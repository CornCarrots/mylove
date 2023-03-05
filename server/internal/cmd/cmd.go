package cmd

import (
	"context"
	"server/internal/consts"
	"server/internal/mw"
	"server/utility/resp"

	"github.com/gogf/gf/v2/util/gmode"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"

	"server/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)
			// OpenApi自定义信息
			oai.Info.Title = `API Reference`
			oai.Config.CommonResponse = resp.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

			// 静态目录设置
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			if uploadPath == "" {
				g.Log().Fatal(ctx, "文件上传配置路径不能为空")
			}
			s.AddStaticPath("/upload", uploadPath)

			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}
			s.Use(
				mw.LogMiddleware,
				ghttp.MiddlewareHandlerResponse,
				ghttp.MiddlewareCORS,
			)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					mw.CtxMiddleware,
				)
				group.Bind(
					controller.Hello,
					controller.UserController,
					controller.NoteController,
				)
				// 权限控制路由

				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						mw.AuthMiddleware,
					)
					group.Bind(
						controller.UserController.Profile,
						controller.NoteController,
					)
				})

			})
			// Custom enhance API document.
			enhanceOpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "zero",
			URL:  "",
		},
	}

	openapi.Tags = &goai.Tags{
		{Name: consts.OpenAPITagNameUser},
		{Name: consts.OpenAPITagNameNote},
	}
}
