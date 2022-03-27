package router

import (
	"BloodPressure/internal/handler/ping"
	"BloodPressure/internal/handler/v1/baseuser"
	"BloodPressure/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// router Router路由接口的默认实现
type router struct {
	uh *baseuser.BaseUserHandler
}

// 新建路由
func NewRouter(_uh *baseuser.BaseUserHandler) *router {
	return &router{
		uh: _uh,
	}
}

// Load 加载中间件和路由信息
func (r *router) Load(g *gin.Engine) {
	// 注册中间件
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache())
	g.Use(middleware.Options())
	g.Use(middleware.Secure())
	g.Use(middleware.RequestId())
	g.Use(middleware.Logger)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404 not found!")
	})

	// ping server - 测试服务器通畅
	g.GET("/ping", ping.Ping())

	// login
	g.POST("/login", r.uh.Login())
	// loginWithOpenid - wechat user login.
	g.POST("/wechatlogin", r.uh.LoginWithOpenid())

	// register
	g.POST("/register", r.uh.Register())

	// user group
	ug := g.Group("/v1/user", middleware.AuthToken())
	{
		ug.GET("", r.uh.GetBaseUserInfo())
		// login
		ug.POST("/login", r.uh.Login())
	}
}
