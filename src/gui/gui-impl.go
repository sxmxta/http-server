package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/messages"
	"gitee.com/snxamdf/http-server/src/entity"
	"strings"
	"time"
)

func (m *TGUIForm) impl() {
	// 底部状态条
	m.stateBar = lcl.NewStatusBar(m)
	m.stateBar.SetParent(m)
	//m.stateBar.SetAutoHint(true)
	m.stateBar.SetSimplePanel(false)
	m.stateBar.Panels().Add().SetText(stateBarText)

	//---- begin 任务栏托盘 tray icon ----
	trayIcon := lcl.NewTrayIcon(m)
	trayIcon.SetHint(m.Caption())
	trayIcon.SetVisible(true)

	m.leftPanel.SetOnResize(func(sender lcl.IObject) {
		w := m.leftPanel.Width() - 3 - 200 //panel - splitter - 200
		m.proxyLogsGridColAddr.SetWidth(w)
	})

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
	//托盘右键菜单
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
	//---- end 任务栏托盘 tray icon ----

	//---- begin gui 窗口变化 ----
	//m.SetOnResize(func(sender lcl.IObject) {
	//	//var mh = m.Height()
	//	//mh = mh - uiHeight - 20
	//	//m.proxyLogsGrid.SetHeight(mh)
	//	fmt.Println("SetOnResize")
	//})
	//m.SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
	//	//var mh = m.Height()
	//	//mh = mh - uiHeight - 20
	//	//m.proxyLogsGrid.SetHeight(mh)
	//	fmt.Println("SetOnConstrainedResize")
	//})
	//---- end gui 窗口变化 ----

	//left ui
	m.initLeftUI()
}

//left ui
func (m *TGUIForm) initLeftUI() {
	m.logs = lcl.NewRichEdit(m.leftPanel)
	m.logs.SetParent(m.leftPanel)
	m.logs.Font().SetSize(10)
	m.logs.SetBounds(0, 21, uiWidth, uiHeight-21)
	m.logs.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
	m.logs.SetScrollBars(types.SsAutoBoth)
	m.logs.SetReadOnly(true)
	m.logs.SetParentColor(true)
	m.logs.SetOnDblClick(func(sender lcl.IObject) {
		clearLogs()
	})
	m.logs.SetHint("双击清空")
	m.logs.SetShowHint(true)
	//m.logs.SetVisible(false)

	//---- begin 显示代理请求日志 checkbox ----
	m.showProxyLogChkBox = lcl.NewCheckBox(m.leftPanel)
	m.showProxyLogChkBox.SetParent(m.leftPanel)
	m.showProxyLogChkBox.SetCaption("显示代理请求日志")
	m.showProxyLogChkBox.SetHint("不影响服务性能\n会增加内存消耗")
	m.showProxyLogChkBox.SetShowHint(true)
	m.showProxyLogChkBox.SetBounds(10, 0, 0, 0)
	m.showProxyLogChkBox.SetAnchors(types.NewSet(types.AkTop, types.AkLeft))
	m.showProxyLogChkBox.SetOnClick(func(sender lcl.IObject) {
		entity.ShowProxyLog = m.showProxyLogChkBox.Checked()
	})
	//m.showProxyLogChkBox.SetChecked(true)
	//entity.ShowProxyLog = true
	//---- end 显示代理请求日志 checkbox ----

	//---- begin 显示普通请求日志 checkbox ----
	m.showStaticLogChkBox = lcl.NewCheckBox(m.leftPanel)
	m.showStaticLogChkBox.SetParent(m.leftPanel)
	m.showStaticLogChkBox.SetCaption("显示普通请求日志")
	m.showStaticLogChkBox.SetHint("不影响服务性能\n会增加内存消耗")
	m.showStaticLogChkBox.SetShowHint(true)
	m.showStaticLogChkBox.SetBounds(m.showProxyLogChkBox.Left()+130, 0, 0, 0)
	m.showStaticLogChkBox.SetAnchors(types.NewSet(types.AkTop, types.AkLeft))
	m.showStaticLogChkBox.SetOnClick(func(sender lcl.IObject) {
		entity.ShowStaticLog = m.showStaticLogChkBox.Checked()
	})
	//m.showStaticLogChkBox.SetChecked(true)
	//entity.ShowStaticLog = true
	//---- end 显示普通请求日志 checkbox ----

	//---- begin 启用代理详情 checkbox ----
	m.enableProxyDetailChkBox = lcl.NewCheckBox(m.leftPanel)
	m.enableProxyDetailChkBox.SetParent(m.leftPanel)
	m.enableProxyDetailChkBox.SetCaption("启用代理跟踪详情")
	m.enableProxyDetailChkBox.SetHint("严重影响服务性能\n适用于调式接口")
	m.enableProxyDetailChkBox.SetShowHint(true)
	m.enableProxyDetailChkBox.SetBounds(m.showStaticLogChkBox.Left()+130, 0, 0, 0)
	m.enableProxyDetailChkBox.SetAnchors(types.NewSet(types.AkTop, types.AkLeft))
	//代理详情checkBox
	m.enableProxyDetailChkBox.SetOnClick(func(sender lcl.IObject) {
		entity.EnableProxyDetail = m.enableProxyDetailChkBox.Checked()
		if entity.EnableProxyDetail {
			m.SetHeight(m.Height() + uiHeightEx)
			m.SetWidth(m.Width() + uiWidthEx)
			m.SetBorderStyle(types.BsSizeable)
			m.proxyLogsGrid.SetBounds(0, uiHeight, m.leftPanel.Width(), m.Height()-uiHeight-20)
			m.proxyLogsGridColAddr.SetWidth(uiWidth - 200)
		} else {
			if m.WindowState() == types.WsMaximized {
				m.SetWindowState(types.WsNormal)
			}
			m.SetHeight(uiHeight)
			m.SetWidth(uiWidth)
			m.SetBorderStyle(types.BsSingle)
			m.leftPanel.SetWidth(uiWidth)
			m.logs.SetWidth(uiWidth)
		}
		if m.rightPanel == nil {
			m.initRightUI()
		}
		m.proxyLogsGrid.SetVisible(entity.EnableProxyDetail)
		m.rightPanel.TPanel.SetVisible(entity.EnableProxyDetail)
	})
	//---- end 启用代理详情 checkbox ----

	//代理日志列表
	m.proxyGrid()
}

func (m *TGUIForm) showHIde() {
	var b = !m.Visible()
	if b {
		m.SetWindowState(types.WsNormal)
		m.Show()
	} else {
		m.Hide()
	}
}

func LogsColor(color int32, message string) {
	lcl.ThreadSync(func() {
		if color >= 0 {
			GUIForm.logs.SetSelStart(GUIForm.logs.GetTextLen())
			GUIForm.logs.SetSelLength(int32(strings.Count(message, "")))
			GUIForm.logs.SelAttributes().SetColor(types.TColor(uint32(color)))
		}
		GUIForm.logs.Lines().Add(message)
		GUIForm.logs.Perform(messages.EM_SCROLLCARET, 7, 0)
	})
	logsLength++
	if logsLength > 10000 {
		clearLogs()
	}
}

func clearLogs() {
	logsLength = 0
	GUIForm.logs.Clear()
	GUIForm.logs.Lines().Clear()
}

func Logs(message ...string) {
	msg := ""
	for _, v := range message {
		msg += v
	}
	LogsColor(-1, msg)
}

func LogsStaticTime(message ...string) {
	if entity.ShowStaticLog {
		LogsTime(message...)
	}
}

func LogsProxyTime(message ...string) {
	if entity.ShowProxyLog {
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
