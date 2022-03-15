package model

import (
	"BloodPressure/pkg/global"
	"BloodPressure/pkg/log"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// 初始化，建立链接
func init() {
	conf := global.GetInstance()

	// 读取ini配置文件获取mysql链接配置
	// root:jx@#ww4D@tcp(localhost:3306)/HighBloodDB?charset=utf8&parseTime=True&loc=Local
	var selection string = "database"
	var dsn strings.Builder
	dsn.WriteString(conf.GetConfigValue(selection, "username") + ":")
	dsn.WriteString(conf.GetConfigValue(selection, "password") + "@")
	dsn.WriteString(conf.GetConfigValue(selection, "function") + "(")
	dsn.WriteString(conf.GetConfigValue(selection, "hostname") + ":")
	dsn.WriteString(conf.GetConfigValue(selection, "port") + ")/")
	dsn.WriteString(conf.GetConfigValue(selection, "dbname") + "?")
	dsn.WriteString("charset=" + conf.GetConfigValue(selection, "charset") + "&")
	dsn.WriteString("parseTime=" + conf.GetConfigValue(selection, "parseTime") + "&")
	dsn.WriteString("loc=" + conf.GetConfigValue(selection, "loc"))

	// 链接数据库
	DB, err = gorm.Open(mysql.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		log.Panic("Database connect", log.WithPair("message", "falied"))
	}
	log.Info("Database connect", log.WithPair("message", "successful!"))

	// 创建数据库连接池
	sqlDB, err := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Panic("Connection pool init", log.WithPair("message", "falied"))
	}
	log.Info("Connection pool init", log.WithPair("message", "successful!"))
}
