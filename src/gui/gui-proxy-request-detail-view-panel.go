package gui

import (
	"encoding/json"
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/http-server/src/consts"
	"gitee.com/snxamdf/http-server/src/entity"
)

//代理详情查看Panel
type RequestDetailViewPanel struct {
	TPanel         *lcl.TPanel
	IdEdit         *lcl.TLabeledEdit
	MethodComboBox *lcl.TComboBox
	HostEdit       *lcl.TLabeledEdit
	SourceEdit     *lcl.TLabeledEdit
	TargetEdit     *lcl.TLabeledEdit
	DetailViewMemo *lcl.TMemo
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
	label.SetBounds(pLeft-50, pTop+5, pWidth, pHeight)
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
	pLeft = m.MethodComboBox.Left() + m.MethodComboBox.Width() + bPLeft
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

	//详情查看
	resetPVars()
	pLeft = 10
	pTop = m.TargetEdit.Top() + m.TargetEdit.Height() + 10
	pWidth = m.TPanel.Width() - 20
	pHeight = m.TPanel.Height() - pTop - 30
	m.DetailViewMemo = lcl.NewMemo(m.TPanel)
	m.DetailViewMemo.SetParent(m.TPanel)
	m.DetailViewMemo.SetScrollBars(types.SsAutoBoth)
	m.DetailViewMemo.SetBounds(pLeft, pTop, pWidth, pHeight)
	m.DetailViewMemo.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkTop, types.AkRight))
}

//更新请求标签UI
func (m *RequestDetailViewPanel) updateRequestSheet(proxyDetail *entity.ProxyDetail) {
	m.IdEdit.SetText(fmt.Sprintf("%v", proxyDetail.ID))
	m.HostEdit.SetText(proxyDetail.Host)
	m.MethodComboBox.SetItemIndex(int32(consts.GetHttpMethodsIdx(proxyDetail.Method)))
	m.SourceEdit.SetText(proxyDetail.SourceUrl)
	m.TargetEdit.SetText(proxyDetail.TargetUrl)
	if jsn, err := json.MarshalIndent(proxyDetail, "", "\t"); err == nil {
		m.DetailViewMemo.SetText(string(jsn))
	}
}

//更新响应标签UI
func (m *RequestDetailViewPanel) updateResponseSheet(proxyDetail *entity.ProxyDetail) {

}
