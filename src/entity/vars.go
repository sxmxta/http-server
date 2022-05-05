package entity

var (
	AppInitSuccess           = true                             //app初始化成功结果
	ShowProxyLog             bool                               //显示代理日志开关
	ShowStaticLog            bool                               //显示普通日志开关
	EnableProxyDetail        bool                               //启用代理详情开关
	GlobalLogMessageChan     = make(chan *LogMessage)           //http-server for select: 全局日志输出通道
	ProxyDetailChan          = make(chan *ProxyDetail)          //gui 组件 for select: 代理详情数据传输通道
	ProxyFlowInterceptChan   = make(chan *ProxyDetail)          //gui 组件 for select: 代理拦截流程数据传输通道 proxy-server > gui，请求同步到UI
	ProxyInterceptSignal     = make(chan int32)                 //代理拦截产生的信号，数字标记状态 10:开始请求拦截 11:结束请求拦截， 20:开始响应拦截 21:结束响应拦截
	ProxyInterceptConfigChan = make(chan *ProxyInterceptConfig) //proxy server for select: 代理拦截配置数据传输通道 gui > proxy-server，拦截规则同步给proxy server处理
	ProxyInterceptWaitChan   = make(chan int32)                 //
	ProxyInterceptEnable     = true                             //代理拦截启用开关
)
