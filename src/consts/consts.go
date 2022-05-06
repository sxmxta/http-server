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
	P2            //发送代理请求
	P3            //代理请求响应失败
	P4            //代理请求响应成功
	P5            //响应给客户端失败
)

const (
	//10:开始请求拦截 11:结束请求拦截， 20:开始响应拦截 21:结束响应拦截
	SIGNAL01 = 1  //等待这个继续处理通知信号
	SIGNAL10 = 10 //开始请求拦截
	SIGNAL11 = 11 //结束请求拦截
	SIGNAL20 = 20 //开始响应拦截
	SIGNAL21 = 21 //结束响应拦截
	SIGNAL22 = 22 //请求响应失败
	SIGNAL23 = 23 //请求响应成功
	SIGNAL24 = 24 //响应客户端失败
	SIGNAL30 = 30 //拦截结束
)

//1=添加 2=删除 3=修改
type PIOption int

const (
	//1=添加 2=删除 3=修改
	PIOption1 = iota + 1
	PIOption2
	PIOption3
)
