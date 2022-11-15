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

type appService struct {
}

var App = new(appService)

func (a *appService) GetAppInfo(appId string) *model.App {
	app := new(model.App)
	if err := dao.App.Ctx(context.TODO()).Where("id", appId).Scan(&app); err != nil {
		g.Log().Error(context.TODO(), err.Error())
	}
	return app
}

func (a *appService) CreateApp(appData *define.AppData) (*model.App, error) {

	app := new(model.App)

	app.Id = guid.S()
	app.Name = appData.Name
	app.Explain = appData.Explain
	app.AccessToken = guid.S()
	app.CreateTime = gconv.Int(gtime.Timestamp())
	if _, err := dao.App.Ctx(context.TODO()).Insert(app); err != nil {
		return nil, err
	}

	return app, nil
}
