package main

import (
	"NoticeServices/app/define"
	"NoticeServices/plugins/sms/provider"
	"NoticeServices/plugins/sms/provider/alisms"
	"NoticeServices/plugins/sms/provider/tencentcloud"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/gconv"
)

type Options struct {
	PayloadURL string
	Secret     string
	Subject    string
	Body       string
}

type WebhookConfig struct {
	PayloadURL string
	Secret     string
}

//goland:noinspection ALL
func Send(sendParam map[string]interface{}, msg *define.InfoData) {

	pluginPath := g.Cfg().MustGet(context.TODO(), "system.PluginPath").String()

	cfgFile := pluginPath + "/webhook/config.toml"
	cfg, err := gcfg.NewAdapterFile(cfgFile)
	if err != nil {
		g.Log().Error(context.TODO(), err)
	}
	defaultSms := cfg.MustGet(context.TODO(), "DefaultSend").String()
	//cfgFile := "config.toml" //本地程序直接测试的时候，把上面两句注释掉，打开这一句。执行本程序中的main方法。
	if defaultSms == "" {
		g.Log().Error(context.TODO(), "获取默认短信服务供应商配置出错")
		return
	}

	smsConfig := cfg.MustGet(context.TODO(), "defaultSms").Map()
	title := gconv.String(smsConfig["title"])
	//初始化上下文
	ctx := &provider.Context{
		ProviderName:  defaultSms,
		ProviderTitle: title,
		SmsConfig:     smsConfig,
		SendParam:     sendParam,
	}
	SmsData(ctx, msg)
}

//SmsData
func SmsData(ctx *provider.Context, msg *define.InfoData) {
	var instance provider.SmsProviderInterface
	switch ctx.ProviderName {
	case "alisms":
		instance = &alisms.Instance{}
	case "tencentcloud":
		instance = &tencentcloud.Instance{}

	default:

		g.Log().Debug(context.TODO(), "未选择短信发送供应商")

	}

	g.Log().Debug(context.TODO(), "发达短信供应商：", ctx.ProviderTitle)
	err := instance.SendSms(ctx, msg)
	if err != nil {
		g.Log().Error(context.TODO(), err)
	}

}

////测试插件是否可用
//func main() {
//
//	sendParam := make(map[string]interface{})
//	sendParam["code"] = "SMS_185570003"
//
//	var msg = new(model.InfoData)
//	msg.AppId = "dfasdfasdf"
//	msg.ConfigId = "3eb5e3d5cd2c71ef6fce3f391c9eabcd"
//	msg.ComeFrom = "xxxadf"
//	msg.Method = "222instant"
//	msg.MethodNum = 1
//	msg.MethodCron = "*"
//	msg.MsgTitle = "3556777744系統故障了"
//	msg.MsgBody = "{\"code\":5432}"
//	msg.MsgUrl = "http://www.baidu.com"
//	msg.UserIds = "aaaa|bbbb|cccc"
//	msg.PartyIds = ""
//	msg.Totag = "[{\"name\":\"mail\",\"value\":\"xinjy@neusoft.com\"},{\"name\":\"webhook\",\"value\":\"cccc\"},{\"name\":\"sms\",\"value\":\"13700005102\"}]"
//	Send(sendParam, msg)
//
//}
