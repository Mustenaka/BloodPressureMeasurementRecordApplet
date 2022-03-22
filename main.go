package main

import (
	"BloodPressure/internal/model"
	"BloodPressure/pkg/config"
	"BloodPressure/pkg/db"
	"BloodPressure/pkg/log"
	"BloodPressure/pkg/version"
	"BloodPressure/server"
	"fmt"
)

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
	ds := db.NewDefaultMysql(c.DBConfig)                    // 创建数据库链接，使用默认的实现方式
	var users []model.BaseUser
	if err := ds.Master().Where(&model.BaseUser{UserName: "李翠花"}).Find(&users); err.Error != nil {
		// 错误处理
		fmt.Println("没有找到该数据333")
	}
	for _, value := range users {
		fmt.Println(value.UserId, value.UserName)
	}
}

func main() {
	RunProgram()
}
