package main

import (
	"NoticeServices/app/define"
	"NoticeServices/plugins/wework/internal"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
)

var logger *glog.Logger

func init() {
	logger = glog.New()
	logger.SetConfigWithMap(g.Map{
		"path":     "./log/wework",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
	})
}

func Send(sendParam map[string]interface{}, msg *define.InfoData) {
	g.Log().Debug(context.TODO(), "wework发送开始")

	pluginPath := g.Cfg().MustGet(context.TODO(), "system.PluginPath").String()

	cfgFile := pluginPath + "/wework/config.toml"
	cfg, err := gcfg.NewAdapterFile(cfgFile)
	if err != nil {
		g.Log().Error(context.TODO(), err)
	}

	//cfgFile := "config.toml" //用于本地main方法直接测试使用，需要将上面的配置注释掉

	var sendObjectList []define.SendObject
	err = gjson.DecodeTo(msg.Totag, &sendObjectList)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
	corpid := cfg.MustGet(context.TODO(), "weworkAlarm.Corpid").String()
	agentID := cfg.MustGet(context.TODO(), "weworkAlarm.AgentID").String()
	secret := cfg.MustGet(context.TODO(), "weworkAlarm.Secret").String()
	token := cfg.MustGet(context.TODO(), "weworkAlarm.Token").String()
	encodingAESKey := cfg.MustGet(context.TODO(), "weworkAlarm.EncodingAESKey").String()
	alarmService := internal.GetInstance(corpid, agentID, secret, token, encodingAESKey)
	for _, object := range sendObjectList {
		if object.Name == "wework" {
			toUser := object.Value
			content := msg.MsgBody
			g.Log().Debug(context.TODO(), toUser, content)
			data, err := alarmService.SendMessage(toUser, content)
			if err != nil {
				g.Log().Error(context.TODO(), err)
			}
			g.Log().Debug(context.TODO(), data)
		}
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
//	msg.MsgBody = "这是一个系統提示信息"
//	msg.MsgUrl = "http://www.baidu.com"
//	msg.UserIds = "aaaa|bbbb|cccc"
//	msg.PartyIds = ""
//	msg.Totag = "[{\"name\":\"mail\",\"value\":\"xinjy@qq.com\"},{\"name\":\"wework\",\"value\":\"JiaYu\"},{\"name\":\"webhook\",\"value\":\"cccc\"},{\"name\":\"sms\",\"value\":\"13400225102\"}]"
//	Send(sendParam, msg)
//
//}
