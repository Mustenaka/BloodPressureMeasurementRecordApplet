package model

import (
	"BloodPressure/global"
	"BloodPressure/tools"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
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

	db, err = gorm.Open(mysql.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect successful!")
}

// 插入数据，测试
func Create() {
	user := BaseUser{
		// UserId:     11,
		OpenId:     tools.RandomString(16),
		UserName:   "张四",
		Tel:        "18278362206",
		Email:      "asdqw123@outlook.com",
		Permission: 3,
		LastTime:   time.Now().Format("2006-01-02 15:04:05"),
	}
	db.Create(&user)
}
