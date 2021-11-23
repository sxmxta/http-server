package server

import (
	"encoding/json"
	"io/ioutil"
)

var serverConfig = make(map[string]string)
var proxyConfig = make(map[string]string)

func init() {
	byt, err := ioutil.ReadFile("hs.conf.json")
	if err != nil {
		panic("读取配置文件错误：" + err.Error())
		return
	}
	var data = make(map[string]interface{})
	err = json.Unmarshal(byt, &data)
	if err != nil {
		panic("解析配置文件错误：" + err.Error())
	}
	var serverKey = "server"
	var proxyKey = "proxy"
	var p = data[proxyKey]
	for k, v := range p.(map[string]interface{}) {
		var val = v.(map[string]interface{})
		var vs = val["target"].(string)
		if vs != "" {
			proxyConfig[k] = vs
		}
	}
	var server = data[serverKey]
	for k, v := range server.(map[string]interface{}) {
		var val = v.(string)
		if val != "" {
			serverConfig[serverKey+"."+k] = val
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
