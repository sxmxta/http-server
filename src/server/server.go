package server

import (
	"bytes"
	"fmt"
	"gitee.com/snxamdf/golcl/emfs"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/common"
	"gitee.com/snxamdf/http-server/src/config"
	"gitee.com/snxamdf/http-server/src/entity"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var contentType = map[string]string{}

func Init() {
	mimeTypes, err := emfs.GetResources("resources/mime-types.conf")
	if err != nil {
		entity.AppInitSuccess = false
		entity.PutColorMessage(colors.ClRed, err.Error())
	} else {
		var types = strings.Split(string(mimeTypes), "\n")
		for _, mime := range types {
			mime = strings.TrimSpace(mime)
			if mime != "" && mime[0] != '#' {
				var m = strings.Split(mime, "=")
				if len(m) == 2 {
					contentType["."+m[0]] = m[1]
				}
			}
		}
	}
}

var sites = "sites"

var routeMapper = make(map[string]HandlerFUNC)

type Route interface {
}

type Controller struct {
	ctx *Context
}

type HandlerFUNC func(ctx *Context)

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	isWrite  bool
}

func (m *Controller) Context() *Context {
	return m.ctx
}

func (m *Context) Write(data interface{}) {
	if !m.isWrite {
		m.isWrite = true
		m.response.Header().Set("Content-Type", "application/json")
		m.response.WriteHeader(200)
		//d, _ := common.ToJson(data)
		r := make(map[string]interface{})
		r["code"] = 200
		r["data"] = data
		var b, err = common.ToJson(r)
		if err != nil {
			entity.PutMessage(err.Error())
		}
		m.response.Write(b)
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

func (m *Context) GetBody() []byte {
	result, err := ioutil.ReadAll(m.request.Body)
	if err != nil {
		return []byte("{}")
	} else {
		return bytes.NewBuffer(result).Bytes()
	}
}

type Handler interface {
}

func RegisterRoute(route string, handler HandlerFUNC) {
	entity.PutMessage("register route -> ", route)
	routeMapper[route] = handler
}

func StartHttpServer() error {
	Init()
	var serverIP = config.Cfg.Server.IP
	var serverPort = config.Cfg.Server.PORT

	if serverIP == "" {
		serverIP = "127.0.0.1"
	}
	if serverPort == "" {
		serverPort = "80"
	}
	addr := serverIP + ":" + serverPort
	mux := http.NewServeMux()
	mux.Handle("/", &HttpServerHandler{})
	t := time.Now()
	msg := t.Format("2006-01-02 15:04:05")
	entity.PutMessage("Http Server 启动中......")
	entity.PutMessage("Http Server 启动时间: " + msg)
	entity.PutMessage(fmt.Sprintf("%v: %v", "Http Server Listen:", addr))
	entity.PutMessage("Http Server Proxy: ", string(config.Cfg.Proxy.ToJSON()))
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		entity.PutMessage("Http Server 启动失败")
		entity.PutColorMessage(colors.ClRed, err.Error())
	}
	return err
}

type HttpServerHandler struct{}

func (*HttpServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			entity.PutMessage("Http 致命错误")
		}
	}()
	var path = r.URL.Path
	if ok, proxyAddr := isProxy(r); ok {
		proxy(proxyAddr, w, r)
	} else {
		//w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		entity.PutMessage("请求URL:", path)
		//if r.Method == "OPTIONS" {
		//	return
		//}
		if fn, ok := routeMapper[path]; ok {
			w.Header().Set("Content-Type", "application/json")
			ctx := &Context{w, r, false}
			fn(ctx)
			ctx.Write(nil)
			return
		}
		if path == "/" {
			path = "/index.html"
		} else if strings.LastIndex(path, "/") == len(path)-1 {
			path = path + "index.html"
		}
		var (
			byt []byte
			err error
		)
		var filePath = fmt.Sprintf("%s%s", sites, path)
		byt, err = ioutil.ReadFile(filePath)
		if err != nil {
			var content = `{"code":"404","data":"你访问的地址不存在"}`
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

func extType(path string) string {
	idx := strings.LastIndex(path, ".")
	if idx != -1 {
		return path[idx:]
	}
	return ""
}
