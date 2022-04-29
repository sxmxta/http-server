package entity

var (
	AppInitSuccess       = true                    //app初始化成功结果
	ShowProxyLog         bool                      //显示代理日志开关
	ShowStaticLog        bool                      //显示普通日志开关
	EnableProxyDetail    bool                      //启用代理详情开关
	GlobalLogMessageChan = make(chan *LogMessage)  //全局日志输出通道
	ProxyDetailChan      = make(chan *ProxyDetail) //代理详情数据传输通道
	ProxyInterceptChan   = make(chan *ProxyDetail) //代理拦截数据传输通道
	ProxyInterceptEnable = false                   //代理拦截启用开关
)
