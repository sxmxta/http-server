package entity

var (
	AppInitSuccess       = true
	ShowProxyLog         bool
	ShowStaticLog        bool
	EnableProxyDetail    bool
	GlobalLogMessageChan = make(chan *LogMessage)
	ProxyDetailChan      = make(chan *ProxyDetail)
)
