package gui

import (
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/http-server/src/consts"
)

//代理详情查看Panel
type RequestDetailViewPanel struct {
	TPanel         *lcl.TPanel
	IdEdit         *lcl.TLabeledEdit
	MethodComboBox *lcl.TComboBox
	HostEdit       *lcl.TLabeledEdit
	SourceEdit     *lcl.TLabeledEdit
	TargetEdit     *lcl.TLabeledEdit
}

//代理详情查看PanelUI
func (m *RequestDetailViewPanel) initUI() {
	//请求源地址
	resetVars()
	//ID
	m.IdEdit = lcl.NewLabeledEdit(m.TPanel)
	m.IdEdit.SetParent(m.TPanel)
	m.IdEdit.SetLabelPosition(types.LpLeft)
	m.IdEdit.EditLabel().SetCaption("序号")
	m.IdEdit.SetBounds(left, top, width, height)
	m.IdEdit.SetReadOnly(true)
	m.IdEdit.SetEnabled(false)

	//请求方法comboBox
	left = m.IdEdit.Left() + m.IdEdit.Width() + bLeft
	label := lcl.NewLabel(m.TPanel)
	label.SetParent(m.TPanel)
	label.SetCaption("请求方法")
	label.SetBounds(left-50, top+5, width, height)
	m.MethodComboBox = lcl.NewComboBox(m.TPanel)
	m.MethodComboBox.SetParent(m.TPanel)
	m.MethodComboBox.SetBounds(left, top, width, height)
	for _, method := range consts.HttpMethods {
		m.MethodComboBox.Items().Add(method)
	}
	m.MethodComboBox.SetItemIndex(0)
	m.MethodComboBox.SetOnChange(func(sender lcl.IObject) {
		fmt.Println(m.MethodComboBox.ItemIndex(), consts.HttpMethods[m.MethodComboBox.ItemIndex()], consts.GetHttpMethodsIdx(consts.HttpMethods[m.MethodComboBox.ItemIndex()]))
	})

	//请求Host
	left = m.MethodComboBox.Left() + m.MethodComboBox.Width() + bLeft
	width = width * 2
	m.HostEdit = lcl.NewLabeledEdit(m.TPanel)
	m.HostEdit.SetParent(m.TPanel)
	m.HostEdit.SetLabelPosition(types.LpLeft)
	m.HostEdit.EditLabel().SetCaption("HOST")
	m.HostEdit.SetBounds(left, top, width, height)

	//请求源地址
	resetVars()
	top = top + height + 10
	width = width * 6
	m.SourceEdit = lcl.NewLabeledEdit(m.TPanel)
	m.SourceEdit.SetParent(m.TPanel)
	m.SourceEdit.SetLabelPosition(types.LpLeft)
	m.SourceEdit.EditLabel().SetCaption("源地址")
	m.SourceEdit.SetBounds(left, top, width, height)

	//请求代理目标地址
	top = m.SourceEdit.Top() + height + 10
	m.TargetEdit = lcl.NewLabeledEdit(m.TPanel)
	m.TargetEdit.SetParent(m.TPanel)
	m.TargetEdit.SetLabelPosition(types.LpLeft)
	m.TargetEdit.EditLabel().SetCaption("目标地址")
	m.TargetEdit.SetBounds(left, top, width, height)
}
