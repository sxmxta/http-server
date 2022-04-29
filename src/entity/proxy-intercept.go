package entity

type ProxyInterceptConfig struct {
	InterceptUrl string
	Enable       bool
}

func (m *ProxyInterceptConfig) SetEnable(enable bool) {
	m.Enable = enable
}
