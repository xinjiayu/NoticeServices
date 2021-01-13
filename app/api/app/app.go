package app

import (
	"github.com/gogf/gf/net/ghttp"
)

// Create 创建一个通知应用
func Create(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
