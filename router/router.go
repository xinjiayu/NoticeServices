package router

import (
	"NoticeServices/app/api/app"
	"NoticeServices/app/api/msg"
	"NoticeServices/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS, middleware.White, middleware.Auth)

		group.POST("/app", app.Create)
		group.GET("/app", app.Get)
		group.POST("/msgsend", msg.Send)

	})
}
