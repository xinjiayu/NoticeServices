package msg

import (
	"NoticeServices/app/model"
	"NoticeServices/app/service"
	"NoticeServices/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func Send(r *ghttp.Request) {

	var msgData *model.InfoData
	err := r.Parse(&msgData)
	if err != nil {
		response.JsonExit(r, 1, "数据有错误", err.Error())

	}
	glog.Info(msgData)
	err = service.Msg.Send(msgData)
	if err != nil {
		response.JsonExit(r, 2, "信息发送失败", err.Error())

	}

	response.JsonExit(r, 0, "信息发送成功")

}
