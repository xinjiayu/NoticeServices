package internal

import (
	"NoticeServices/plugins/wework/model"
	"encoding/json"
	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/message"
	"github.com/gogf/gf/encoding/gparser"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

type Alarm struct {
	Corp      *corporation.Corporation
	CorpApp   *corporation.App
	AppConfig corporation.AppConfig
}

/**
 * 建立私有变量
 */
var instance *Alarm

func GetInstance(corpid, agentID, secret, token, encodingAESKey string) *Alarm {
	instance = new(Alarm)
	// 加载应用的配置
	appConfigInfo := g.Cfg().Get("wework.alarm", corporation.AppConfig{})
	p := gparser.New(appConfigInfo)
	p.Struct(&instance.AppConfig)

	instance.AppConfig.Token = token
	instance.AppConfig.AgentId = agentID
	instance.AppConfig.Secret = secret
	instance.AppConfig.EncodingAESKey = encodingAESKey

	instance.Corp = corporation.New(corporation.Config{Corpid: corpid})
	instance.CorpApp = instance.Corp.NewApp(instance.AppConfig)

	return instance
}

func (e *Alarm) SendMessage(toUser, content string) (interface{}, error) {
	if content == "" {
		return nil, gerror.New("发送的内容为空值")
	}

	sendMsg := new(model.Text)
	sendMsg.Agentid = e.AppConfig.AgentId
	sendMsg.Touser = toUser //"@all"
	sendMsg.Msgtype = "text"
	sendMsg.Text.Content = content
	sendMsg.DuplicateCheckInterval = 1800

	json := ToJson(sendMsg)

	resp, err := message.Send(e.CorpApp, json)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ToJson(data interface{}) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		glog.Error(err)
	}
	return json
}
