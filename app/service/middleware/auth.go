package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"net/http"
)

func Auth(r *ghttp.Request) {

	secretKey := g.Config().GetString("system.SecretKey")
	if secretKey != "" {
		if secretKey == r.Header.Get("secret-key") {
			r.Middleware.Next()
		} else {
			glog.Info("SecretKey 不一致")
			r.Response.WriteStatus(http.StatusForbidden)
		}
	} else {
		r.Middleware.Next()

	}

}
