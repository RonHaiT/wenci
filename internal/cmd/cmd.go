package cmd

import (
	"context"

	"wenci/internal/controller/account"
	"wenci/internal/controller/users"
	"wenci/internal/controller/words"
	"wenci/internal/logic/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// OpenAPI/Swagger: 显式声明鉴权格式（swagger 才能展示需要认证与 Authorize）
			openapi := s.GetOpenApi()
			openapi.Components.SecuritySchemes = goai.SecuritySchemes{
				// 标准 Bearer Token：Authorization: Bearer <jwt>
				"AuthToken": goai.SecuritySchemeRef{
					Value: &goai.SecurityScheme{
						Type:   "http",
						Scheme: "bearer",
					},
				},
			}

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Bind(
						users.NewV1(),
					)
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.Bind(
							account.NewV1(),
							words.NewV1(),
						)
					})
				})

			})
			s.Run()
			return nil
		},
	}
)
