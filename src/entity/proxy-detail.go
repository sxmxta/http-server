package entity

import (
	"gitee.com/snxamdf/golcl/lcl/types"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/consts"
)

//代理请求详情
type ProxyDetail struct {
	Error     error               `json:"-"`
	State     consts.ProxyFlow    `json:"-"`
	ID        int32               `json:"-"`
	Method    string              `json:"method"`
	SourceUrl string              `json:"source_url"`
	TargetUrl string              `json:"target_url"`
	Host      string              `json:"host"`
	Request   ProxyRequestDetail  `json:"request"`
	Response  ProxyResponseDetail `json:"response"`
	Row       int32               `json:"-"`
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
		return "发送请求", colors.ClLawngreen
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
