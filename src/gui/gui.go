package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/http-server/src/entity"
)

var GUIForm = &TGUIForm{}

type TGUIForm struct {
	*lcl.TForm
	width                   int32
	height                  int32
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
	m.init()
	m.SetCaption("Http Web Server")
	m.SetPosition(types.PoScreenCenter)
	//m.EnabledMaximize(false)
	m.SetBorderStyle(types.BsSingle)
	m.SetWidth(m.width)
	m.SetHeight(m.height)
	m.ProxyDetails = make(map[int32]*entity.ProxyDetail)
	m.impl()
	//数据监听
	go m.dataListen()
}

func (m *TGUIForm) init() {
	m.width = 600
	m.height = 350
	icon := lcl.NewIcon()
	icon.LoadFromFSFile("resources/app.ico")
	m.SetIcon(icon)
}

func (m *TGUIForm) dataListen() {
	for {
		select {
		case proxyDetail := <-entity.ProxyDetailChan:
			//fmt.Printf("%+v\n", proxyDetail)
			m.setProxyDetail(proxyDetail)
		}
	}
}
