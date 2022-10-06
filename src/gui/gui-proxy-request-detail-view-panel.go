package gui

import (
	"encoding/json"
	"fmt"
	"gitee.com/snxamdf/http-server/src/consts"
	"gitee.com/snxamdf/http-server/src/entity"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//代理详情查看Panel
type RequestDetailViewPanel struct {
	TPanel                 *lcl.TPanel
	IdEdit                 *lcl.TLabeledEdit
	MethodComboBox         *lcl.TComboBox
	HostEdit               *lcl.TLabeledEdit
	SourceEdit             *lcl.TLabeledEdit
	TargetEdit             *lcl.TLabeledEdit
	RequestDetailViewMemo  *lcl.TMemo
	ResponseDetailViewMemo *lcl.TMemo
}

//代理详情查看PanelUI
func (m *RequestDetailViewPanel) initUI() {
	var enable = true
	//请求源地址
	resetPVars()
	//ID
	m.IdEdit = lcl.NewLabeledEdit(m.TPanel)
	m.IdEdit.SetParent(m.TPanel)
	m.IdEdit.SetLabelPosition(types.LpLeft)
	m.IdEdit.EditLabel().SetCaption("序号")
	m.IdEdit.SetBounds(pLeft, pTop, pWidth, pHeight)
	m.IdEdit.SetReadOnly(true)
	m.IdEdit.SetEnabled(false)

	//请求方法comboBox
	pLeft = m.IdEdit.Left() + m.IdEdit.Width() + bPLeft
	label := lcl.NewLabel(m.TPanel)
	label.SetParent(m.TPanel)
	label.SetCaption("请求方法")
	label.SetBounds(pLeft-52, pTop+5, pWidth, pHeight)
	m.MethodComboBox = lcl.NewComboBox(m.TPanel)
	m.MethodComboBox.SetParent(m.TPanel)
	m.MethodComboBox.SetBounds(pLeft, pTop, pWidth, pHeight)
	for _, method := range consts.HttpMethods {
		m.MethodComboBox.Items().Add(method)
	}
	m.MethodComboBox.SetItemIndex(0)
	m.MethodComboBox.SetOnChange(func(sender lcl.IObject) {
		fmt.Println(m.MethodComboBox.ItemIndex(), consts.HttpMethods[m.MethodComboBox.ItemIndex()], consts.GetHttpMethodsIdx(consts.HttpMethods[m.MethodComboBox.ItemIndex()]))
	})
	m.MethodComboBox.SetEnabled(enable)

	//请求Host
	pLeft = m.MethodComboBox.Left() + m.MethodComboBox.Width() + bPLeft - 20
	pWidth = pWidth * 2
	m.HostEdit = lcl.NewLabeledEdit(m.TPanel)
	m.HostEdit.SetParent(m.TPanel)
	m.HostEdit.SetLabelPosition(types.LpLeft)
	m.HostEdit.EditLabel().SetCaption("HOST")
	m.HostEdit.SetBounds(pLeft, pTop, pWidth, pHeight)
	m.HostEdit.SetEnabled(enable)

	//请求源地址
	resetPVars()
	pTop = pTop + pHeight + 10
	pWidth = pWidth * 6
	m.SourceEdit = lcl.NewLabeledEdit(m.TPanel)
	m.SourceEdit.SetParent(m.TPanel)
	m.SourceEdit.SetLabelPosition(types.LpLeft)
	m.SourceEdit.EditLabel().SetCaption("源地址")
	m.SourceEdit.SetBounds(pLeft, pTop, pWidth, pHeight)
	m.SourceEdit.SetEnabled(enable)

	//请求代理目标地址
	pTop = m.SourceEdit.Top() + pHeight + 10
	m.TargetEdit = lcl.NewLabeledEdit(m.TPanel)
	m.TargetEdit.SetParent(m.TPanel)
	m.TargetEdit.SetLabelPosition(types.LpLeft)
	m.TargetEdit.EditLabel().SetCaption("目标地址")
	m.TargetEdit.SetBounds(pLeft, pTop, pWidth, pHeight)
	m.TargetEdit.SetEnabled(enable)

	//详情查看 memo 展示格式化json串
	resetPVars()
	pLeft = 0
	pTop = m.TargetEdit.Top() + m.TargetEdit.Height() + 10
	pWidth = m.TPanel.Width()
	pHeight = m.TPanel.Height() - pTop
	var pageControl = lcl.NewPageControl(m.TPanel)
	pageControl.SetParent(m.TPanel)
	pageControl.SetBounds(pLeft, pTop, pWidth, pHeight)
	pageControl.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	var requestSheet = lcl.NewTabSheet(pageControl) //标签页
	requestSheet.SetPageControl(pageControl)
	requestSheet.SetCaption("　Request View JSON　")
	requestSheet.SetAlign(types.AlClient)
	var responseSheet = lcl.NewTabSheet(pageControl) //标签页
	responseSheet.SetPageControl(pageControl)
	responseSheet.SetCaption("　Response View JSON　")
	responseSheet.SetAlign(types.AlClient)
	//request memo
	pLeft = 0
	pTop = 0
	pWidth = m.TPanel.Width() - 10
	pHeight = pageControl.Height() - 30
	m.RequestDetailViewMemo = lcl.NewMemo(requestSheet)
	m.RequestDetailViewMemo.SetParent(requestSheet)
	m.RequestDetailViewMemo.SetScrollBars(types.SsAutoBoth)
	m.RequestDetailViewMemo.SetBounds(pLeft, pTop, pWidth, pHeight)
	m.RequestDetailViewMemo.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	m.RequestDetailViewMemo.SetBorderStyle(types.BsNone)

	//response memo
	pLeft = 0
	pTop = 0
	pWidth = m.TPanel.Width() - 10
	pHeight = pageControl.Height() - 30
	m.ResponseDetailViewMemo = lcl.NewMemo(responseSheet)
	m.ResponseDetailViewMemo.SetParent(responseSheet)
	m.ResponseDetailViewMemo.SetScrollBars(types.SsAutoBoth)
	m.ResponseDetailViewMemo.SetBounds(pLeft, pTop, pWidth, pHeight)
	m.ResponseDetailViewMemo.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
	m.ResponseDetailViewMemo.SetBorderStyle(types.BsNone)
}

//更新请求标签UI
func (m *RequestDetailViewPanel) updateRequestSheet(proxyDetail *entity.ProxyDetail) {
	m.IdEdit.SetText(fmt.Sprintf("%v", proxyDetail.ID))
	m.HostEdit.SetText(proxyDetail.Host)
	m.MethodComboBox.SetItemIndex(int32(consts.GetHttpMethodsIdx(proxyDetail.Method)))
	m.SourceEdit.SetText(proxyDetail.SourceUrl)
	m.TargetEdit.SetText(proxyDetail.TargetUrl)
	if jsn, err := json.MarshalIndent(proxyDetail.Request, "", "\t"); err == nil {
		m.RequestDetailViewMemo.SetText(string(jsn))
	}
}

//更新响应标签UI
func (m *RequestDetailViewPanel) updateResponseSheet(proxyDetail *entity.ProxyDetail) {
	if jsn, err := json.MarshalIndent(proxyDetail.Response, "", "\t"); err == nil {
		m.ResponseDetailViewMemo.SetText(string(jsn))
	}
}
