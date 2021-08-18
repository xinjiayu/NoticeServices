package main

import (
	"NoticeServices/app/define"
	"NoticeServices/plugins/sms/provider"
	"NoticeServices/plugins/sms/provider/alisms"
	"NoticeServices/plugins/sms/provider/tencentcloud"
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

//goland:noinspection ALL
func Send(sendParam map[string]interface{}, msg *define.InfoData) {

	pluginPath := g.Config().GetString("system.PluginPath")
	cfgFile := pluginPath + "/sms/config.toml"
	//cfgFile := "config.toml" //本地程序直接测试的时候，把上面两句注释掉，打开这一句。执行本程序中的main方法。
	cfg := gcfg.New(cfgFile)
	defaultSms := cfg.GetString("DefaultSend")
	if defaultSms == "" {
		glog.Error("获取默认短信服务供应商配置出错")
		return
	}
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
func SmsData(ctx *provider.Context, msg *define.InfoData) {
	var instance provider.SmsProviderInterface
	switch ctx.ProviderName {
	case "alisms":
		instance = &alisms.Instance{}
	case "tencentcloud":
		instance = &tencentcloud.Instance{}

	default:

		glog.Info("未选择短信发送供应商")

	}

	glog.Info("发达短信供应商：", ctx.ProviderTitle)
	err := instance.SendSms(ctx, msg)
	if err != nil {
		glog.Error(err)
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
