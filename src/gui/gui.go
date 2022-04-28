package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/http-server/src/entity"
)

var GUIForm = &TGUIForm{}

var (
	uiWidth    int32 = 600
	uiHeight   int32 = 350
	uiWidthEx  int32 = 400
	uiHeightEx int32 = 400
)

type TGUIForm struct {
	*lcl.TForm
	logs                    *lcl.TRichEdit
	proxyLogsGrid           *lcl.TStringGrid              //代理详情列表UI
	ProxyDetails            map[int32]*entity.ProxyDetail //代理详情数据集合
	ProxyDetailUI           *ProxyDetailPanel             //代理PanelUI
	stbar                   *lcl.TStatusBar
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
	m.impl()
	//数据监听 channel
	go m.dataListen()
}

//数据监听 channel
func (m *TGUIForm) dataListen() {
	for {
		select {
		case proxyDetail := <-entity.ProxyDetailChan:
			//fmt.Printf("%+v\n", proxyDetail)
			m.setProxyDetail(proxyDetail)
		}
	}
}
