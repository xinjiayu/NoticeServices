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
	"github.com/gogf/gf/util/gutil"
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

//标记用户的通知信息为已讯状态
func (m *msgService) MarkRead(id string) error {
	_, err := dao.UserInfo.
		Data(g.Map{dao.UserInfo.Columns.Status: 1}).
		WherePri(id).Update()
	if err != nil {
		return err
	}
	return nil
}

// GetInfoByUserID 通过app_id 与user_id 获取用户的通知信息列表
func (m *msgService) GetInfoByUserID(appId, userId string) ([]*model.EntityInfo, error) {

	uWhere := g.Map{
		dao.UserInfo.Columns.AppId:  appId,
		dao.UserInfo.Columns.UserId: userId,
	}
	var entityInfos []*model.EntityInfo

	err := dao.UserInfo.Fields("*").Where(uWhere).
		ScanList(&entityInfos, "UserInfo")
	if err != nil {
		return nil, err
	}

	err = dao.Info.Fields("id,msg_title,msg_body,msg_url").Where(dao.Info.Columns.Id, gutil.ListItemValues(&entityInfos, "UserInfo", "InfoId")).
		ScanList(&entityInfos, "Info")
	if err != nil {
		return nil, err
	}

	return entityInfos, nil
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
	resData, err := dao.Info.FieldsEx(dao.Info.Columns.Id).Data(info).Insert()
	if err != nil {
		glog.Error(err)
		return
	}
	infoId, err2 := resData.LastInsertId()
	if err2 != nil {
		glog.Error(err2)
		return
	}

	if message.UserIds == "" {
		return
	}

	//存入用户关系表
	userInfo := new(model.UserInfo)
	userInfo.InfoId = gconv.Int(infoId)
	userInfo.AppId = message.AppId
	userInfo.Status = "0"
	userInfo.CreateTime = gconv.Int(gtime.Timestamp())
	//获取目标用户列表
	userList := gstr.Explode("|", message.UserIds)
	if userList != nil {
		for _, u := range userList {
			userInfo.UserId = u
			if _, err := dao.UserInfo.FieldsEx(dao.UserInfo.Columns.Id).Insert(userInfo); err != nil {
				glog.Error(err)
				return
			}
		}
	}

}

//getInfoConfig 读取通知信息的配置文件
func (m *msgService) getInfoConfig(configId string) (*model.Config, error) {
	infoConfig, err := dao.Config.FindOne(dao.Config.Columns.Id, configId)
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
			glog.Error(err)
			return
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
