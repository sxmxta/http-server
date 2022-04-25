package gui

import (
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
)

var GUIForm = &TGUIForm{}

type TGUIForm struct {
	*lcl.TForm
	width                   int32
	height                  int32
	logs                    *lcl.TRichEdit
	proxyLogsGrid           *lcl.TStringGrid       //代理详情列表
	ProxyDetails            map[int32]*ProxyDetail //代理详情数据
	stbar                   *lcl.TStatusBar
	showProxyLogChkBox      *lcl.TCheckBox
	ShowProxyLog            bool
	showStaticLogChkBox     *lcl.TCheckBox
	ShowStaticLog           bool
	enableProxyDetailChkBox *lcl.TCheckBox
	EnableProxyDetail       bool
}

type ProxyDetail struct {
	ID        int32
	Method    string
	SourceUrl string
	TargetUrl string
	Host      string
	Request   ProxyRequestDetail
	Response  ProxyResponseDetail
}

type ProxyRequestDetail struct {
	Header     map[string][]string
	Body       string
	URLParams  map[string][]string
	FormParams map[string][]string
}

type ProxyResponseDetail struct {
	Header map[string][]string
	Body   string
	Size   int64
}

func (m *TGUIForm) OnFormCreate(sender lcl.IObject) {
	m.init()
	m.SetCaption("Http Web Server")
	m.SetPosition(types.PoScreenCenter)
	//m.EnabledMaximize(false)
	m.SetWidth(m.width)
	m.SetHeight(m.height)
	//m.SetBorderStyle(types.BsSingle)
	m.ProxyDetails = make(map[int32]*ProxyDetail)
	m.impl()
}

func (m *TGUIForm) init() {
	m.width = 900
	m.height = 350
	icon := lcl.NewIcon()
	icon.LoadFromFSFile("resources/app.ico")
	m.SetIcon(icon)
}
