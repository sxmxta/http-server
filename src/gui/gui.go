package gui

import (
	"gitee.com/snxamdf/http-server/src/entity"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

var GUIForm = &TGUIForm{}

var (
	logsLength   int
	uiWidth      int32 = 600
	uiHeight     int32 = 350
	uiWidthEx    int32 = 650
	uiHeightEx   int32 = 400
	stateBarText       = "https://gitee.com/snxamdf/http-server"
)

type TGUIForm struct {
	*lcl.TForm
	leftPanel               *lcl.TPanel
	rightPanel              *RightPanelUI //代理PanelUI
	lrSplitter              *lcl.TSplitter
	logs                    *lcl.TRichEdit
	proxyLogsGrid           *lcl.TStringGrid                        //代理详情列表UI
	proxyLogsGridColAddr    *lcl.TGridColumn                        //
	ProxyLogsGridRowStyle   map[int32]*entity.ProxyLogsGridRowStyle //每行的样式
	proxyLogsGridCountRow   int32                                   //表格总行数
	proxyMouseMoveIndex     int32
	ProxyDetails            map[int32]*entity.ProxyDetail //代理详情数据集合
	stateBar                *lcl.TStatusBar
	showProxyLogChkBox      *lcl.TCheckBox
	showStaticLogChkBox     *lcl.TCheckBox
	enableProxyDetailChkBox *lcl.TCheckBox
}

func (m *TGUIForm) OnFormCreate(sender lcl.IObject) {
	m.Icon().LoadFromFSFile("resources/app.ico")
	m.SetCaption("Http Web Server")
	m.SetPosition(types.PoScreenCenter)
	//m.EnabledMaximize(false)
	m.SetBorderStyle(types.BsSingle)
	m.SetWidth(uiWidth)
	m.SetHeight(uiHeight)

	m.leftPanel = lcl.NewPanel(m)
	m.leftPanel.SetParent(m)
	m.leftPanel.SetAlign(types.AlLeft)
	m.leftPanel.SetWidth(uiWidth)

	m.lrSplitter = lcl.NewSplitter(m)
	m.lrSplitter.SetParent(m)
	m.lrSplitter.SetAlign(types.AlLeft)
	m.lrSplitter.SetWidth(3)
	m.lrSplitter.SetLeft(uiWidth)
	m.lrSplitter.SetName("LRSplitter")

	m.ProxyDetails = make(map[int32]*entity.ProxyDetail)
	m.proxyLogsGridCountRow = 1
	m.ProxyLogsGridRowStyle = make(map[int32]*entity.ProxyLogsGridRowStyle)

	//初始化所有实现UI
	m.impl()
	//数据监听 channel
	go m.dataListen()
}

//数据监听 channel
func (m *TGUIForm) dataListen() {
	for {
		select {
		case proxyDetail, ok := <-entity.ProxyDetailGridChan:
			if ok {
				m.setProxyDetail(proxyDetail)
			}
		case proxyInterceptDetail, ok := <-entity.ProxyFlowInterceptChan:
			if ok {
				//添加到拦截队列任务
				m.rightPanel.ProxyInterceptConfigPanel.interceptQueue(proxyInterceptDetail)
			}
		}
	}
}
