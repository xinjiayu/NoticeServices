package middleware

import "github.com/gogf/gf/v2/net/ghttp"

// 允许接口跨域请求
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
