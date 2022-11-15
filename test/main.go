package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("Post:/test/webhook", func(r *ghttp.Request) {

		g.Log().Debug(context.TODO(), r.GetBody())

	})
	s.SetPort(8180)
	s.Run()

}
