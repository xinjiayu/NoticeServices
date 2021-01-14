package service

import (
	"NoticeServices/app/dao"
	"NoticeServices/app/model"
	"NoticeServices/boot"
	"NoticeServices/library/tools"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"plugin"
)

type msgService struct{}

var Msg = new(msgService)

//Send 发送信息
func (m *msgService) Send(message *model.InfoData) error {

	//将接收到通知信息存入数据库
	go m.save(message)

	switch message.Method {
	case boot.Instant:
		//调用发送通道进行发送
		m.gateWaySend(message)

	case boot.Appointment:
		//TODO:添加预约发送处理
		glog.Info("预约发送处理==========")

	case boot.Regular:
		//TODO:添加定期发送处理
		glog.Info("定期发送处理==========")

	default:
		//调用发送通道进行发送
		m.gateWaySend(message)
	}

	return nil

}

//save 将接收到的发送信息存入到数据库
func (m *msgService) save(message *model.InfoData) {
	var info *model.Info
	if err := gconv.Struct(message, &info); err != nil {
		glog.Error(err)
		return
	}

	info.State = "1"
	info.CreateTime = gconv.Int(gtime.Timestamp())
	if _, err := dao.Info.FieldsEx("id").Data(info).Insert(); err != nil {
		glog.Error(err)
		return
	}
}

//getInfoConfig 读取通知信息的配置文件
func (m *msgService) getInfoConfig(configId string) (*model.Config, error) {
	infoConfig, err := dao.Config.FindOne("id", configId)
	if err != nil {
		glog.Error(err.Error())
		return nil, err
	}
	return infoConfig, nil
}

//gateWaySend 通过发送通道进行发送
func (m *msgService) gateWaySend(message *model.InfoData) {

	//获取指定通知的配置信息
	config, _ := m.getInfoConfig(message.ConfigId)
	sendGatewayList := gstr.Explode("|", config.SendGateway)
	if sendGatewayList == nil {
		return
	}

	messageBaseBody := message.MsgBody
	for _, gatewayName := range sendGatewayList {
		message.MsgBody = messageBaseBody

		//获取发送通道的通知模板
		where := g.Map{
			"config_id":    message.ConfigId,
			"send_gateway": gatewayName,
		}
		template, err := dao.Template.FindOne(where)
		if template != nil {
			paramDataMap := gconv.Map(message)
			message.MsgBody = tools.StringLiteralTemplate(template.Content, paramDataMap)

		}

		// 加载插件
		filename := "plugins/" + gatewayName + ".so"
		p, err := plugin.Open(filename)
		if err != nil {
			panic(err)
		}

		// 查找插件里的指定函数
		symbol, err := p.Lookup("Send")
		if err != nil {
			panic(err)
		}
		sendFunc, ok := symbol.(func(*model.InfoData))

		if !ok {
			glog.Error(gerror.New("Plugin has no Send function"))
			return
		}
		// 调用插件函数
		sendFunc(message)
	}
}
