package mw

import (
	"server/internal/model"
	"server/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CtxMiddleware(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			UserId:   user.UserId,
			Passport: user.Passport,
			Nickname: user.Nickname,
		}
	}
	// Continue execution of next middleware.
	r.Middleware.Next()
}
