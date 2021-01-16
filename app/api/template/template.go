package template

import (
	"NoticeServices/app/model"
	"NoticeServices/app/service"
	"NoticeServices/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type NoticeTemplate struct {
}

func (n *NoticeTemplate) Get(r *ghttp.Request) {

	//获取单条记录
	id := r.GetString("id")
	if id != "" {
		data, err := service.Template.GetOneTemplate(id)
		if err != nil {
			response.JsonExit(r, 1, "数据获取错误", err.Error())
		}
		response.JsonExit(r, 0, "数据获取成功", data)
	}

	//获取多条记录
	configId := r.GetString("config_id")
	if configId != "" {
		data, err := service.Template.GetTemplateList(configId)
		if err != nil {
			response.JsonExit(r, 1, "数据获取错误", err.Error())
		}
		response.JsonExit(r, 0, "数据获取成功", data)
	}

	response.JsonExit(r, 0, "请传入模板ID（id）或是配置文件ID（config_id）")

}

func (n *NoticeTemplate) Post(r *ghttp.Request) {
	var data *model.TemplateData
	err := r.Parse(&data)
	if err != nil {
		response.JsonExit(r, 1, "数据有错误", err.Error())

	}

	resData, err2 := service.Template.CreateTemplate(data)
	if err2 != nil {
		response.JsonExit(r, 1, "数据有错误", err2.Error())

	}
	response.JsonExit(r, 0, "配置创建成功", resData)
}

func (n *NoticeTemplate) Put(r *ghttp.Request) {
	var data *model.TemplateUpData
	err := r.Parse(&data)
	if err != nil {
		response.JsonExit(r, 1, "数据有错误", err.Error())
	}

	err = service.Template.UpdateTemplate(data)
	if err != nil {
		response.JsonExit(r, 1, "修改失败", err.Error())

	}
	response.JsonExit(r, 0, "修改成功")
}
func (n *NoticeTemplate) Delete(r *ghttp.Request) {
	id := r.GetString("id")
	err := service.Template.DeleteTemplate(id)
	if err != nil {
		response.JsonExit(r, 1, "删除失败", err.Error())

	}
	response.JsonExit(r, 0, "删除成功")
}
