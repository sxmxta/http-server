package gui

import (
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/http-server/src/consts"
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
		case proxyDetail := <-entity.ProxyDetailChan:
			m.setProxyDetail(proxyDetail)
		case proxyInterceptDetail := <-entity.ProxyFlowInterceptChan:
			//请求拦截-此时会阻塞，等待确定发送请求或响应
			if proxyInterceptDetail.State == consts.P2 {
				m.ProxyDetailUI.ProxyInterceptConfigPanel.updateRequestUI(proxyInterceptDetail)
			} else if proxyInterceptDetail.State == consts.P4 {
				m.ProxyDetailUI.ProxyInterceptConfigPanel.updateResponseUI(proxyInterceptDetail)
			}
		case signal := <-entity.ProxyInterceptSignal:
			//10:开始请求拦截 11:结束请求拦截， 20:开始响应拦截 21:结束响应拦截
			fmt.Println(signal)
			m.ProxyDetailUI.ProxyInterceptConfigPanel.State = signal
			lcl.ThreadSync(func() {
				if signal == consts.SIGNAL10 { //10:开始请求拦截 - 阻塞请求
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateOkBtn.SetVisible(true)
					m.ProxyDetailUI.ProxyInterceptConfigPanel.switchRequestPage()
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.Font().SetColor(0x8000FF)
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.SetCaption("请求拦截，待确认")
				} else if signal == consts.SIGNAL20 { //20:开始响应拦截 - 阻塞响应
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateOkBtn.SetVisible(true)
					m.ProxyDetailUI.ProxyInterceptConfigPanel.switchResponsePage()
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.Font().SetColor(0x8000FF)
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.SetCaption("响应拦截，待确认")
				} else if signal == consts.SIGNAL22 { //请求超时-请求响应失败
					m.ProxyDetailUI.ProxyInterceptConfigPanel.switchResponsePage()
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.Font().SetColor(0x8000FF)
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.SetCaption("请求响应失败-超时")
				} else if signal == consts.SIGNAL23 { //请求超时-响应成功
					m.ProxyDetailUI.ProxyInterceptConfigPanel.switchResponsePage()
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.Font().SetColor(0x8000FF)
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.SetCaption("请求响应成功")
				} else if signal == consts.SIGNAL24 { //请求超时-响应客户端失败
					m.ProxyDetailUI.ProxyInterceptConfigPanel.switchResponsePage()
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.Font().SetColor(0x8000FF)
					m.ProxyDetailUI.ProxyInterceptConfigPanel.StateLabel.SetCaption("响应客户端失败")
				}
			})
		}
	}
}
