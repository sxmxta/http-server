package server

import (
	"bytes"
	"fmt"
	"gitee.com/snxamdf/http-server/src/config"
	"gitee.com/snxamdf/http-server/src/gui"
	"io"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"strings"
)

var id int
var jar, _ = cookiejar.New(nil)

func proxy(proxyAddr *proxyAddr, w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			gui.LogsTime("Http proxy 致命错误")
		}
	}()
	//fmt.Println("url: ", r.URL)
	//查看url各个信息
	//fmt.Print(r.Host, " ", r.Method, " \nr.URL.String ", r.URL.String(), " r.URL.Host ", r.URL.Host, " r.URL.Fragment ", r.URL.Fragment, " r.URL.Hostname ", r.URL.Hostname(), " r.URL.RequestURI ", r.URL.RequestURI(), " r.URL.Scheme ", r.URL.Scheme)
	var (
		request     *http.Request
		response    *http.Response
		err         error
		proxyDetail *gui.ProxyDetail
	)
	cli := &http.Client{
		Jar: jar,
	}
	//启用代理详情 记录 请求 详情
	if gui.GUIForm.EnableProxyDetail {
		id++
		//err = r.ParseForm()
		proxyDetail = &gui.ProxyDetail{
			ID:     id,
			URL:    proxyAddr.targetUrl,
			Method: r.Method,
			Host:   r.Host,
			Request: gui.ProxyRequestDetail{
				URLParams:  r.URL.Query(),
				FormParams: r.PostForm,
			},
			Response: gui.ProxyResponseDetail{},
		}
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(r.Body)
		if err == nil {
			proxyDetail.Request.Body = buf.String()
			proxyDetail.Request.Header = r.Header
			request, err = http.NewRequest(r.Method, proxyAddr.targetUrl, buf)
		}
	} else {
		request, err = http.NewRequest(r.Method, proxyAddr.targetUrl, r.Body)
	}
	if err != nil {
		gui.LogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy http.NewRequest ", err.Error())
		return
	}

	//头设置
	//fmt.Println("proxy.request.Host",request.Host,"  r.Host",r.Host)
	//request.Host = r.Host
	for k, v := range r.Header {
		for _, vs := range v {
			request.Header.Add(k, vs)
		}
	}
	//cookie设置
	//for _, cookie := range r.Cookies() {
	//	request.AddCookie(cookie)
	//}

	//fmt.Println("proxy.request.Host",request.Host,"  r.Host",r.Host)
	//发起代理请求
	response, err = cli.Do(request)
	if err != nil {
		gui.LogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy error:", err.Error())
		return
	}
	defer response.Body.Close()
	//处理代理原样返回给客户端
	w.WriteHeader(response.StatusCode)
	for k, v := range response.Header {
		for _, vs := range v {
			w.Header().Add(k, vs)
		}
	}
	//for _, cookie := range response.Cookies() {
	//	r.AddCookie(cookie)
	//}

	var wi int64
	//启用代理详情 记录 请求 详情
	if gui.GUIForm.EnableProxyDetail && proxyDetail != nil {
		buf := new(bytes.Buffer)
		wi, err = buf.ReadFrom(response.Body)
		if err == nil {
			proxyDetail.Response.Body = buf.String()
			proxyDetail.Response.Header = response.Header
			proxyDetail.Response.Size = wi
			gui.GUIForm.SetProxyDetail(proxyDetail)
			_, err = w.Write(buf.Bytes())
		}
	} else {
		wi, err = io.Copy(w, response.Body)
	}
	if err != nil {
		gui.LogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy response size:", strconv.Itoa(int(wi)), err.Error())
	} else {
		gui.LogsProxyTime("proxy url:  ", proxyAddr.targetUrl, "  method: ", r.Method, "  proxy response size:", strconv.Itoa(int(wi)))
	}
}

func isProxy(r *http.Request) (bool, *proxyAddr) {
	urlPath := r.URL.String()
	proxyMap := config.Cfg.Proxy.Proxy
	p := &proxyAddr{sourceUrl: urlPath}
	for matchUrl, v := range proxyMap {
		if strings.Index(urlPath, matchUrl) == 0 {
			p.matchUrl = matchUrl
			p.targetUrl = v.Target
			p.rewrite = v.Rewrite
			//初始化代理地址
			p.init()
			return true, p
		}
	}
	return false, nil
}

type proxyAddr struct {
	sourceUrl string
	matchUrl  string
	targetUrl string
	rewrite   map[string]string
}

func (m *proxyAddr) init() {
	// /dam/service/configs/get_one?name=configurable.web.dam.appname&callback=jQuery360046561208122962516_1650637061068&_=1650637061069
	//判断是否替换 /url/ /url
	var idxLast = strings.LastIndex(m.matchUrl, "/")
	var isReplace = len(m.matchUrl)-1 == idxLast
	if isReplace {
		//替换
		m.sourceUrl = m.sourceUrl[idxLast:]
	} else {
		//不替换
	}
	if m.rewrite != nil {
		for k, v := range m.rewrite {
			compile := regexp.MustCompile(k)
			m.sourceUrl = compile.ReplaceAllString(m.sourceUrl, v)
		}
	}
	m.targetUrl = fmt.Sprintf("%s%s", m.targetUrl, m.sourceUrl)
	//var isTargetUrlLast = len(m.targetUrl)-1 == strings.LastIndex(m.targetUrl, "/")
	//var isSourceUrlFirst = len(m.sourceUrl)-1 == strings.Index(m.sourceUrl, "/")
	//if !isTargetUrlLast && !isSourceUrlFirst {
	//	m.targetUrl = fmt.Sprintf("%s/%s", m.targetUrl, m.sourceUrl)
	//} else {
	//	m.targetUrl = fmt.Sprintf("%s%s", m.targetUrl, m.sourceUrl)
	//}
}
