package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

func StartHttpServer() {
	var serverIP = GetServerConf("server.ip")
	var serverPort = GetServerConf("server.port")

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
	for proxyUrl, v := range proxyConfig {
		if strings.Index(path, proxyUrl) == 0 {
			return true, proxyUrl, v
		}
	}
	return false, "", ""
}
