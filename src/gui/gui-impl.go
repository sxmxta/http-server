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
	//m.logs.SetWidth(m.width)
	//m.logs.SetHeight(m.height - 25)

	//label := lcl.NewLabel(m)
	//label.SetParent(m)
	//label.SetCaption("运行时间：")
	//label.SetBounds(m.width/3, m.height-20, 0, 0)
	//label.Font().SetColor(colors.ClBlue)
	//label.Hide()

	trayIcon := lcl.NewTrayIcon(m)
	trayIcon.SetHint(m.Caption())
	trayIcon.SetVisible(true)

	var c = true
	lcl.Application.SetOnMinimize(func(sender lcl.IObject) {
		m.showHIde()
		if c {
			trayIcon.SetBalloonTitle("隐藏托盘提示")
			trayIcon.SetBalloonTimeout(5000)
			trayIcon.SetBalloonHint("我隐藏在托盘拉，单击我显示。")
			trayIcon.ShowBalloonHint()
			c = false
		}
	})
	trayIcon.SetOnClick(func(sender lcl.IObject) {
		m.showHIde()
	})
	pm := lcl.NewPopupMenu(m)
	item := lcl.NewMenuItem(m)
	item.SetCaption("显　示")
	item.SetOnClick(func(lcl.IObject) {
		m.showHIde()
	})
	pm.Items().Add(item)

	item = lcl.NewMenuItem(m)
	item.SetCaption("退　出")
	item.SetOnClick(func(lcl.IObject) {
		m.Close()
	})
	pm.Items().Add(item)
	trayIcon.SetPopupMenu(pm)
}

var b = true

func (m *TGUIForm) showHIde() {
	b = !b
	if b {
		m.SetWindowState(types.WsNormal)
		m.Show()
	} else {
		m.Hide()
	}
}

func Logs(message ...string) {
	lcl.ThreadSync(func() {
		msg := ""
		for _, v := range message {
			msg += v
		}
		GUIForm.logs.Lines().Add(msg)
		GUIForm.logs.Perform(messages.EM_SCROLLCARET, 7, 0)
	})
}
func LogsTime(message ...string) {
	go func() {
		t := time.Now()
		msg := t.Format("2006-01-02 15:04:05") + " "
		for _, v := range message {
			msg += v + " "
		}
		Logs(msg)
	}()
}
