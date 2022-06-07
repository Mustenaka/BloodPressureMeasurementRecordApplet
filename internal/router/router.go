package router

import (
	"BloodPressure/internal/handler/copyright"
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

	// Copyright - 版权信息彩蛋
	g.GET("/copyright", copyright.Copyright())

	// login & wechat login
	g.POST("/login", r.uh.Login())
	g.POST("/wxlogin", r.uh.WeLogin())

	// wechat register
	g.POST("/wxregister", r.uh.WeRegister())

	// user group (wechat)
	ug := g.Group("/v1/user", middleware.AuthToken())
	{
		// index info
		ug.GET("", r.uh.GetBaseUserInfo())

		// login(relogin)
		ug.POST("/wxlogin", r.uh.WeLogin())

		// 用户基本信息
		ug.PUT("/userpassword", r.uh.UpdateUserPassword())
		ug.PUT("/user", r.uh.UpdateUserDetail())
		ug.GET("/user", r.uh.GetBaseUserInfo())
		// 禁止个人用户删除自己信息

		// 血压测量记录功能
		ug.POST("/bprecord", r.uh.RecordBp())
		ug.GET("/bprecord", r.uh.GetRecordBp())

		// 用户治疗方案
		ug.POST("/treatmentplan", r.uh.AddPlan())
		ug.GET("/treatmentplan", r.uh.GetPlans())

		// 患者信息记录
		ug.POST("/patientinfo", r.uh.AddPatientInfo())
		ug.POST("/wxpatientinfo", r.uh.WxUpdatePatientInfo()) // 微信用接口自动更新或者添加
		ug.GET("/patientinfo", r.uh.GetPatientInfo())
		ug.PUT("/patientinfo", r.uh.UpdatePatientInfo())

		// 用户检查报告

		// 用户上传照片 - 先抛弃，啥也不做
		ug.POST("/upload", r.uh.DiscardUserUploadedPhotos())
	}

	// admin group (administrator)
	ag := g.Group("/v1/admin", middleware.AuthToken())
	{
		// index info
		ag.GET("", r.uh.GetBaseUserInfo())
	}
}
