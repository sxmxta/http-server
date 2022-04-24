package config

import (
	"encoding/json"
	"gitee.com/snxamdf/http-server/src/consts"
	"io/ioutil"
)

var Cfg = &Config{}

type Config struct {
	Server Server
	Proxy
	Sqlite3 Sqlite3
}

type Server struct {
	IP   string `json:"ip"`
	PORT string `json:"port"`
}

type Proxy struct {
	Proxy map[string]ProxyTarget `json:"proxy"`
}

type ProxyTarget struct {
	Target string `json:"target"`
	ProxyTargetRewrite
}

type ProxyTargetRewrite struct {
	Rewrite map[string]string `json:"rewrite"`
}

type Sqlite3 struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

func (m *Config) ToJSON() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

func (m *Proxy) ToJSON() []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

func init() {
	defer func() {
		if err := recover(); err != nil {
			consts.GlobalPanicRecoverString = "读取配置文件 致命错误 " + (err.(error)).Error()
		}
	}()
	byt, err := ioutil.ReadFile("hs.conf.json")
	if err != nil {
		//panic("读取配置文件错误：" + err.Error())
		consts.GlobalPanicRecoverString = "读取配置文件错误：" + err.Error()
		return
	}
	err = json.Unmarshal(byt, Cfg)
	if err != nil {
		//panic("解析配置文件错误：" + err.Error())
		consts.GlobalPanicRecoverString = "解析配置文件错误：" + err.Error()
		return
	}
}
