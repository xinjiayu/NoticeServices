package service

import (
	"NoticeServices/app/dao"
	"NoticeServices/app/define"
	"NoticeServices/app/model"
	"NoticeServices/app/notifieer"
	"NoticeServices/boot"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

type msgService struct{}

var Msg = new(msgService)

//Send 发送信息
func (m *msgService) Send(message *define.InfoData) error {

	//将接收到通知信息存入数据库

	infoId, err := m.save(message)
	if err != nil {
		return err
	}

	jd := new(model.Job)
	jd.Name = message.MsgTitle
	jd.Group = message.Method
	jd.Params = gconv.String(infoId)       //将通知信息存入库后，将信息ID记录到任务参数中
	jd.InvokeTarget = "JobSendMessage"     //执行的方法 调用目标字符串
	jd.CronExpression = message.MethodCron // cron执行表达式
	jd.Status = 0                          // 状态（0正常 1暂停）
	jd.CreateTime = gconv.Int(gtime.Timestamp())

	switch message.Method {
	case boot.Instant:
		//调用发送通道进行发送
		//m.gateWaySend(message)
		notifieer.Instance.GateWaySend(message)

	case boot.Appointment:
		g.Log().Debug(context.TODO(), "预约发送处理==========")
		jd.MisfirePolicy = 2 // 计划执行策略（执行一次）
		jobId, err := JobAdd(jd)
		if err != nil {
			g.Log().Error(context.TODO(), err)
		}
		jd.Id = gconv.Int(jobId)
		if err = JobStart(jd); err != nil {
			g.Log().Error(context.TODO(), err)
		}

	case boot.Regular:
		g.Log().Debug(context.TODO(), "定期发送处理==========")
		jd.MisfirePolicy = 1 // 计划执行策略（多次执行）
		jobId, err := JobAdd(jd)
		if err != nil {
			g.Log().Error(context.TODO(), err)
		}
		jd.Id = gconv.Int(jobId)
		if err = JobStart(jd); err != nil {
			g.Log().Error(context.TODO(), err)
		}

	default:
		//调用发送通道进行发送
		//m.gateWaySend(message)
		notifieer.Instance.GateWaySend(message)

	}

	return nil

}

//MarkRead 标记用户的通知信息为已讯状态
func (m *msgService) MarkRead(id string) error {
	_, err := dao.UserInfo.Ctx(context.TODO()).
		Data(g.Map{dao.UserInfo.Columns.Status: 1}).
		WherePri(id).Update()
	if err != nil {
		return err
	}
	return nil
}

// GetInfoByUserID 通过app_id 与user_id 获取用户的通知信息列表
func (m *msgService) GetInfoByUserID(appId, userId string) ([]*define.EntityInfo, error) {

	uWhere := g.Map{
		dao.UserInfo.Columns.AppId:  appId,
		dao.UserInfo.Columns.UserId: userId,
	}
	var entityInfos []*define.EntityInfo

	err := dao.UserInfo.Ctx(context.TODO()).Fields("*").Where(uWhere).
		ScanList(&entityInfos, "UserInfo")
	if err != nil {
		return nil, err
	}

	err = dao.Info.Ctx(context.TODO()).Fields("id,msg_title,msg_body,msg_url").Where(dao.Info.Columns.Id, gutil.ListItemValues(&entityInfos, "UserInfo", "InfoId")).
		ScanList(&entityInfos, "Info")
	if err != nil {
		return nil, err
	}

	return entityInfos, nil
}

//save 将接收到的发送信息存入到数据库
func (m *msgService) save(message *define.InfoData) (int, error) {
	var info *model.Info
	if err := gconv.Struct(message, &info); err != nil {
		g.Log().Error(context.TODO(), err)
		return 0, err
	}

	info.State = "1"
	info.CreateTime = gconv.Int(gtime.Timestamp())
	resData, err := dao.Info.Ctx(context.TODO()).FieldsEx(dao.Info.Columns.Id).Data(info).Insert()
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return 0, err
	}
	infoId, err2 := resData.LastInsertId()
	if err2 != nil {
		g.Log().Error(context.TODO(), err2)
		return 0, err
	}

	if message.UserIds == "" {
		return 0, err
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
			if _, err := dao.UserInfo.Ctx(context.TODO()).FieldsEx(dao.UserInfo.Columns.Id).Insert(userInfo); err != nil {
				g.Log().Error(context.TODO(), err)
				return 0, err
			}
		}
	}
	return gconv.Int(infoId), nil
}
