package entity

//消息日志
type LogMessage struct {
	Type    int
	Message []string
	Color   int32
}

//普通消息
func PutMessage(message ...string) {
	go func() { GlobalLogMessageChan <- &LogMessage{Type: 0, Message: message} }()
}

//带颜色的
func PutColorMessage(color int32, message ...string) {
	go func() { GlobalLogMessageChan <- &LogMessage{Type: 1, Message: message, Color: color} }()
}

//带日期的
func PutTimeMessage(message ...string) {
	go func() { GlobalLogMessageChan <- &LogMessage{Type: 2, Message: message} }()
}

//代理日志
func PutLogsProxyTime(message ...string) {
	if ShowProxyLog {
		go func() { GlobalLogMessageChan <- &LogMessage{Type: 3, Message: message} }()
	}
}

//普通日志
func PutLogsStaticTime(message ...string) {
	if ShowStaticLog {
		go func() { GlobalLogMessageChan <- &LogMessage{Type: 4, Message: message} }()
	}
}
