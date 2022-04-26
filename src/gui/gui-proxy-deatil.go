package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
)

var (
	bTop    int32 = 20
	bLeft   int32 = 70
	bWidth  int32 = 80
	bHeight int32 = 25
	top           = bTop
	left          = bLeft
	width         = bWidth
	height        = bHeight
)

func resetVars() {
	left = bLeft
	top = bTop
	height = bHeight
	width = bWidth
}

type ProxyDetailPanel struct {
	TPanel                    *lcl.TPanel
	IdEdit                    *lcl.TLabeledEdit
	MethodComboBox            *lcl.TComboBox
	HostEdit                  *lcl.TLabeledEdit
	SourceEdit                *lcl.TLabeledEdit
	TargetEdit                *lcl.TLabeledEdit
	RequestDetailViewPanel    *RequestDetailViewPanel //代理详情查看
	ProxyInterceptConfigPanel *ProxyInterceptPanel    //代理拦截配置Panel
}

//初始化子组件对象
func (m *ProxyDetailPanel) init() {
	m.RequestDetailViewPanel = &RequestDetailViewPanel{}
	m.ProxyInterceptConfigPanel = &ProxyInterceptPanel{
		ProxyInterceptRequestPanel:  &ProxyInterceptRequestPanel{},
		ProxyInterceptResponsePanel: &ProxyInterceptResponsePanel{},
		ProxyInterceptSettingPanel:  &ProxyInterceptSettingPanel{},
	}
}

func (m *TGUIForm) proxyDetailPanelInit() {
	m.ProxyDetailUI = &ProxyDetailPanel{}
	m.ProxyDetailUI.TPanel = lcl.NewPanel(m)
	m.ProxyDetailUI.TPanel.SetParent(m)
	m.ProxyDetailUI.TPanel.SetBounds(m.width, 0, m.width, m.height+400)
	m.ProxyDetailUI.TPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))

	//初始化子组件对象
	m.ProxyDetailUI.init()

	//请求响应tabs标签
	resetVars()
	left = 0
	top = 0
	width = m.ProxyDetailUI.TPanel.Width()
	height = m.ProxyDetailUI.TPanel.Height()
	m.ProxyDetailUI.proxyPages(left, top, width, height)

	//代理详情查看PanelUI
	m.ProxyDetailUI.RequestDetailViewPanel.initUI()
	//代理详情查看PanelUI
	m.ProxyDetailUI.ProxyInterceptConfigPanel.initUI()

}

//请求响应Page
func (m *ProxyDetailPanel) proxyPages(left, top, width, height int32) {
	pagePanel := lcl.NewPanel(m.TPanel) //创建一个tabs的父组件，可以根据客户端变更大小
	pagePanel.SetParent(m.TPanel)
	pagePanel.SetBounds(left, top, width, height)
	pagePanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	pagePanel.SetBevelOuter(types.BvNone) //去除panel边框

	pageControl := lcl.NewPageControl(pagePanel) //Tabs 的控制标签
	pageControl.SetParent(pagePanel)
	pageControl.SetBounds(left, top, width, height)
	pageControl.SetAlign(types.AlClient)

	sheet := lcl.NewTabSheet(pagePanel) //标签页
	sheet.SetPageControl(pageControl)
	sheet.SetCaption("　代理详情查看　")
	sheet.SetAlign(types.AlClient)
	m.RequestDetailViewPanel.TPanel = lcl.NewPanel(pagePanel) //ProxyInterceptRequestPanel 标签页
	m.RequestDetailViewPanel.TPanel.SetParent(sheet)
	m.RequestDetailViewPanel.TPanel.SetBounds(0, 0, width, height)
	m.RequestDetailViewPanel.TPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))

	sheet = lcl.NewTabSheet(pagePanel) //标签页
	sheet.SetPageControl(pageControl)
	sheet.SetCaption("　代理拦截　")
	m.ProxyInterceptConfigPanel.TPanel = lcl.NewPanel(pagePanel) //responsePanel 标签页
	m.ProxyInterceptConfigPanel.TPanel.SetParent(sheet)
	m.ProxyInterceptConfigPanel.TPanel.SetBounds(0, 0, width, height)
	m.ProxyInterceptConfigPanel.TPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
}
