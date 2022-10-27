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
)

var jar, _ = cookiejar.New(nil)
var interceptConfig *[]*entity.ProxyInterceptConfig

//监听 代理拦截配置数据传输通道
func proxyInterceptConfigChanListen() {
	for {
		select {
		case picc := <-entity.ProxyInterceptConfigChan:
			interceptConfig = picc
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
		//signal              int32 //信号
	)
	cli := &http.Client{
		Jar: jar,
	}
	//启用代理详情 记录 详情 请求
	if entity.EnableProxyDetail {
		request, proxyDetailGridData, isInter, err = handlerEnableProxy(entity.ID.Get(), proxyAddr, r)
		//err = r.ParseForm()
		//proxyDetailGridData = &entity.ProxyDetail{
		//	ID:             id,
		//	SourceUrl:      proxyAddr.sourceUrl,
		//	TargetUrl:      proxyAddr.targetUrl,
		//	Method:         r.Method,
		//	Host:           r.Host,
		//	State:          consts.P0,
		//	IsAddTaskQueue: true,
		//	StateCode:      0,
		//	Request: entity.ProxyRequestDetail{
		//		URLParams: r.URL.Query(),
		//		PostForm:  r.PostForm,
		//	},
		//	Response:             entity.ProxyResponseDetail{},
		//	ProxyInterceptSignal: make(chan int32),
		//}
		////r.MultipartForm
		////r.Form
		////r.PostForm
		////r.ParseForm()
		////r.ParseMultipartForm()
		////读取请求体内容
		//buf := new(bytes.Buffer)
		//wi, err = buf.ReadFrom(r.Body)
		//if err == nil {
		//	//设置数据到 proxyDetailGridData
		//	proxyDetailGridData.Request.Body = buf.String()
		//	proxyDetailGridData.Request.Header = r.Header
		//	proxyDetailGridData.Request.Size = wi
		//	proxyDetailGridData.State = consts.P2
		//	//发送数据到 列表展示
		//	entity.ProxyDetailGridChan <- proxyDetailGridData
		//
		//	//开启拦截配置
		//	if entity.ProxyInterceptConfigEnable {
		//		//判断发送请求是否被拦截
		//		var pic *entity.ProxyInterceptConfig
		//		for _, pic = range *interceptConfig {
		//			if pic.Enable() && strings.TrimSpace(pic.InterceptUrl()) != "" {
		//				//当前地址被拦截
		//				isInter = strings.Contains(proxyAddr.targetUrl, pic.InterceptUrl())
		//				if isInter {
		//					break
		//				}
		//			}
		//		}
		//		//此处将会得到被拦截的请求修改的信息包括：请求参数、头，体、cookie等，体现在 SIGNAL10 之后
		//		if isInter && pic != nil {
		//			//添加到队列准备-然后等待通知信号继续处理
		//			entity.ProxyFlowInterceptChan <- proxyDetailGridData
		//			//等待继续处理通知信号 01, 所有后面被拦截的都会阻塞在这里等待通知
		//			signal = <-proxyDetailGridData.ProxyInterceptSignal
		//			if signal != consts.SIGNAL01 {
		//				fmt.Println("程序混乱 signal 信号不正确,应为 01, 实际值", signal)
		//			}
		//			//发送信号触发request拦截 10
		//			proxyDetailGridData.ProxyInterceptSignal <- consts.SIGNAL10
		//			//等待触发拦截结果 11
		//			signal = <-proxyDetailGridData.ProxyInterceptSignal
		//			if signal != consts.SIGNAL11 {
		//				fmt.Println("程序混乱 signal 信号不正确,应为 11, 实际值", signal)
		//			}
		//			//fmt.Println("signal", signal)
		//		}
		//	}
		//	//创建新的代理请求对象
		//	request, err = http.NewRequest(r.Method, proxyAddr.targetUrl, buf)
		//}
		if err == nil {
			//之后的操作都不向任务列队中添加
			proxyDetailGridData.IsAddTaskQueue = false
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

	//修改请求头
	updateProxyRequestHeader(request, r, proxyDetailGridData)

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
		for k, v := range response.Header {
			fmt.Println("header:", k, v)
			for _, vs := range v {
				w.Header().Add(k, vs)
			}
		}
		//go的要写在w.Header().Add之后
		w.WriteHeader(response.StatusCode)
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
				if isInter && entity.ProxyInterceptConfigEnable {
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

//处理启用拦截
func handlerEnableProxy(id int32, proxyAddr *proxyAddr, r *http.Request) (request *http.Request, proxyDetailGridData *entity.ProxyDetail, isInter bool, err error) {
	var (
		signal int32 //信号
		wi     int64
	)
	//err = r.ParseForm()
	//if err != nil {
	//	return request, proxyDetailGridData, isInter, err
	//}
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
			URLParams: r.URL.Query(),
			PostForm:  r.PostForm,
			//Form:      r.Form,
		},
		Response:             entity.ProxyResponseDetail{},
		ProxyInterceptSignal: make(chan int32),
	}
	//r.MultipartForm
	//r.Form
	//r.PostForm
	//r.ParseForm()
	//r.ParseMultipartForm()
	//读取请求体内容
	buf := new(bytes.Buffer)
	wi, err = buf.ReadFrom(r.Body)
	if err == nil {
		//设置数据到 proxyDetailGridData
		proxyDetailGridData.Request.Body = buf.String()
		proxyDetailGridData.Request.Header = r.Header
		proxyDetailGridData.Request.Size = wi
		proxyDetailGridData.State = consts.P2
		//发送数据到 列表展示
		entity.ProxyDetailGridChan <- proxyDetailGridData

		//开启拦截配置
		if entity.ProxyInterceptConfigEnable {
			//判断发送请求是否被拦截
			var pic *entity.ProxyInterceptConfig
			for _, pic = range *interceptConfig {
				if pic.Enable() && strings.TrimSpace(pic.InterceptUrl()) != "" {
					//当前地址被拦截
					isInter = strings.Contains(proxyAddr.targetUrl, pic.InterceptUrl())
					if isInter {
						break
					}
				}
			}
			//此处将会得到被拦截的请求修改的信息包括：请求参数、头，体、cookie等，体现在 SIGNAL10 之后
			if isInter && pic != nil {
				//添加到队列准备-然后等待通知信号继续处理
				entity.ProxyFlowInterceptChan <- proxyDetailGridData
				//等待继续处理通知信号 01, 所有后面被拦截的都会阻塞在这里等待通知 上一个结束
				signal = <-proxyDetailGridData.ProxyInterceptSignal
				if signal != consts.SIGNAL01 {
					fmt.Println("程序混乱 signal 信号不正确,应为 01, 实际值", signal)
				}
				//发送信号触发request拦截 10
				proxyDetailGridData.ProxyInterceptSignal <- consts.SIGNAL10
				//等待触发拦截结果 11 button
				signal = <-proxyDetailGridData.ProxyInterceptSignal
				if signal != consts.SIGNAL11 {
					fmt.Println("程序混乱 signal 信号不正确,应为 11, 实际值", signal)
				}
				//fmt.Println("signal", signal)
			}
		}
	}
	//创建新的代理请求对象
	request, err = http.NewRequest(r.Method, proxyAddr.targetUrl, buf)
	return request, proxyDetailGridData, isInter, err
}

//修改代理请求
func updateProxyRequestHeader(newRequest *http.Request, oldRequest *http.Request, proxyDetailGridData *entity.ProxyDetail) {
	//request.Host = r.Host
	//nil 未开启代理
	if proxyDetailGridData == nil {
		for k, v := range oldRequest.Header {
			for _, vs := range v {
				newRequest.Header.Add(k, vs)
			}
		}
	} else {
		// 开启代理，使用代理UI界面修改后的数据
		for k, v := range proxyDetailGridData.Request.Header {
			for _, vs := range v {
				newRequest.Header.Add(k, vs)
			}
		}
	}
}

//修改代理响应
func updateProxyResponse(proxyDetailGridData *entity.ProxyDetail) {

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
