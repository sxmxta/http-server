package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/messages"
	"time"
)

func (m *TGUIForm) impl() {
	m.logs = lcl.NewMemo(m)
	m.logs.SetParent(m)
	m.logs.Font().SetSize(10)
	m.logs.SetAlign(types.AlClient)
	m.logs.SetScrollBars(types.SsAutoBoth)
}

func Logs(message ...string) {
	go func() {
		lcl.ThreadSync(func() {
			msg := ""
			for _, v := range message {
				msg += v
			}
			GUIForm.logs.Lines().Add(msg)
			GUIForm.logs.Perform(messages.EM_SCROLLCARET, 7, 0)
		})
	}()
}
func LogsTime(message ...string) {
	go func() {
		lcl.ThreadSync(func() {
			t := time.Now()
			msg := t.Format("2006-01-02 15:04:05") + " "
			for _, v := range message {
				msg += v + " "
			}
			GUIForm.logs.Lines().Add(msg)
			GUIForm.logs.Perform(messages.EM_SCROLLCARET, 7, 0)
		})
	}()
}
