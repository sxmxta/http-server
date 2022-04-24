package consts

var (
	GlobalPanicRecoverString string
	GlobalMessageChan        = make(chan string)
)

func PutMessage(message string) {
	go func() { GlobalMessageChan <- message }()
}
