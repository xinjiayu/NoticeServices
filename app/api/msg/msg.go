package msg

import (
	"NoticeServices/app/model"
	"NoticeServices/app/service"
	"NoticeServices/library/response"
	"github.com/gogf/gf/net/ghttp"
)

func Send(r *ghttp.Request) {
	var msgData *model.InfoData
	err := r.Parse(&msgData)
	if err != nil {
		response.JsonExit(r, 1, "数据有错误", err.Error())
	}
	err = service.Msg.Send(msgData)
	if err != nil {
		response.JsonExit(r, 2, "信息发送失败", err.Error())
	}
	response.JsonExit(r, 0, "信息发送成功")

}

func GetMsg(r *ghttp.Request) {
	appId := r.GetString("app_id")
	userId := r.GetString("user_id")

	resData, err := service.Msg.GetInfoByUserID(appId, userId)
	if err != nil {
		response.JsonExit(r, 1, "信息获取失败", err.Error())
	}
	response.JsonExit(r, 0, "信息获取成功", resData)

}

func MarkRead(r *ghttp.Request) {
	id := r.GetString("id")
	err := service.Msg.MarkRead(id)
	if err != nil {
		response.JsonExit(r, 1, "标记通知信息为已读状态失败", err.Error())
	}
	response.JsonExit(r, 0, "标记为已读状态成功")
}
