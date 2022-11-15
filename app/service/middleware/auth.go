package middleware

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

func Auth(r *ghttp.Request) {

	secretKey := g.Cfg().MustGet(context.TODO(), "system.SecretKey").String()
	if secretKey != "" {
		if secretKey == r.Header.Get("secret-key") {
			r.Middleware.Next()
		} else {
			g.Log().Debug(context.TODO(), "SecretKey 不一致")
			r.Response.WriteStatus(http.StatusForbidden)
		}
	} else {
		r.Middleware.Next()

	}

}
