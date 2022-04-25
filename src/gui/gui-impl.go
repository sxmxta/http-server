package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/messages"
	"math"
	"strings"
	"time"
)

func (m *TGUIForm) impl() {
	m.logs = lcl.NewRichEdit(m)
	m.logs.SetParent(m)
	m.logs.Font().SetSize(10)
	m.logs.SetWidth(m.width)
	m.logs.SetHeight(m.height - 20)
	m.logs.SetTop(20)
	m.logs.SetScrollBars(types.SsAutoBoth)
	m.logs.SetReadOnly(true)
	m.logs.SetOnDblClick(func(sender lcl.IObject) {
		logsLength = 0
		m.logs.Lines().Clear()
	})
	m.logs.SetHint("双击清空")
	m.logs.SetShowHint(true)
	m.logs.SetVisible(false)

	//代理
	m.proxyGrid()
	// 底部状态条
	m.stbar = lcl.NewStatusBar(m)
	m.stbar.SetParent(m)
	m.stbar.SetAutoHint(true)
	m.stbar.SetSimplePanel(true)

	m.showProxyLogChkBox = lcl.NewCheckBox(m)
	m.showProxyLogChkBox.SetParent(m)
	m.showProxyLogChkBox.SetCaption("显示代理请求日志")
	m.showProxyLogChkBox.SetBounds(10, 0, 0, 0)
	m.showProxyLogChkBox.SetAnchors(types.NewSet(types.AkTop, types.AkLeft))
	m.showProxyLogChkBox.SetOnClick(func(sender lcl.IObject) {
		m.ShowProxyLog = m.showProxyLogChkBox.Checked()
	})
	m.showProxyLogChkBox.SetChecked(true)
	m.ShowProxyLog = true

	m.enableProxyDetailChkBox = lcl.NewCheckBox(m)
	m.enableProxyDetailChkBox.SetParent(m)
	m.enableProxyDetailChkBox.SetCaption("启用代理详情")
	m.enableProxyDetailChkBox.SetBounds(m.showProxyLogChkBox.Left()+130, 0, 0, 0)
	m.enableProxyDetailChkBox.SetAnchors(types.NewSet(types.AkTop, types.AkLeft))
	m.enableProxyDetailChkBox.SetOnClick(func(sender lcl.IObject) {
		m.EnableProxyDetail = m.enableProxyDetailChkBox.Checked()
		if m.EnableProxyDetail {
			m.SetHeight(m.Height() + 400)
		} else {
			m.SetHeight(m.Height() - 400)
		}
		m.proxyLogsGrid.SetVisible(m.EnableProxyDetail)

	})

	m.showStaticLogChkBox = lcl.NewCheckBox(m)
	m.showStaticLogChkBox.SetParent(m)
	m.showStaticLogChkBox.SetCaption("显示普通请求日志")
	m.showStaticLogChkBox.SetBounds(m.enableProxyDetailChkBox.Left()+110, 0, 0, 0)
	m.showStaticLogChkBox.SetAnchors(types.NewSet(types.AkTop, types.AkLeft))
	m.showStaticLogChkBox.SetOnClick(func(sender lcl.IObject) {
		m.ShowStaticLog = m.showStaticLogChkBox.Checked()
	})
	m.showStaticLogChkBox.SetChecked(true)
	m.ShowStaticLog = true

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

var logsLength int

func LogsColor(color int32, message string) {
	lcl.ThreadSync(func() {
		if color >= 0 {
			GUIForm.logs.SetSelStart(int32(logsLength))
			GUIForm.logs.SetSelLength(int32(strings.Count(message, "")))
			GUIForm.logs.SelAttributes().SetColor(types.TColor(uint32(color)))
		}
		GUIForm.logs.Lines().Add(message)
		GUIForm.logs.Perform(messages.EM_SCROLLCARET, 7, 0)
	})
	logsLength += strings.Count(message, "")
	if logsLength >= math.MaxInt32 || logsLength < 0 {
		logsLength = 0
	}
}

func Logs(message ...string) {
	msg := ""
	for _, v := range message {
		msg += v
	}
	LogsColor(-1, msg)
}

func LogsStaticTime(message ...string) {
	if GUIForm.ShowStaticLog {
		LogsTime(message...)
	}
}

func LogsProxyTime(message ...string) {
	if GUIForm.ShowProxyLog {
		LogsTime(message...)
	}
}

func LogsTime(message ...string) {
	go func() {
		t := time.Now()
		msg := t.Format("2006-01-02 15:04:05") + " "
		for _, v := range message {
			msg += v + " "
		}
		LogsColor(-1, msg)
	}()
}
