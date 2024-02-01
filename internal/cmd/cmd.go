package cmd

import (
	"01-practice/internal/controller"
	"01-practice/internal/controller/loginAuth"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func MainRun() {
	var ctx context.Context
	s := g.Server()
	v1 := s.Group("/api/v1")
	// MiddlewareCORS 允许跨域
	v1.Middleware(ghttp.MiddlewareHandlerResponse, MiddlewareCORS)
	v1.GET("/", controller.GetSkuSceneController{})     // 查询分类列表
	v1.POST("/goods", controller.GoodsListController{}) // 查询商品列表
	v1.POST("/", controller.SignInController{})         // 注册

	// 中间件一下的路由将会 token 验证
	loginAuth.NewAuthToken().Login()
	err := loginAuth.BfToken.Middleware(context.Background(), v1)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	s.Run()
}

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
