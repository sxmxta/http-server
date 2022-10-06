package config

import (
	"bytes"
	"encoding/json"
	"gitee.com/snxamdf/http-server/src/entity"
	"github.com/energye/golcl/lcl/types/colors"
	"io/ioutil"
)

var Cfg = &Config{
	ROOT: "root", //解析网站文件目录,默认root文件夹
}

type Config struct {
	ROOT   string `json:"root"`
	Server Server
	Proxy
	Sqlite3 Sqlite3
}

type Server struct {
	IP          string `json:"ip"`
	PORT        string `json:"port"`
	SSL         bool   `json:"ssl"`
	SSLCert     string `json:"sslCert"`
	SSLKey      string `json:"sslKey"`
	SSLHttp     bool   `json:"sslHttp"`
	SSLHttpPort int    `json:"sslHttpPort"`
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
			entity.AppInitSuccess = false
			entity.PutColorMessage(colors.ClRed, "读取配置文件 致命错误 ", (err.(error)).Error())
		}
	}()
	byt, err := ioutil.ReadFile("hs.conf.json")
	if err != nil {
		entity.AppInitSuccess = false
		entity.PutColorMessage(colors.ClRed, "读取配置文件错误：", err.Error())
		return
	}
	byt = bytes.TrimPrefix(byt, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(byt, Cfg)
	if err != nil {
		entity.AppInitSuccess = false
		entity.PutColorMessage(colors.ClRed, "解析配置文件错误：", err.Error())
		return
	}
}
