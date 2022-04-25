package consts

import "strings"

var (
	AppInitSuccess = true
	GlobalMessage  = make(chan MessageChannel)
	HttpMethods    = []string{"GET", "POST", "HEAD", "PUT", "DELETE", "CONNECT", "OPTIONS"}
)

type MessageChannel struct {
	Type    int
	Message []string
	Color   int32
}

func GetHttpMethodsIdx(methodName string) int {
	methodName = strings.ToUpper(strings.Trim(methodName, " "))
	for i, method := range HttpMethods {
		if method == methodName {
			return i
		}
	}
	return -1
}

//普通消息
func PutMessage(message ...string) {
	go func() { GlobalMessage <- MessageChannel{Type: 0, Message: message} }()
}

//带颜色的
func PutColorMessage(color int32, message ...string) {
	go func() { GlobalMessage <- MessageChannel{Type: 1, Message: message, Color: color} }()
}

//带日期的
func PutTimeMessage(message ...string) {
	go func() { GlobalMessage <- MessageChannel{Type: 2, Message: message} }()
}

//代理日志
func PutLogsProxyTime(message ...string) {
	go func() { GlobalMessage <- MessageChannel{Type: 3, Message: message} }()
}

//普通日志
func PutLogsStaticTime(message ...string) {
	go func() { GlobalMessage <- MessageChannel{Type: 4, Message: message} }()
}
