package mw

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func LogMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
	errStr := ""
	if err := r.GetError(); err != nil {
		errStr = err.Error()
	}
	g.Log().Info(r.Context(), r.Response.Status, r.URL.Path, errStr)
}
