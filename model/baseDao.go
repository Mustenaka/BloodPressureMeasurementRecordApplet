package model

import (
	"BloodPressure/global"
)

func Connect() {
	var fileName = "./config/config.ini"
	conf := global.GetInstance(fileName)
	conf.
		gorm.Open(mysql)
}
