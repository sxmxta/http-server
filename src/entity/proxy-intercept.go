package entity

import "sync"

//锁
var PIC = &pic{
	proxyConfigLock: sync.RWMutex{},
	proxyConfig:     map[string]*proxyInterceptConfig{},
}

type pic struct {
	proxyConfigLock sync.RWMutex
	proxyConfig     map[string]*proxyInterceptConfig
}

//代理拦截配置
type proxyInterceptConfig struct {
	interceptUrl string //拦截的URL
	enable       bool   //是否启用
}

func (m *proxyInterceptConfig) Enable() bool {
	return m.enable
}

func (m *proxyInterceptConfig) InterceptUrl() string {
	return m.interceptUrl
}

//添加一个拦截配置
func (m *pic) AddProxyInterceptConfig(interceptUrl string) *proxyInterceptConfig {
	m.proxyConfigLock.Lock()
	defer m.proxyConfigLock.Unlock()
	if pc, ok := m.proxyConfig[interceptUrl]; ok {
		pc.interceptUrl = interceptUrl
		return pc
	} else {
		pc = &proxyInterceptConfig{interceptUrl: interceptUrl, enable: true}
		m.proxyConfig[interceptUrl] = pc
		return pc
	}
}

//添加一个拦截配置
func (m *pic) GetProxyInterceptConfig(interceptUrl string) *proxyInterceptConfig {
	m.proxyConfigLock.Lock()
	defer m.proxyConfigLock.Unlock()
	return m.proxyConfig[interceptUrl]
}

//删除一个拦截配置
func (m *pic) DelProxyInterceptConfig(interceptUrl string) {
	m.proxyConfigLock.Lock()
	defer m.proxyConfigLock.Unlock()
	delete(m.proxyConfig, interceptUrl)
}
