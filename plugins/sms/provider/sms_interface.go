package provider

import (
	"NoticeServices/app/define"
)

// Context 上下文
type Context struct {
	//供应商选择
	ProviderName  string                 `json:"provider_name"`
	ProviderTitle string                 `json:"provider_title"`
	SmsConfig     map[string]interface{} `json:"sms_config"`
	SendParam     map[string]interface{} `json:"send_param"`
}

type SmsProviderInterface interface {
	SendSms(ctx *Context, msg *define.InfoData) error
}
