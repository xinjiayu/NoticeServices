package app

import (
	"NoticeServices/app/model"
	"NoticeServices/app/service"
	"NoticeServices/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// Create 创建一个通知应用
func Create(r *ghttp.Request) {

	var appData *model.AppData
	err := r.Parse(&appData)
	if err != nil {
		response.JsonExit(r, 1, "数据有错误", err.Error())

	}

	data, err2 := service.App.CreateApp(appData)
	if err2 != nil {
		response.JsonExit(r, 1, "数据有错误", err2.Error())

	}

	response.JsonExit(r, 0, "应用申请成功", data)

}

func Get(r *ghttp.Request) {

	appId := r.GetString("app_id")
	data := service.App.GetAppInfo(appId)

	response.JsonExit(r, 0, "数据源返回数据值", data)

}
