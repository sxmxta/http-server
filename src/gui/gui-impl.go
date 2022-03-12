package gui

import (
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/rtl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/messages"
	"gitee.com/snxamdf/golcl/lcl/win"
	"time"
)

func (m *TGUIForm) impl() {
	m.logs = lcl.NewMemo(m)
	m.logs.SetParent(m)
	m.logs.Font().SetSize(10)
	m.logs.SetAlign(types.AlClient)
	m.logs.SetScrollBars(types.SsAutoBoth)
	//m.logs.SetReadOnly(true)
	m.ScrollInView(m.logs)
	m.scrollBar = lcl.NewScrollBar(m)
	m.scrollBar.SetParent(m)
	m.scrollBar.SetHeight(500)
	m.ScrollInView(m.scrollBar)
	m.VertScrollBar().SetVisible(true)
	m.SetAutoScroll(true)
}

func Logs(message ...string) {
	msg := ""
	for _, v := range message {
		msg += v
	}
	GUIForm.logs.Lines().Add(msg)
}
func LogsTime(message ...string) {
	t := time.Now()
	msg := t.Format("2006-01-02 15:04:05") + " "
	for _, v := range message {
		msg += v
	}
	GUIForm.logs.Lines().Add(msg)
	rr := GUIForm.logs.Perform(messages.EM_SCROLLCARET, 7, 0)
	fmt.Println(rr)
	r := rtl.SendMessage(GUIForm.logs.Handle(), messages.EM_SCROLL, win.SB_BOTTOM, 0)
	fmt.Println(r)
	//GUIForm.logs.SetSelLength()

	//rtl.SendMessage(GUIForm.logs.Handle(), WM_VSCROLL, SB_BOTTOM, 0)
}
