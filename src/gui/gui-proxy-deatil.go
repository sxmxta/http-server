package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
)

//详情使用的基础变量
var (
	bPTop    int32 = 20
	bPLeft   int32 = 70
	bPWidth  int32 = 80
	bPHeight int32 = 25
	pTop           = bPTop
	pLeft          = bPLeft
	pWidth         = bPWidth
	pHeight        = bPHeight
)

//重置基础变量
func resetPVars() {
	pLeft = bPLeft
	pTop = bPTop
	pHeight = bPHeight
	pWidth = bPWidth
}

//代理Panel
type ProxyDetailPanel struct {
	TPanel                    *lcl.TPanel
	RequestDetailViewPanel    *RequestDetailViewPanel //代理详情查看 Sheet Panel
	ProxyInterceptConfigPanel *ProxyInterceptPanel    //代理拦截配置 Sheet Panel
}

//初始化子组件对象
func (m *ProxyDetailPanel) init() {
	m.RequestDetailViewPanel = &RequestDetailViewPanel{}
	m.ProxyInterceptConfigPanel = &ProxyInterceptPanel{
		ProxyInterceptRequestPanel:  &ProxyInterceptRequestPanel{ParamsRow: 1, HeadersRow: 1},
		ProxyInterceptResponsePanel: &ProxyInterceptResponsePanel{},
		ProxyInterceptSettingPanel:  &ProxyInterceptSettingPanel{},
	}
}

func (m *TGUIForm) proxyDetailPanelInit() {
	m.ProxyDetailUI = &ProxyDetailPanel{}
	m.ProxyDetailUI.TPanel = lcl.NewPanel(m)
	m.ProxyDetailUI.TPanel.SetParent(m)
	m.ProxyDetailUI.TPanel.SetBounds(uiWidth, 0, uiWidth, uiHeight+uiHeightEx)
	m.ProxyDetailUI.TPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))

	//初始化子组件对象
	m.ProxyDetailUI.init()

	//请求响应tabs标签
	resetPVars()
	pLeft = 0
	pTop = 0
	pWidth = m.ProxyDetailUI.TPanel.Width()
	pHeight = m.ProxyDetailUI.TPanel.Height()
	m.ProxyDetailUI.proxyPages(pLeft, pTop, pWidth, pHeight)

	//代理详情查看PanelUI
	m.ProxyDetailUI.RequestDetailViewPanel.initUI()
	//代理详情查看PanelUI
	m.ProxyDetailUI.ProxyInterceptConfigPanel.initUI()

}

//UI右侧请求响应sheet Page
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
