package router

import (
	"NoticeServices/app/api/app"
	"NoticeServices/app/api/config"
	"NoticeServices/app/api/msg"
	"NoticeServices/app/api/template"
	"NoticeServices/app/service/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS, middleware.White, middleware.Auth)

		group.POST("/app", app.Create)
		group.GET("/app", app.Get)

		group.GET("/markread", msg.MarkRead)
		group.GET("/msg", msg.GetMsg)
		group.POST("/msgsend", msg.Send)
		group.REST("/config", new(config.NoticeConfig))
		group.REST("/template", new(template.NoticeTemplate))

	})
}
