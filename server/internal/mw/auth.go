package mw

import (
	"net/http"
	"server/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func AuthMiddleware(r *ghttp.Request) {
	if service.User().IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
