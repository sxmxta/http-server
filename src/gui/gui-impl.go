package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
)

func (m *TGUIForm) impl() {
	m.logs = lcl.NewMemo(m)
	m.logs.SetParent(m)
	m.logs.Font().SetSize(10)
	m.logs.SetAlign(types.AlClient)
	m.logs.SetScrollBars(types.SsAutoBoth)
	m.logs.SetReadOnly(true)
}

func Logs(message ...string) {
	msg := ""
	for _, v := range message {
		msg += v
	}
	GUIForm.logs.Lines().Add(msg)
}
