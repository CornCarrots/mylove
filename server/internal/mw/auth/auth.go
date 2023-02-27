package auth

import (
	"net/http"
	"server/internal/mw"
	"server/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	authMiddleware struct{}
)

func init() {
	mw.RegisterMiddleware(NewAuthMiddleware())
}

func NewAuthMiddleware() *authMiddleware {
	return &authMiddleware{}
}

// Do validates the request to allow only signed-in users visit.
func (s *authMiddleware) Do(r *ghttp.Request) {
	if service.User().IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
