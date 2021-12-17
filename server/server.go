package server

import (
	"bytes"
	"fmt"
	"http_server/common"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var contentType = map[string]string{
	".css":  "text/css",
	".js":   "application/javascript",
	".html": "text/html",
	".htm":  "text/html",
	".txt":  "text/plain",
	".png":  "image/png",
	".gif":  "image/gif",
	".jpg":  "image/jpeg",
	".bmp":  "image/bmp",
	".jpeg": "image/jpeg",
	".ico":  "image/ico",
	".json": "application/json",
}

var sites = "sites"

var routeMapper = make(map[string]HandlerFUNC)

type HandlerFUNC func(ctx *Context)

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	isWrite  bool
}

func (m *Context) Write(data []byte) {
	if !m.isWrite {
		m.isWrite = true
		m.response.Header().Set("Content-Type", "application/json")
		m.response.WriteHeader(200)
		m.response.Write(data)
	}
}

func (m *Context) Get(key string) string {
	var val string
	val = m.request.URL.Query().Get(key)
	if val == "" {
		if m.Form() == nil {
			m.request.ParseForm()
		}
		val = m.Form().Get(key)
		if val == "" {
			val = m.PostForm().Get(key)
		}
	}
	return val
}

func (m *Context) Form() url.Values {
	return m.request.Form
}

func (m *Context) PostForm() url.Values {
	return m.request.PostForm
}

func (m *Context) GetBody() string {
	result, err := ioutil.ReadAll(m.request.Body)
	if err != nil {
		return "{}"
	} else {
		return bytes.NewBuffer(result).String()
	}
}

type Handler interface {
}

func RegisterRoute(route string, handler HandlerFUNC) {
	fmt.Println("register route -> ", route)
	routeMapper[route] = handler
}

func StartHttpServer() {
	var serverIP = common.GetServerConf("server.ip")
	var serverPort = common.GetServerConf("server.port")

	if serverIP == "" {
		serverIP = "127.0.0.1"
	}
	if serverPort == "" {
		serverPort = "80"
	}
	addr := serverIP + ":" + serverPort
	fmt.Println("http server listen:", addr)
	mux := http.NewServeMux()
	mux.Handle("/", &HttpServerHandler{})
	_ = http.ListenAndServe(addr, mux)
}

type HttpServerHandler struct{}

func (*HttpServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path
	if ok, proxyUrl, target := isProxy(path); ok {
		proxy(proxyUrl, target, w, r)
	} else {
		if fn, ok := routeMapper[path]; ok {
			w.Header().Set("Content-Type", "application/json")
			ctx := &Context{w, r, false}
			fn(ctx)
			ctx.Write(nil)
			return
		}
		if path == "/" {
			path = "/index.html"
		}
		var (
			byt []byte
			err error
		)
		var filePath = fmt.Sprintf("%s%s", sites, path)
		byt, err = ioutil.ReadFile(filePath)
		w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			return
		}
		if err != nil {
			var content = `{"code":"911","data":"未找到内部操作资源"}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			_, _ = w.Write([]byte(content))
		} else {
			et := extType(path)
			if et != "" {
				if ct, ok := contentType[et]; ok {
					w.Header().Set("Content-Type", ct)
				}
			}
			w.WriteHeader(200)
			_, _ = w.Write(byt)
		}
	}
}

func proxy(proxyUrl, target string, w http.ResponseWriter, r *http.Request) {
	//fmt.Println("url: ", r.URL)
	//查看url各个信息
	//fmt.Print(r.Host, " ", r.Method, " \nr.URL.String ", r.URL.String(), " r.URL.Host ", r.URL.Host, " r.URL.Fragment ", r.URL.Fragment, " r.URL.Hostname ", r.URL.Hostname(), " r.URL.RequestURI ", r.URL.RequestURI(), " r.URL.Scheme ", r.URL.Scheme)
	cli := &http.Client{}
	reqUrl := r.URL.String()
	reqUrl = reqUrl[len(proxyUrl):]
	reqUrl = fmt.Sprintf("%s%s", target, reqUrl)
	fmt.Println("proxy url:", reqUrl, "method:", r.Method)
	req, err := http.NewRequest(r.Method, reqUrl, r.Body)
	if err != nil {
		fmt.Print("proxy http.NewRequest ", err.Error())
		return
	}
	for k, v := range r.Header {
		for _, vs := range v {
			req.Header.Add(k, vs)
		}
	}
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println("proxy error:", err.Error())
		return
	}
	defer res.Body.Close()
	for k, v := range res.Header {
		for _, vs := range v {
			w.Header().Add(k, vs)
		}
	}
	wi, err := io.Copy(w, res.Body)
	fmt.Println("proxy response size:", wi, err)
}

func extType(path string) string {
	idx := strings.LastIndex(path, ".")
	if idx != -1 {
		return path[idx:]
	}
	return ""
}

func isProxy(path string) (bool, string, string) {
	for proxyUrl, v := range common.GetProxyConfig() {
		if strings.Index(path, proxyUrl) == 0 {
			return true, proxyUrl, v
		}
	}
	return false, "", ""
}
