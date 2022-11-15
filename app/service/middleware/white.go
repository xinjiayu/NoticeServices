package middleware

import (
	"NoticeServices/library/tools"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

func White(r *ghttp.Request) {
	accessOk := true
	//获取客户端IP
	cip := r.GetRemoteIp()
	ipArr := g.Cfg().MustGet(context.TODO(), "system.whitelist").Array()
	if len(ipArr) > 0 {
		accessOk = tools.IsContains(cip, ipArr)
	}

	if accessOk {
		r.Middleware.Next()
	} else {
		g.Log().Debug(context.TODO(), cip, "此IP无权访问服务")
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
