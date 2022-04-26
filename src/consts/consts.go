package consts

import "strings"

var (
	HttpMethods = []string{"GET", "POST", "HEAD", "PUT", "DELETE", "CONNECT", "OPTIONS"}
)

//获取HttpMethod 下标
func GetHttpMethodsIdx(methodName string) int {
	methodName = strings.ToUpper(strings.Trim(methodName, " "))
	for i, method := range HttpMethods {
		if method == methodName {
			return i
		}
	}
	return -1
}

type ProxyFlow int32

const (
	P0 = iota + 0 //初始代理请求
	P1            //创建代理请求失败
	P2            //代理请求响应失败
	P3            //代理请求响应成功
	P4            //响应给客户端失败
)
