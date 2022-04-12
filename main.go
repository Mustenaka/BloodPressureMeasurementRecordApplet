package main

import (
	"BloodPressure/internal/handler/v1/baseuser"
	"BloodPressure/internal/repo/mysql"
	"BloodPressure/internal/router"
	"BloodPressure/internal/service"
	"BloodPressure/pkg/config"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	"BloodPressure/pkg/version"
	"BloodPressure/server"
)

// 初始化路由
func initRouter(ds db.IDataSource) server.Router {
	// 用户
	userRepo := mysql.NewBaseUserRepo(ds)
	userService := service.NewBaseUserService(userRepo)
	// 血压记录
	bprRepo := mysql.NewPatientBpRecordRepo(ds)
	bprService := service.NewPBPRecordService(bprRepo)
	// 生成Handler并且传递至Router服务
	userHandler := baseuser.NewBaseUserHandler(userService, bprService)
	routerRouter := router.NewRouter(userHandler)
	return routerRouter
}

// 获取路由
func getRouters(ds db.IDataSource) []server.Router {
	rts := make([]server.Router, 0)
	rt := initRouter(ds)

	if rt != nil {
		rts = append(rts, rt)
	}
	return rts
}

// 启动应用代码
func RunProgram() {
	// 解析服务器配置参数
	appOpt := &server.AppOptions{}
	server.ResolveAppOptions(appOpt)
	if appOpt.PrintVersion {
		version.PrintVersion()
	}

	// 加载配置文件
	c := config.Load(appOpt.ConfigFilePath)
	log.InitLogger(&c.LogConfig, c.BasicinfoConfig.AppName) // 日志

	// 创建数据库链接，使用默认的实现方式
	ds := db.NewDefaultMysql(c.DBConfig)
	routers := getRouters(ds)

	// 创建HTTPServer
	srv := server.NewHttpServer(config.GlobalConfig)
	srv.RegisterOnShutdown(func() {
		if ds != nil {
			ds.Close()
		}
	})

	// 启动服务
	srv.Run(routers...)
}

func main() {
	// example.ExampleRun()
	RunProgram()
}
