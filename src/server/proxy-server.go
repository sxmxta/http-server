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

func proxy(proxyAddr *proxyAddr, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			entity.PutLogsProxyTime("Http proxy 致命错误")
		}
	}()
	//查看url各个信息
	//fmt.Print(r.Host, " ", r.Method, " \nr.URL.String ", r.URL.String(), " r.URL.Host ", r.URL.Host, " r.URL.Fragment ", r.URL.Fragment, " r.URL.Hostname ", r.URL.Hostname(), " r.URL.RequestURI ", r.URL.RequestURI(), " r.URL.Scheme ", r.URL.Scheme)
	var (
		request     *http.Request
		response    *http.Response
		err         error
		proxyDetail *entity.ProxyDetail
		wi          int64
	)
	cli := &http.Client{
		Jar: jar,
	}
	//启用代理详情 记录 详情 请求
	if entity.EnableProxyDetail {
		atomic.AddInt32(&id, 1)
		var id = atomic.LoadInt32(&id)
		//err = r.ParseForm()
		proxyDetail = &entity.ProxyDetail{
			ID:        id,
			SourceUrl: proxyAddr.sourceUrl,
			TargetUrl: proxyAddr.targetUrl,
			Method:    r.Method,
			Host:      r.Host,
			State:     consts.P0,
			StateCode: 0,
			Request: entity.ProxyRequestDetail{
				URLParams:  r.URL.Query(),
				FormParams: r.PostForm,
			},
			Response: entity.ProxyResponseDetail{},
		}
		buf := new(bytes.Buffer)
		wi, err = buf.ReadFrom(r.Body)
		if err == nil {
			proxyDetail.Request.Body = buf.String()
			proxyDetail.Request.Header = r.Header
			proxyDetail.Request.Size = wi
			request, err = http.NewRequest(r.Method, proxyAddr.targetUrl, buf)
		}
	} else {
		request, err = http.NewRequest(r.Method, proxyAddr.targetUrl, r.Body)
	}
	if err != nil {
		//启用代理详情 记录 详情 请求
		if proxyDetail != nil {
			proxyDetail.Error = err
			proxyDetail.State = consts.P1
			proxyDetail.StateCode = 1
			entity.ProxyDetailChan <- proxyDetail
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
	if proxyDetail != nil {
		proxyDetail.State = consts.P2
		entity.ProxyDetailChan <- proxyDetail
	}
	response, err = cli.Do(request)
	if err != nil {
		entity.PutLogsProxyTime(err.Error())
		//启用代理详情 记录 详情 请求
		if proxyDetail != nil {
			proxyDetail.Error = err
			proxyDetail.State = consts.P3
			proxyDetail.StateCode = 504
			entity.ProxyDetailChan <- proxyDetail
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
		if proxyDetail != nil {
			buf := new(bytes.Buffer)
			wi, err = buf.ReadFrom(response.Body)
			if err == nil {
				proxyDetail.Response.Body = buf.String()
				proxyDetail.Response.Header = response.Header
				proxyDetail.Response.Size = wi
				proxyDetail.State = consts.P4
				proxyDetail.StateCode = response.StatusCode
				entity.ProxyDetailChan <- proxyDetail
				_, err = w.Write(buf.Bytes())
			}
		} else {
			wi, err = io.Copy(w, response.Body)
		}
		if err != nil {
			//启用代理详情 记录 详情 请求
			if proxyDetail != nil {
				proxyDetail.Error = err
				proxyDetail.State = consts.P5
				proxyDetail.StateCode = response.StatusCode
				entity.ProxyDetailChan <- proxyDetail
			}
			entity.PutLogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy response size:", strconv.Itoa(int(wi)), err.Error())
		} else {
			entity.PutLogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy response size:", strconv.Itoa(int(wi)))
		}
	}
}

type proxyAddr struct {
	sourceUrl string
	matchUrl  string
	targetUrl string
	rewrite   map[string]string
}

//url = / /path /path/ /path/path /path/path/
func isProxy(r *http.Request) (bool, *proxyAddr) {
	urlPath := r.URL.String()
	for matchUrl, v := range config.Cfg.Proxy.Proxy {
		if strings.Index(urlPath, matchUrl) == 0 {
			p := &proxyAddr{sourceUrl: urlPath}
			p.matchUrl = matchUrl
			p.targetUrl = v.Target
			p.rewrite = v.Rewrite
			//解析代理配置地址
			p.init()
			return true, p
		}
	}
	return false, nil
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
}
