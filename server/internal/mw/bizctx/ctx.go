package bizctx

import (
	"server/internal/model"
	"server/internal/mw"
	"server/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ctxMiddleware struct{}
)

func init() {
	mw.RegisterMiddleware(NewCtxMiddleware())
}

func NewCtxMiddleware() *ctxMiddleware {
	return &ctxMiddleware{}
}

// Do injects custom business context variable into context of current request.
func (s *ctxMiddleware) Do(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Passport: user.Passport,
			Nickname: user.Nickname,
		}
	}
	// Continue execution of next middleware.
	r.Middleware.Next()
}
