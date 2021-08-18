package config

import (
	"NoticeServices/app/define"
	"NoticeServices/app/service"
	"NoticeServices/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type NoticeConfig struct{}

func (nc *NoticeConfig) Get(r *ghttp.Request) {

	//获取单条记录
	id := r.GetString("id")
	if id != "" {
		data, err := service.Config.GetOneConfig(id)
		if err != nil {
			response.JsonExit(r, 1, "数据获取错误", err.Error())
		}
		response.JsonExit(r, 0, "数据获取成功", data)
	}

	//获取多条记录
	var reqData *define.ConfigServiceGetListReq
	err := r.Parse(&reqData)
	if err != nil {
		response.JsonExit(r, 1, "提交的参数数据有错误", err.Error())
	}

	resData, err2 := service.Config.GetConfigList(reqData)
	if err2 != nil {
		response.JsonExit(r, 1, "数据获取出错", err2.Error())

	}
	response.JsonExit(r, 0, "数据获取成功", resData)
}

func (nc *NoticeConfig) Post(r *ghttp.Request) {
	var data *define.ConfigData
	err := r.Parse(&data)
	if err != nil {
		response.JsonExit(r, 1, "数据有错误", err.Error())

	}

	resData, err2 := service.Config.CreateConfig(data)
	if err2 != nil {
		response.JsonExit(r, 1, "数据有错误", err2.Error())

	}
	response.JsonExit(r, 0, "配置创建成功", resData)

}

func (nc *NoticeConfig) Put(r *ghttp.Request) {
	var data *define.ConfigUpData
	err := r.Parse(&data)
	if err != nil {
		response.JsonExit(r, 1, "数据有错误", err.Error())
	}

	err = service.Config.UpdateConfig(data)
	if err != nil {
		response.JsonExit(r, 1, "修改失败", err.Error())

	}
	response.JsonExit(r, 0, "修改成功")
}

func (nc *NoticeConfig) Delete(r *ghttp.Request) {
	id := r.GetString("id")
	err := service.Config.DeleteConfig(id)
	if err != nil {
		response.JsonExit(r, 1, "删除失败", err.Error())

	}
	response.JsonExit(r, 0, "删除成功")
}
