package entity

import "sync/atomic"

var (
	AppInitSuccess             = true                                //app初始化成功结果
	ShowProxyLog               bool                                  //显示代理日志开关
	ShowStaticLog              bool                                  //显示普通日志开关
	EnableProxyDetail          bool                                  //启用代理详情开关 checkbox
	GlobalLogMessageChan       = make(chan *LogMessage)              //http-server for select: 全局日志输出通道
	ProxyDetailGridChan        = make(chan *ProxyDetail)             //gui 组件 for select: 代理列表和详情数据传输通道
	ProxyFlowInterceptChan     = make(chan *ProxyDetail)             //gui 组件 for select: 代理拦截流程数据传输通道 proxy-server > gui，请求同步到UI
	ProxyInterceptConfigChan   = make(chan *[]*ProxyInterceptConfig) //proxy server for select: 代理拦截配置数据传输通道 gui > proxy-server，拦截规则同步给proxy server处理
	ProxyInterceptConfigEnable = true                                //代理拦截启用开关 button
	ID                         = &id{}
)

type id struct {
	value int32
}

func (m *id) Get() int32 {
	atomic.AddInt32(&m.value, 1)
	return atomic.LoadInt32(&m.value)
}

func (m *id) Reset() {
	atomic.StoreInt32(&m.value, 0)
}
