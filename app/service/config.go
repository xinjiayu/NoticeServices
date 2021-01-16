package service

import (
	"NoticeServices/app/dao"
	"NoticeServices/app/model"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/guid"
)

type configService struct{}

var Config = new(configService)

func (c *configService) CreateConfig(data *model.ConfigData) (*model.Config, error) {
	var cfg *model.Config
	if err := gconv.Struct(data, &cfg); err != nil {
		glog.Error(err)
		return nil, err
	}

	cfg.Id = guid.S()
	cfg.CreateTime = gconv.Int(gtime.Timestamp())
	if _, err := dao.Config.Insert(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

//SaveConfig 修改
func (c *configService) UpdateConfig(data *model.ConfigUpData) error {
	_, err := dao.Config.Data(data).
		FieldsEx(dao.Config.Columns.Id).
		Where(dao.Config.Columns.Id, data.Id).
		Update()
	return err
}

//Delete 删除
func (c *configService) DeleteConfig(id string) error {
	_, err := dao.Config.Where(dao.Config.Columns.Id, id).Delete()
	return err
}

//GetOneConfig 获取一条配置记录
func (c *configService) GetOneConfig(id string) (*model.Config, error) {
	return dao.Config.Where(dao.Config.Columns.Id, id).One()
}

//GetConfigList 获取多条配置记录
func (c *configService) GetConfigList(r *model.ConfigServiceGetListReq) (*model.ConfigServiceGetListRes, error) {
	m := dao.Config.Fields("*")

	if r.Type != "" {
		m = m.Where(dao.Config.Columns.Type, r.Type)
	}
	if r.Key != "" {
		likePattern := `%` + r.Key + `%`
		m = m.Where(dao.Config.Columns.Name+" LIKE ?", likePattern).Or(dao.Config.Columns.SendGateway+" LIKE ?", likePattern)
	}

	listModel := m.Page(r.Page, r.Size)

	configEntities, err := listModel.M.All()
	if err != nil {
		return nil, err
	}
	// 没有数据
	if configEntities.IsEmpty() {
		return nil, nil
	}
	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &model.ConfigServiceGetListRes{
		Page:  r.Page,
		Size:  r.Size,
		Total: total,
	}
	// Config
	getListRes.List = configEntities

	return getListRes, nil

}
