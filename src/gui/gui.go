package gui

import (
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
)

var GUIForm = &TGUIForm{}

type TGUIForm struct {
	*lcl.TForm
	width  int32
	height int32
}

func (m *TGUIForm) OnFormCreate(sender lcl.IObject) {
	// init
	m.init()
	m.SetCaption("web-server")

	fmt.Println("大小：", lcl.Screen.DesktopTop(), lcl.Screen.DesktopLeft(), lcl.Screen.DesktopHeight(), lcl.Screen.DesktopWidth())
	fmt.Println("大小: ", lcl.Screen.DesktopRect())
	fmt.Println("数量: ", lcl.Screen.MonitorCount())
	fmt.Println("0： ", lcl.Screen.Monitors(0))

	lcl.Screen.Monitors(0).Instance()

	m.SetPosition(types.PoScreenCenter)
	m.EnabledMaximize(false)
	m.SetWidth(m.width)
	m.SetHeight(m.height)
	m.SetBorderStyle(types.BsSingle)
}

func (m *TGUIForm) init() {
	m.width = 600
	m.height = 400
	icon := lcl.NewIcon()
	icon.LoadFromFSFile("resources/icon.ico")
	m.SetIcon(icon)
}
