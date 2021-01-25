package main

import (
	"NoticeServices/app/model"
	"NoticeServices/plugins/wework/internal"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/os/glog"
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

func Send(sendParam map[string]interface{}, msg *model.InfoData) {
	logger.Info("wework发送开始")
	pluginPath := g.Config().GetString("system.PluginPath")
	cfgFile := pluginPath + "/wework/config.toml"

	//cfgFile := "config.toml" //用于本地main方法直接测试使用，需要将上面的配置注释掉

	cfg := gcfg.New(cfgFile)
	var sendObjectList []model.SendObject
	err := gjson.DecodeTo(msg.Totag, &sendObjectList)
	if err != nil {
		logger.Error(err)
		return
	}
	corpid := cfg.GetString("weworkAlarm.Corpid")
	agentID := cfg.GetString("weworkAlarm.AgentID")
	secret := cfg.GetString("weworkAlarm.Secret")
	token := cfg.GetString("weworkAlarm.Token")
	encodingAESKey := cfg.GetString("weworkAlarm.EncodingAESKey")
	alarmService := internal.GetInstance(corpid, agentID, secret, token, encodingAESKey)
	for _, object := range sendObjectList {
		if object.Name == "wework" {
			toUser := object.Value
			content := msg.MsgBody
			logger.Info(toUser, content)
			data, err := alarmService.SendMessage(toUser, content)
			if err != nil {
				logger.Error(err)
			}
			logger.Info(data)
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
