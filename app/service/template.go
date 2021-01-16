package service

import (
	"NoticeServices/app/dao"
	"NoticeServices/app/model"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/guid"
)

type templateService struct {
}

var Template = new(templateService)

//CreateTemplate 创建模板
func (c *templateService) CreateTemplate(data *model.TemplateData) (*model.Template, error) {
	var tpl *model.Template
	if err := gconv.Struct(data, &tpl); err != nil {
		glog.Error(err)
		return nil, err
	}
	tpl.Id = guid.S()
	tpl.CreateTime = gconv.Int(gtime.Timestamp())
	if _, err := dao.Template.Insert(tpl); err != nil {
		return nil, err
	}
	return tpl, nil
}

//UpdateTemplate 修改
func (c *templateService) UpdateTemplate(data *model.TemplateUpData) error {
	_, err := dao.Template.Data(data).
		FieldsEx(dao.Template.Columns.Id).
		Where(dao.Template.Columns.Id, data.Id).
		Update()
	return err
}

//DeleteTemplate 删除
func (c *templateService) DeleteTemplate(id string) error {
	_, err := dao.Template.Where(dao.Template.Columns.Id, id).Delete()
	return err
}

//GetOneTemplate 获取一条记录
func (c *templateService) GetOneTemplate(id string) (*model.Template, error) {
	return dao.Template.Where(dao.Template.Columns.Id, id).One()
}

//GetTemplateList 获取多条记录
func (c *templateService) GetTemplateList(ConfigId string) ([]*model.Template, error) {
	return dao.Template.FindAll(dao.Template.Columns.ConfigId, ConfigId)
}
