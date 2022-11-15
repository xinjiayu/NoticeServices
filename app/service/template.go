package service

import (
	"NoticeServices/app/dao"
	"NoticeServices/app/define"
	"NoticeServices/app/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

type templateService struct {
}

var Template = new(templateService)

//CreateTemplate 创建模板
func (c *templateService) CreateTemplate(data *define.TemplateData) (*model.Template, error) {
	var tpl *model.Template
	if err := gconv.Struct(data, &tpl); err != nil {
		g.Log().Error(context.TODO(), err)
		return nil, err
	}
	tpl.Id = guid.S()
	tpl.CreateTime = gconv.Int(gtime.Timestamp())
	if _, err := dao.Template.Ctx(context.TODO()).Insert(tpl); err != nil {
		return nil, err
	}
	return tpl, nil
}

//UpdateTemplate 修改
func (c *templateService) UpdateTemplate(data *define.TemplateUpData) error {
	_, err := dao.Template.Ctx(context.TODO()).Data(data).
		FieldsEx(dao.Template.Columns.Id).
		Where(dao.Template.Columns.Id, data.Id).
		Update()
	return err
}

//DeleteTemplate 删除
func (c *templateService) DeleteTemplate(id string) error {
	_, err := dao.Template.Ctx(context.TODO()).Where(dao.Template.Columns.Id, id).Delete()
	return err
}

//GetOneTemplate 获取一条记录
func (c *templateService) GetOneTemplate(id string) (data *model.Template, err error) {
	err = dao.Template.Ctx(context.TODO()).Where(dao.Template.Columns.Id, id).Scan(&data)
	return
}

//GetTemplateList 获取多条记录
func (c *templateService) GetTemplateList(ConfigId string) (datas []*model.Template, err error) {
	err = dao.Template.Ctx(context.TODO()).Where(dao.Template.Columns.ConfigId, ConfigId).Scan(&datas)
	return
}
