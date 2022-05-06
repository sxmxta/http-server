package server

import (
	"bytes"
	"fmt"
	"gitee.com/snxamdf/http-server/src/config"
	"gitee.com/snxamdf/http-server/src/consts"
	"gitee.com/snxamdf/http-server/src/entity"
	"io"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
)

var id int32
var jar, _ = cookiejar.New(nil)
var interceptConfig = make(map[int32]*entity.ProxyInterceptConfig)

//监听 代理拦截配置数据传输通道
func proxyInterceptConfigChanListen() {
	for {
		select {
		case picc := <-entity.ProxyInterceptConfigChan:
			if picc.Option == consts.PIOption1 || picc.Option == consts.PIOption3 {
				//添加
				interceptConfig[picc.Index] = picc
			} else if picc.Option == consts.PIOption2 {
				//删除
				//var before = interceptConfig[:picc.Index]
				//var after = interceptConfig[picc.Index+1:]
				//interceptConfig = append(before, after...)
				delete(interceptConfig, picc.Index)
			}
			fmt.Println(picc.Index, picc.Option, picc.Enable(), picc.InterceptUrl(), len(interceptConfig))
		}
	}
}

//代理服务
func proxyServer(proxyAddr *proxyAddr, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			entity.PutTimeMessage("Http Proxy Server 致命错误:" + err.(error).Error())
		}
	}()
	//查看url各个信息
	//fmt.Print(r.Host, " ", r.Method, " \nr.URL.String ", r.URL.String(), " r.URL.Host ", r.URL.Host, " r.URL.Fragment ", r.URL.Fragment, " r.URL.Hostname ", r.URL.Hostname(), " r.URL.RequestURI ", r.URL.RequestURI(), " r.URL.Scheme ", r.URL.Scheme)
	var (
		request             *http.Request
		response            *http.Response
		err                 error
		proxyDetailGridData *entity.ProxyDetail
		wi                  int64
		isInter             bool
		signal              int32 //信号
	)
	cli := &http.Client{
		Jar: jar,
	}
	//启用代理详情 记录 详情 请求
	if entity.EnableProxyDetail {
		atomic.AddInt32(&id, 1)
		var id = atomic.LoadInt32(&id)
		//err = r.ParseForm()
		proxyDetailGridData = &entity.ProxyDetail{
			ID:             id,
			SourceUrl:      proxyAddr.sourceUrl,
			TargetUrl:      proxyAddr.targetUrl,
			Method:         r.Method,
			Host:           r.Host,
			State:          consts.P0,
			IsAddTaskQueue: true,
			StateCode:      0,
			Request: entity.ProxyRequestDetail{
				URLParams:  r.URL.Query(),
				FormParams: r.PostForm,
			},
			Response:             entity.ProxyResponseDetail{},
			ProxyInterceptSignal: make(chan int32),
		}
		buf := new(bytes.Buffer)
		wi, err = buf.ReadFrom(r.Body)
		if err == nil {
			proxyDetailGridData.Request.Body = buf.String()
			proxyDetailGridData.Request.Header = r.Header
			proxyDetailGridData.Request.Size = wi
			request, err = http.NewRequest(r.Method, proxyAddr.targetUrl, buf)
		}
	} else {
		request, err = http.NewRequest(r.Method, proxyAddr.targetUrl, r.Body)
	}
	if err != nil {
		//启用代理详情 记录 详情 请求
		if proxyDetailGridData != nil {
			proxyDetailGridData.Error = err
			proxyDetailGridData.State = consts.P1
			proxyDetailGridData.StateCode = 1
			entity.ProxyDetailGridChan <- proxyDetailGridData
		}
		entity.PutLogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy http.NewRequest ", err.Error())
		return
	}

	//头设置
	//request.Host = r.Host
	for k, v := range r.Header {
		for _, vs := range v {
			request.Header.Add(k, vs)
		}
	}

	//发起代理请求
	//启用代理详情 记录 详情 请求
	if proxyDetailGridData != nil {
		proxyDetailGridData.State = consts.P2
		entity.ProxyDetailGridChan <- proxyDetailGridData
	}
	//开启拦截
	if entity.ProxyInterceptEnable {
		//没啥问题 判断是否为发送请求
		if proxyDetailGridData != nil && proxyDetailGridData.State == consts.P2 {
			var pic *entity.ProxyInterceptConfig
			for _, pic = range interceptConfig {
				if pic.Enable() && strings.TrimSpace(pic.InterceptUrl()) != "" {
					//判断当前地址是被拦截
					isInter = strings.Contains(proxyAddr.targetUrl, pic.InterceptUrl())
					if isInter {
						//fmt.Println(pic.InterceptUrl(), proxyAddr.targetUrl)
						break
					}
				}
			}
			//请求地址拦截
			if isInter && pic != nil {
				//添加到队列准备-然后等待通知信号继续处理
				entity.ProxyFlowInterceptChan <- proxyDetailGridData
				//等待继续处理通知信号 01, 所有后面被拦截的都会阻塞在这里等待通知
				signal = <-proxyDetailGridData.ProxyInterceptSignal
				if signal != consts.SIGNAL01 {
					fmt.Println("程序混乱 signal 信号不正确,应为 01, 实际值", signal)
				}
				proxyDetailGridData.ProxyInterceptSignal <- consts.SIGNAL10 //发送信号触发request拦截 10
				signal = <-proxyDetailGridData.ProxyInterceptSignal         //等待触发拦截结果 11
				if signal != consts.SIGNAL11 {
					fmt.Println("程序混乱 signal 信号不正确,应为 11, 实际值", signal)
				}
				//fmt.Println("signal", signal)
			}
		}
	}
	if proxyDetailGridData != nil {
		//之后的操作都不向列队中添加
		proxyDetailGridData.IsAddTaskQueue = false
	}
	response, err = cli.Do(request)
	if err != nil {
		entity.PutLogsProxyTime(err.Error())
		//启用代理详情 记录 详情 请求
		if proxyDetailGridData != nil { //出错了
			proxyDetailGridData.Error = err
			proxyDetailGridData.State = consts.P3
			proxyDetailGridData.StateCode = 504
			entity.ProxyDetailGridChan <- proxyDetailGridData
			//响应失败
			if isInter {
				entity.ProxyFlowInterceptChan <- proxyDetailGridData
				proxyDetailGridData.ProxyInterceptSignal <- consts.SIGNAL22 //信号 失败
			}
		}
		buf := new(bytes.Buffer)
		buf.WriteString(err.Error())
		w.WriteHeader(504)
		wi, err = io.Copy(w, buf)
	} else {
		defer response.Body.Close()
		//处理代理原样返回给客户端
		w.WriteHeader(response.StatusCode)
		for k, v := range response.Header {
			for _, vs := range v {
				w.Header().Add(k, vs)
			}
		}
		//启用代理详情 记录 详情 请求
		if proxyDetailGridData != nil {
			buf := new(bytes.Buffer)
			wi, err = buf.ReadFrom(response.Body)
			if err == nil {
				proxyDetailGridData.Response.Body = buf.String()
				proxyDetailGridData.Response.Header = response.Header
				proxyDetailGridData.Response.Size = wi
				proxyDetailGridData.State = consts.P4
				proxyDetailGridData.StateCode = response.StatusCode
				entity.ProxyDetailGridChan <- proxyDetailGridData
				//响应成功
				if isInter && entity.ProxyInterceptEnable {
					entity.ProxyFlowInterceptChan <- proxyDetailGridData
					proxyDetailGridData.ProxyInterceptSignal <- consts.SIGNAL20
					signal := <-proxyDetailGridData.ProxyInterceptSignal //等待触发拦截结果 21
					fmt.Println("signal", signal)
				}
				_, err = w.Write(buf.Bytes())
			}
		} else {
			wi, err = io.Copy(w, response.Body)
		}
		if err != nil {
			//启用代理详情 记录 详情 请求
			if proxyDetailGridData != nil {
				proxyDetailGridData.Error = err
				proxyDetailGridData.State = consts.P5
				proxyDetailGridData.StateCode = response.StatusCode
				entity.ProxyDetailGridChan <- proxyDetailGridData
				//响应客户端失败
				if isInter {
					entity.ProxyFlowInterceptChan <- proxyDetailGridData
					proxyDetailGridData.ProxyInterceptSignal <- consts.SIGNAL24
				}
			}
			entity.PutLogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy response size:", strconv.Itoa(int(wi)), err.Error())
		} else {
			//响应客户端成功
			if isInter && proxyDetailGridData != nil {
				entity.ProxyFlowInterceptChan <- proxyDetailGridData
				proxyDetailGridData.ProxyInterceptSignal <- consts.SIGNAL23
			}
			entity.PutLogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy response size:", strconv.Itoa(int(wi)))
		}
	}
	//响应完成
	if isInter && proxyDetailGridData != nil {
		proxyDetailGridData.ProxyInterceptSignal <- consts.SIGNAL30
	}
}

type proxyAddr struct {
	cliReqHost string
	sourceUrl  string
	matchUrl   string
	targetUrl  string
	rewrite    map[string]string
}

//url = / /path /path/ /path/path /path/path/
func isProxy(r *http.Request) (*proxyAddr, bool) {
	urlPath := r.URL.String()
	for matchUrl, v := range config.Cfg.Proxy.Proxy {
		if strings.Index(urlPath, matchUrl) == 0 {
			var buf strings.Builder
			if r.Proto[4] == '/' {
				buf.WriteString("http://")
			} else if r.Proto[4] == '5' {
				buf.WriteString("https://")
			}
			buf.WriteString(r.Host)
			p := &proxyAddr{sourceUrl: urlPath, cliReqHost: buf.String()}
			p.matchUrl = matchUrl
			p.targetUrl = v.Target
			p.rewrite = v.Rewrite
			//解析代理配置地址
			p.init()
			return p, true
		}
	}
	return nil, false
}

func (m *proxyAddr) init() {
	// /dam/service/configs/get_one?name=configurable.web.dam.appname&callback=jQuery360046561208122962516_1650637061068&_=1650637061069
	//判断是否替换 /url/ /url
	var idxLast = strings.LastIndex(m.matchUrl, "/")
	var isReplace = len(m.matchUrl)-1 == idxLast
	var url = m.sourceUrl
	if isReplace {
		//替换
		url = url[idxLast:]
	}
	//重写替换
	if m.rewrite != nil {
		for k, v := range m.rewrite {
			compile := regexp.MustCompile(k)
			url = compile.ReplaceAllString(url, v)
		}
	}
	m.targetUrl = fmt.Sprintf("%s%s", m.targetUrl, url)
	m.sourceUrl = fmt.Sprintf("%s%s", m.cliReqHost, m.sourceUrl)
}
