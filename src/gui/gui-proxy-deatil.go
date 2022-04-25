package gui

import (
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/consts"
)

type ProxyDetailUI struct {
	DetailPanel      *lcl.TPanel
	IdEdit           *lcl.TLabeledEdit
	MethodComboBox   *lcl.TComboBox
	HostEdit         *lcl.TLabeledEdit
	SourceEdit       *lcl.TLabeledEdit
	TargetEdit       *lcl.TLabeledEdit
	Parent           *TGUIForm
	requestPanel     *lcl.TPanel
	responsePanel    *lcl.TPanel
	proxyConfigPanel *lcl.TPanel
}

func (m *TGUIForm) proxyDetailPanelInit() {
	m.ProxyDetailUI = &ProxyDetailUI{}
	m.ProxyDetailUI.Parent = m
	m.ProxyDetailUI.DetailPanel = lcl.NewPanel(m)
	m.ProxyDetailUI.DetailPanel.SetParent(m)
	m.ProxyDetailUI.DetailPanel.SetBounds(m.width, 0, m.width, m.height+400)
	m.ProxyDetailUI.DetailPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	//m.ProxyDetailPanel.SetColor(colors.ClRed)
	var (
		mLeft         = m.ProxyDetailUI.DetailPanel.Left()
		mTop          = m.ProxyDetailUI.DetailPanel.Top()
		bLeft   int32 = 70
		left          = bLeft
		bTop    int32 = 20
		top           = bTop
		bHeight int32 = 25
		height        = bHeight
		bWidth  int32 = 80
		width         = bWidth
	)
	fmt.Println(mLeft, mTop)
	var reset = func() {
		left = bLeft
		top = bTop
		height = bHeight
		width = bWidth
	}
	//ID
	m.ProxyDetailUI.IdEdit = lcl.NewLabeledEdit(m.ProxyDetailUI.DetailPanel)
	m.ProxyDetailUI.IdEdit.SetParent(m.ProxyDetailUI.DetailPanel)
	m.ProxyDetailUI.IdEdit.SetLabelPosition(types.LpLeft)
	m.ProxyDetailUI.IdEdit.EditLabel().SetCaption("RID")
	m.ProxyDetailUI.IdEdit.SetBounds(left, top, width, height)
	m.ProxyDetailUI.IdEdit.SetReadOnly(true)
	m.ProxyDetailUI.IdEdit.SetEnabled(false)

	//请求方法组件
	left = m.ProxyDetailUI.IdEdit.Left() + m.ProxyDetailUI.IdEdit.Width() + bLeft
	m.ProxyDetailUI.requestMethod(left, top, width, height)

	//请求Host
	left = m.ProxyDetailUI.MethodComboBox.Left() + m.ProxyDetailUI.MethodComboBox.Width() + bLeft
	width = width * 2
	m.ProxyDetailUI.HostEdit = lcl.NewLabeledEdit(m.ProxyDetailUI.DetailPanel)
	m.ProxyDetailUI.HostEdit.SetParent(m.ProxyDetailUI.DetailPanel)
	m.ProxyDetailUI.HostEdit.SetLabelPosition(types.LpLeft)
	m.ProxyDetailUI.HostEdit.EditLabel().SetCaption("Host")
	m.ProxyDetailUI.HostEdit.SetBounds(left, top, width, height)

	//请求源地址
	reset()
	top = top + height + 10
	width = width * 6
	m.ProxyDetailUI.SourceEdit = lcl.NewLabeledEdit(m.ProxyDetailUI.DetailPanel)
	m.ProxyDetailUI.SourceEdit.SetParent(m.ProxyDetailUI.DetailPanel)
	m.ProxyDetailUI.SourceEdit.SetLabelPosition(types.LpLeft)
	m.ProxyDetailUI.SourceEdit.EditLabel().SetCaption("SourceUrl")
	m.ProxyDetailUI.SourceEdit.SetBounds(left, top, width, height)

	//请求代理目标地址
	top = m.ProxyDetailUI.SourceEdit.Top() + height + 10
	m.ProxyDetailUI.TargetEdit = lcl.NewLabeledEdit(m.ProxyDetailUI.DetailPanel)
	m.ProxyDetailUI.TargetEdit.SetParent(m.ProxyDetailUI.DetailPanel)
	m.ProxyDetailUI.TargetEdit.SetLabelPosition(types.LpLeft)
	m.ProxyDetailUI.TargetEdit.EditLabel().SetCaption("TargetUrl")
	m.ProxyDetailUI.TargetEdit.SetBounds(left, top, width, height)

	//请求响应tabs标签
	reset()
	left = 0
	top = m.ProxyDetailUI.TargetEdit.Top() + m.ProxyDetailUI.TargetEdit.Height() + 10
	width = m.width
	height = 650
	m.ProxyDetailUI.requestResponsePage(left, top, width, height)

	//请求sheet
	m.ProxyDetailUI.requestSheet()
	//响应sheet
	m.ProxyDetailUI.responseSheet()
	//代理拦截配置sheet
	m.ProxyDetailUI.proxyInterceptSheet()
}

//请求sheet
func (m *ProxyDetailUI) requestSheet() {

}

//响应sheet
func (m *ProxyDetailUI) responseSheet() {

}

//代理拦截配置sheet
func (m *ProxyDetailUI) proxyInterceptSheet() {

}

//请求响应Page
func (m *ProxyDetailUI) requestResponsePage(left int32, top int32, width int32, height int32) {
	pagePanel := lcl.NewPanel(m.DetailPanel) //创建一个tabs的父组件，可以根据客户端变更大小
	pagePanel.SetParent(m.DetailPanel)
	pagePanel.SetBounds(left, top, width, height)
	pagePanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	pagePanel.SetBevelOuter(types.BvNone) //去除panel边框

	pageControl := lcl.NewPageControl(pagePanel) //Tabs 的控制标签
	pageControl.SetParent(pagePanel)
	pageControl.SetBounds(left, top, width, height)
	pageControl.SetAlign(types.AlClient)
	pageControl.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))

	sheet := lcl.NewTabSheet(pagePanel) //标签页
	sheet.SetPageControl(pageControl)
	sheet.SetCaption(" 请求信息-request")
	font := lcl.NewFont()
	font.SetColor(colors.ClGreen)
	sheet.SetFont(font)
	sheet.SetAlign(types.AlClient)
	sheet.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	m.requestPanel = lcl.NewPanel(pagePanel) //requestPanel 标签页
	m.requestPanel.SetParent(sheet)
	m.requestPanel.SetBounds(0, 0, width, height)
	m.requestPanel.SetAlignment(types.AlClient)
	m.requestPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))

	sheet = lcl.NewTabSheet(pagePanel) //标签页
	sheet.SetPageControl(pageControl)
	sheet.SetCaption(" 响应信息-response")
	sheet.SetShowHint(true)
	m.responsePanel = lcl.NewPanel(pagePanel) //responsePanel 标签页
	m.responsePanel.SetParent(sheet)
	m.responsePanel.SetBounds(0, 0, width, height)
	m.responsePanel.SetAlignment(types.AlClient)
	m.responsePanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))

	sheet = lcl.NewTabSheet(pagePanel) //标签页
	sheet.SetPageControl(pageControl)
	sheet.SetCaption(" 代理拦截配置 ")
	m.proxyConfigPanel = lcl.NewPanel(pagePanel) //responsePanel 标签页
	m.proxyConfigPanel.SetParent(sheet)
	m.proxyConfigPanel.SetBounds(0, 0, width, height)
	m.proxyConfigPanel.SetAlignment(types.AlClient)
	m.proxyConfigPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
}

//请求方法下拉框
func (m *ProxyDetailUI) requestMethod(left int32, top int32, width int32, height int32) {
	label := lcl.NewLabel(m.DetailPanel)
	label.SetParent(m.DetailPanel)
	label.SetCaption("Method")
	label.SetBounds(left-50, top+5, width, height)
	m.MethodComboBox = lcl.NewComboBox(m.DetailPanel)
	m.MethodComboBox.SetParent(m.DetailPanel)
	m.MethodComboBox.SetBounds(left, top, width, height)
	for _, method := range consts.HttpMethods {
		m.MethodComboBox.Items().Add(method)
	}
	m.MethodComboBox.SetItemIndex(0)
	m.MethodComboBox.SetOnChange(func(sender lcl.IObject) {
		fmt.Println(m.MethodComboBox.ItemIndex(), consts.HttpMethods[m.MethodComboBox.ItemIndex()], consts.GetHttpMethodsIdx(consts.HttpMethods[m.MethodComboBox.ItemIndex()]))
	})
	//m.MethodComboBox.Items().Assign(lcl.Printer.Printers())
}

func (m *ProxyDetailUI) updateRequestSheet(proxyDetail *ProxyDetail) {
	m.IdEdit.SetText(fmt.Sprintf("%v", proxyDetail.ID))
	m.HostEdit.SetText(proxyDetail.Host)
	m.MethodComboBox.SetItemIndex(int32(consts.GetHttpMethodsIdx(proxyDetail.Method)))
	m.SourceEdit.SetText(proxyDetail.SourceUrl)
	m.TargetEdit.SetText(proxyDetail.TargetUrl)
}
func (m *ProxyDetailUI) updateResponseSheet(proxyDetail *ProxyDetail) {

}
