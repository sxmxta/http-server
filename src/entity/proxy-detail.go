package entity

import (
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/consts"
)

type ProxyLogsGridRowStyle struct {
	cols map[int32]*ProxyLogsGridColStyle
}

type ProxyLogsGridColStyle struct {
	isColor     bool
	isFontStyle bool
	text        string
	tColor      types.TColor
	fontStyle   types.TFontStyle
}

//代理请求详情
type ProxyDetail struct {
	Error                error               `json:"-"`
	State                consts.ProxyFlow    `json:"-"`
	ID                   int32               `json:"-"`
	ProxyInterceptSignal chan int32          `json:"-"` //代理拦截产生的信号，数字标记状态 10:开始请求拦截 11:结束请求拦截， 20:开始响应拦截 21:结束响应拦截
	IsAddTaskQueue       bool                `json:"-"` //是否添加到任务队列 ，只有在初始请求时第一次添加 true 添加
	StateCode            int                 `json:"state_code"`
	Method               string              `json:"method"`
	SourceUrl            string              `json:"source_url"`
	TargetUrl            string              `json:"target_url"`
	Host                 string              `json:"host"`
	Request              ProxyRequestDetail  `json:"request"`
	Response             ProxyResponseDetail `json:"response"`
}

type ProxyRequestDetail struct {
	Header     map[string][]string `json:"header"`
	Body       string              `json:"body"`
	URLParams  map[string][]string `json:"url_params"`
	FormParams map[string][]string `json:"form_params"`
	Size       int64               `json:"size"`
}

type ProxyResponseDetail struct {
	Header map[string][]string `json:"header"`
	Body   string              `json:"body"`
	Size   int64               `json:"size"`
}

func (m *ProxyDetail) GetState() (string, types.TColor) {
	if m.State == consts.P0 {
		return "初始化", colors.ClBlack
	} else if m.State == consts.P1 {
		return "初始失败", colors.ClRed
	} else if m.State == consts.P2 {
		return "发送请求", colors.ClMediumseagreen
	} else if m.State == consts.P3 {
		return "发送失败", colors.ClRed
	} else if m.State == consts.P4 {
		return "发送成功", colors.ClGreen
	} else if m.State == consts.P5 {
		return "响应失败", colors.ClRed
	} else {
		return " - - ", colors.ClGray
	}
}

func (m *ProxyLogsGridRowStyle) SetColStyle(col int32, style *ProxyLogsGridColStyle) {
	m.cols[col] = style
}
func (m *ProxyLogsGridRowStyle) GetColStyle(col int32) *ProxyLogsGridColStyle {
	if colStyle, ok := m.cols[col]; ok {
		return colStyle
	} else {
		colStyle = NewColStyle(0, 0)
		colStyle.SetTColor(colors.ClBlack)
		m.cols[col] = colStyle
		return colStyle
	}
}
func (m *ProxyLogsGridRowStyle) GetCols() map[int32]*ProxyLogsGridColStyle {
	return m.cols
}

func NewRowStyle() *ProxyLogsGridRowStyle {
	return &ProxyLogsGridRowStyle{cols: map[int32]*ProxyLogsGridColStyle{}}
}

func NewColStyle(tColor types.TColor, font types.TFontStyle) *ProxyLogsGridColStyle {
	return &ProxyLogsGridColStyle{tColor: tColor, fontStyle: font}
}

func (m *ProxyLogsGridColStyle) TColor() types.TColor {
	return m.tColor
}
func (m *ProxyLogsGridColStyle) FontStyle() types.TFontStyle {
	return m.fontStyle
}
func (m *ProxyLogsGridColStyle) SetTColor(color types.TColor) *ProxyLogsGridColStyle {
	m.isColor = true
	m.tColor = color
	return m
}
func (m *ProxyLogsGridColStyle) SetFontStyle(fontStyle types.TFontStyle) *ProxyLogsGridColStyle {
	m.isFontStyle = true
	m.fontStyle = fontStyle
	return m
}

func (m *ProxyLogsGridColStyle) IsFontStyle() bool {
	return m.isFontStyle
}
func (m *ProxyLogsGridColStyle) IsColor() bool {
	return m.isColor
}
func (m *ProxyLogsGridColStyle) Text() string {
	return m.text
}
func (m *ProxyLogsGridColStyle) SetText(text string) *ProxyLogsGridColStyle {
	m.text = text
	return m
}
