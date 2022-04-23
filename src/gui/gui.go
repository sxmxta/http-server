package gui

import (
	"encoding/json"
	"fmt"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types"
)

var GUIForm = &TGUIForm{}

type TGUIForm struct {
	*lcl.TForm
	width                   int32
	height                  int32
	logs                    *lcl.TRichEdit
	logGrid                 *lcl.TStringGrid
	stbar                   *lcl.TStatusBar
	showProxyLogChkBox      *lcl.TCheckBox
	ShowProxyLog            bool
	showStaticLogChkBox     *lcl.TCheckBox
	ShowStaticLog           bool
	enableProxyDetailChkBox *lcl.TCheckBox
	EnableProxyDetail       bool
	ProxyDetail             map[int]*ProxyDetail
}

type ProxyDetail struct {
	ID       int
	Method   string
	URL      string
	Host     string
	Request  ProxyRequestDetail
	Response ProxyResponseDetail
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
	m.ProxyDetail = make(map[int]*ProxyDetail)
	m.impl()
}

func (m *TGUIForm) SetProxyDetail(proxyDetail *ProxyDetail) {
	m.ProxyDetail[proxyDetail.ID] = proxyDetail
	//add list grid
	d, _ := json.Marshal(proxyDetail)
	fmt.Println("\nproxyDetail:", proxyDetail.URL, " JSON:", string(d))
}

func (m *TGUIForm) init() {
	m.width = 600
	m.height = 400
	icon := lcl.NewIcon()
	icon.LoadFromFSFile("resources/app.ico")
	m.SetIcon(icon)
}
