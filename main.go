package main

import (
	"BloodPressure/internal/handler/v1/baseuser"
	"BloodPressure/internal/repo/mysql"
	"BloodPressure/internal/router"
	"BloodPressure/internal/service"
	"BloodPressure/pkg/config"
	"BloodPressure/pkg/copyright"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	"BloodPressure/pkg/version"
	"BloodPressure/server"
)

// 初始化路由 <- 随着项目越来越大，这里一定要IOT处理，尽快熟悉dig库运行
func initRouter(ds db.IDataSource) server.Router {
	// 用户
	userRepo := mysql.NewBaseUserRepo(ds)
	userService := service.NewBaseUserService(userRepo)
	// 血压记录
	bprRepo := mysql.NewPatientBpRecordRepo(ds)
	bprService := service.NewPBPRecordService(bprRepo)
	// 治疗方案记录
	trplanRepo := mysql.NewTreatmentPlanRepo(ds)
	trplanService := service.NewTreatmentPlanService(trplanRepo)
	// 患者信息记录
	pinfoRepo := mysql.NewPatientInfoRepo(ds)
	pinfoService := service.NewPatientInfoService(pinfoRepo)
	// 舌苔迈向记录
	tongueRepo := mysql.NewTongueDetailRepo(ds)
	tongueService := service.NewTongueDetailService(tongueRepo)
	// 24小时 bpr记录
	bpr24Repo := mysql.NewMedicalReport24HoursbprRepo(ds)
	bpr24Service := service.NewMedicalReport24HoursbprService(bpr24Repo)
	// 24小时 ecg记录
	ecg24Repo := mysql.NewMedicalReport24hoursecgRepo(ds)
	ecg24Service := service.NewMedicalReport24hoursecgService(ecg24Repo)
	// ecg记录
	ecgRepo := mysql.NewMedicalReportEcgRepo(ds)
	ecgService := service.NewMedicalReportEcgService(ecgRepo)
	// 心超记录
	echoRepo := mysql.NewMedicalReportEchocardiographyRepo(ds)
	echoService := service.NewMedicalReportEchocardiographyService(echoRepo)
	// BNP记录
	bnpRepo := mysql.NewTestIndicatorBnpRepo(ds)
	bnpService := service.NewTestIndicatorBnpService(bnpRepo)
	// 肌酐参数
	creatinineRepo := mysql.NewTestIndicatorCreatinine(ds)
	creatinineService := service.NewTestIndicatorCreatinineService(creatinineRepo)

	// 生成Handler并且传递至Router服务
	userHandler := baseuser.NewBaseUserHandler(
		userService, bprService, trplanService, pinfoService, tongueService,
		bpr24Service, ecg24Service, ecgService, echoService,
		bnpService, creatinineService)
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

	// 加载配置文件,日志文件
	c := config.Load(appOpt.ConfigFilePath)
	copyright.GetInstance().LoadCopyright(c.BasicinfoConfig)
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

	// 启动HTTPS服务
	srv.RunTLS(routers...)
}

func main() {
	// example.ExampleRun()
	RunProgram()
}
