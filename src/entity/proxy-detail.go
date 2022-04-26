package entity

import "gitee.com/snxamdf/http-server/src/consts"

//代理请求详情
type ProxyDetail struct {
	Error     error               `json:"-"`
	State     consts.ProxyFlow    `json:"-"`
	ID        int32               `json:"id"`
	Method    string              `json:"method"`
	SourceUrl string              `json:"source_url"`
	TargetUrl string              `json:"target_url"`
	Host      string              `json:"host"`
	Request   ProxyRequestDetail  `json:"request"`
	Response  ProxyResponseDetail `json:"response"`
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
