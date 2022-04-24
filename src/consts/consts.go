package consts

var (
	AppInitSuccess = true
	GlobalMessage  = make(chan MessageChannel)
)

type MessageChannel struct {
	Type    int
	Message []string
	Color   int32
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
