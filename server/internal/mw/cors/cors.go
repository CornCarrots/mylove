package cors

import (
	"server/internal/mw"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	corsMiddleware struct{}
)

func init() {
	mw.RegisterMiddleware(NewCorsMiddleware())
}

func NewCorsMiddleware() *corsMiddleware {
	return &corsMiddleware{}
}

// Do validates the request to allow only signed-in users visit.
func (s *corsMiddleware) Do(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
