package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/http-server/src/entity"
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
	logs                    *lcl.TRichEdit
	proxyLogsGrid           *lcl.TStringGrid                        //代理详情列表UI
	ProxyLogsGridRowStyle   map[int32]*entity.ProxyLogsGridRowStyle //每行的样式
	proxyLogsGridCountRow   int32                                   //表格总行数
	proxyMouseMoveIndex     int32
	ProxyDetails            map[int32]*entity.ProxyDetail //代理详情数据集合
	ProxyDetailUI           *ProxyDetailPanel             //代理PanelUI
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
	m.ProxyDetails = make(map[int32]*entity.ProxyDetail)
	m.proxyLogsGridCountRow = 1
	m.ProxyLogsGridRowStyle = make(map[int32]*entity.ProxyLogsGridRowStyle)
	m.impl()
	//数据监听 channel
	go m.dataListen()
}

//数据监听 channel
func (m *TGUIForm) dataListen() {
	for {
		select {
		case proxyDetail, ok := <-entity.ProxyDetailChan:
			if ok {
				m.setProxyDetail(proxyDetail)
			}
		}
	}
}
