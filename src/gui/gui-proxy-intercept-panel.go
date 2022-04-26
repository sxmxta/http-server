package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
)

//代理拦截配置Panel
type ProxyInterceptPanel struct {
	TPanel                      *lcl.TPanel
	StateLabel                  *lcl.TLabel                  //拦截状态
	ProxyInterceptRequestPanel  *ProxyInterceptRequestPanel  //代理拦截请求Panel
	ProxyInterceptResponsePanel *ProxyInterceptResponsePanel //代理拦截响应Panel
	ProxyInterceptSettingPanel  *ProxyInterceptSettingPanel  //代理拦截配置Panel
}

//代理拦截请求Panel
type ProxyInterceptRequestPanel struct {
	TPanel *lcl.TPanel
}

//代理拦截响应Panel
type ProxyInterceptResponsePanel struct {
	TPanel *lcl.TPanel
}

//代理拦截配置Panel
type ProxyInterceptSettingPanel struct {
	TPanel *lcl.TPanel
}

//request
func (m *ProxyInterceptRequestPanel) initUI() {
	resetVars()
	left = 0
	top = 0
	width = m.TPanel.Width()
	height = m.TPanel.Height()

	reqPageControl := lcl.NewPageControl(m.TPanel) //Tabs 的控制标签
	reqPageControl.SetParent(m.TPanel)
	reqPageControl.SetBounds(left, top, width, height)
	reqPageControl.SetAlign(types.AlClient)

	sheet := lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　Request Query Params　")
	sheet.SetAlign(types.AlClient)
	paramsPanel := lcl.NewPanel(m.TPanel) // 标签页
	paramsPanel.SetParent(sheet)
	paramsPanel.SetBounds(0, 0, width, height)
	paramsPanel.SetAlign(types.AlClient)

	sheet = lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　Request Headers　")
	sheet.SetAlign(types.AlClient)
	headersPanel := lcl.NewPanel(m.TPanel) // 标签页
	headersPanel.SetParent(sheet)
	headersPanel.SetBounds(0, 0, width, height)
	headersPanel.SetAlign(types.AlClient)

	sheet = lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　Request Body　")
	sheet.SetAlign(types.AlClient)
	bodyPanel := lcl.NewPanel(m.TPanel) // 标签页
	bodyPanel.SetParent(sheet)
	bodyPanel.SetBounds(0, 0, width, height)
	bodyPanel.SetAlign(types.AlClient)
}

//response
func (m *ProxyInterceptResponsePanel) initUI() {
	resetVars()
	left = 0
	top = 25
	width = m.TPanel.Width()
	height = m.TPanel.Height()

	reqPageControl := lcl.NewPageControl(m.TPanel) //Tabs 的控制标签
	reqPageControl.SetParent(m.TPanel)
	reqPageControl.SetBounds(left, top, width, height)
	reqPageControl.SetAlign(types.AlClient)

	sheet := lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　Response Headers　")
	sheet.SetAlign(types.AlClient)
	headersPanel := lcl.NewPanel(m.TPanel) // 标签页
	headersPanel.SetParent(sheet)
	headersPanel.SetBounds(0, 0, width, height)
	headersPanel.SetAlign(types.AlClient)

	sheet = lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　Response Body　")
	sheet.SetAlign(types.AlClient)
	bodyPanel := lcl.NewPanel(m.TPanel) // 标签页
	bodyPanel.SetParent(sheet)
	bodyPanel.SetBounds(0, 0, width, height)
	bodyPanel.SetAlign(types.AlClient)
}

//setting
func (m *ProxyInterceptSettingPanel) initUI() {

}

//代理拦截配置Panel
func (m *ProxyInterceptPanel) initUI() {
	resetVars()
	left = 0
	top = 0
	width = m.TPanel.Width()
	height = m.TPanel.Height()

	reqPageControl := lcl.NewPageControl(m.TPanel) //Tabs 的控制标签
	reqPageControl.SetParent(m.TPanel)
	reqPageControl.SetBounds(left, top, width, height)
	reqPageControl.SetAlign(types.AlClient)

	sheet := lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　拦截请求　")
	sheet.SetAlign(types.AlClient)
	m.ProxyInterceptRequestPanel.TPanel = lcl.NewPanel(m.TPanel) //ProxyInterceptRequestPanel 标签页
	m.ProxyInterceptRequestPanel.TPanel.SetParent(sheet)
	m.ProxyInterceptRequestPanel.TPanel.SetBounds(0, 0, width, height)
	m.ProxyInterceptRequestPanel.TPanel.SetAlign(types.AlClient)

	sheet = lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　拦截响应　")
	m.ProxyInterceptResponsePanel.TPanel = lcl.NewPanel(m.TPanel) //responsePanel 标签页
	m.ProxyInterceptResponsePanel.TPanel.SetParent(sheet)
	m.ProxyInterceptResponsePanel.TPanel.SetBounds(0, 0, width, height)
	m.ProxyInterceptResponsePanel.TPanel.SetAlign(types.AlClient)

	sheet = lcl.NewTabSheet(reqPageControl) //标签页
	sheet.SetPageControl(reqPageControl)
	sheet.SetCaption("　拦截配置　")
	m.ProxyInterceptSettingPanel.TPanel = lcl.NewPanel(m.TPanel) //responsePanel 标签页
	m.ProxyInterceptSettingPanel.TPanel.SetParent(sheet)
	m.ProxyInterceptSettingPanel.TPanel.SetBounds(0, 0, width, height)
	m.ProxyInterceptSettingPanel.TPanel.SetAlign(types.AlClient)

	//初始化子组件
	m.ProxyInterceptRequestPanel.initUI()
	m.ProxyInterceptResponsePanel.initUI()
	m.ProxyInterceptSettingPanel.initUI()
}
