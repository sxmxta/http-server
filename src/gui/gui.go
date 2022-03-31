package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
)

var GUIForm = &TGUIForm{}

type TGUIForm struct {
	*lcl.TForm
	width     int32
	height    int32
	logs      *lcl.TMemo
	scrollBar *lcl.TScrollBar
}

func (m *TGUIForm) OnFormCreate(sender lcl.IObject) {
	m.init()
	m.SetCaption("Web Server")
	m.SetPosition(types.PoScreenCenter)
	m.EnabledMaximize(false)
	m.SetWidth(m.width)
	m.SetHeight(m.height)
	m.SetBorderStyle(types.BsSingle)
	m.impl()
}

func (m *TGUIForm) init() {
	m.width = 600
	m.height = 400
	icon := lcl.NewIcon()
	icon.LoadFromFSFile("resources/app.ico")
	m.SetIcon(icon)
}
