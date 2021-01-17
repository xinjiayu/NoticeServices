package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func main() {
	s := g.Server()
	s.BindHandler("Post:/test/webhook", func(r *ghttp.Request) {

		glog.Info(r.GetBody())

	})
	s.SetPort(8180)
	s.Run()

}
