package main

import (
	"NoticeServices/app/model"
	"NoticeServices/plugins/sms/provider"
	"NoticeServices/plugins/sms/provider/alisms"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
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

func Send(sendParam map[string]interface{}, msg *model.InfoData) {

	pluginPath := g.Config().GetString("system.PluginPath")
	cfgFile := pluginPath + "/sms/config.toml"
	cfg := gcfg.New(cfgFile)
	defaultSms := cfg.GetString("DefaultSend")
	smsConfig := cfg.GetMap(defaultSms)
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
func SmsData(ctx *provider.Context, msg *model.InfoData) {
	var instance provider.SmsProviderInterface

	switch ctx.ProviderName {
	case "alisms":
		instance = &alisms.Instance{}
	case "qcloud":
		glog.Info("qcloud 短信提供商进行发送")

	default:

		glog.Info("未选择短信发送供应商")

	}

	err := instance.SendSms(ctx, msg)
	if err != nil {
		glog.Error(err)
	}

}

//测试插件是否可用
//func main() {
//
//	var msg = new(model.InfoData)
//	msg.AppId = "dfasdfasdf"
//	msg.ConfigId = "3eb5e3d5cd2c71ef6fce3f391c9eabcd"
//	msg.ComeFrom = "xxxadf"
//	msg.Method = "instant"
//	msg.MethodNum = 1
//	msg.MethodTask = "*"
//	msg.MsgTitle = "3556777744系統故障了"
//	msg.MsgBody = "我们的内容信息"
//	msg.MsgUrl = "http://www.baidu.com"
//	msg.UserIds = "aaaa|bbbb|cccc"
//	msg.PartyIds = ""
//	msg.Totag = "[{\"name\":\"mail\",\"value\":\"xinjy@neusoft.com\"},{\"name\":\"webhook\",\"value\":\"cccc\"},{\"name\":\"sms\",\"value\":\"13700005102\"}]"
//	unit.Logger.Info("开始进行测试。。。。")
//	Send(msg)
//
//}
