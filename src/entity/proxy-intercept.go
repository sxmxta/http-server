package entity

import "gitee.com/snxamdf/http-server/src/consts"

//代理拦截配置
type ProxyInterceptConfig struct {
	Index        int32
	Option       consts.PIOption
	interceptUrl string //拦截的URL
	enable       bool   //是否启用
}

func (m *ProxyInterceptConfig) Enable() bool {
	return m.enable
}
func (m *ProxyInterceptConfig) SetEnable(enable bool) {
	m.enable = enable
}

func (m *ProxyInterceptConfig) InterceptUrl() string {
	return m.interceptUrl
}

func (m *ProxyInterceptConfig) SetInterceptUrl(interceptUrl string) {
	m.interceptUrl = interceptUrl
}
