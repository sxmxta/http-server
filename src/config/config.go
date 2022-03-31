package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config map[string]string

var serverConfig = make(Config)
var proxyConfig = make(Config)
var dbConfig = make(Config)

func (m Config) Get(key string) string {
	return m[key]
}

func (m Config) ToJSONString() string {
	s, _ := json.Marshal(m)
	return string(s)
}

func GetServerConfig() Config {
	return serverConfig
}

func GetProxyConfig() Config {
	return proxyConfig
}
func GetDBConfig() Config {
	return dbConfig
}

var ParseConfigErr string

func init() {
	defer func() {
		if err := recover(); err != nil {
			ParseConfigErr = "读取配置文件 致命错误"
		}
	}()
	byt, err := ioutil.ReadFile("hs.conf.json")
	if err != nil {
		//panic("读取配置文件错误：" + err.Error())
		ParseConfigErr = "读取配置文件错误：" + err.Error()
		return
	}
	var data = make(map[string]interface{})
	err = json.Unmarshal(byt, &data)
	if err != nil {
		//panic("解析配置文件错误：" + err.Error())
		ParseConfigErr = "解析配置文件错误：" + err.Error()
		return
	}
	var serverKey = "server"
	var proxyKey = "proxy"
	var databaseKey = "database"
	var p = data[proxyKey]
	for k, v := range p.(map[string]interface{}) {
		var val = v.(map[string]interface{})
		var vs = val["target"].(string)
		if vs != "" {
			proxyConfig[k] = vs
		}
	}
	var server = data[serverKey]
	if server != nil {
		for k, v := range server.(map[string]interface{}) {
			var val = v.(string)
			if val != "" {
				serverConfig[serverKey+"."+k] = val
			}
		}
	}

	var db = data[databaseKey]
	if db != nil {
		for k, v := range db.(map[string]interface{}) {
			var val = v.(string)
			if val != "" {
				dbConfig[databaseKey+"."+k] = val
			}
		}
	}
	//fmt.Println("解析 服务 配置")
	//for k, v := range serverConfig {
	//	fmt.Println("服务：", k, "=", v)
	//}
	//fmt.Println("解析 代理 配置")
	//for k, v := range proxyConfig {
	//	fmt.Println("代理：", k, "=", v)
	//}
}

func GetServerConf(key string) string {
	if key, ok := serverConfig[key]; ok {
		return key
	}
	return ""
}

func GetProxyConf(key string) string {
	if key, ok := proxyConfig[key]; ok {
		return key
	}
	return ""
}
