package model

import (
	"BloodPressure/global"
	"fmt"
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

	fmt.Println(dsn.String())

	// 链接数据库
	DB, err = gorm.Open(mysql.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect mysql successful!")

	// 创建数据库连接池
	sqlDB, err := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic("Connection pool init failed.")
	}
	fmt.Println("Connection pool init successful!")
}

// // 插入一个记录
// func Create(value interface{}) (tx *gorm.DB) {
// 	return DB.Create(value)
// }

// // 查询第一个记录(主键升序)
// func First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
// 	return DB.First(dest, conds)
// }

// // 查询一条记录(无指定排序字段)
// func Take(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
// 	return DB.Take(dest, conds)
// }

// // 查询最后一条记录(逐渐降序)
// func Last(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
// 	return DB.Last(dest, conds)
// }
