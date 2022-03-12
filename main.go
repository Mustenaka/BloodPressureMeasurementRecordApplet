package main

import (
	"BloodPressure/global"
	"fmt"
)

// Test hello
func RunProgram() {
	var fileName = "./config/config.ini"

	// var dadtabase = "database"
	// username := conf.GetValue(dadtabase, "username")
	// password := conf.GetValue(dadtabase, "password")
	// hostname := conf.GetValue(dadtabase, "hostname")

	// fmt.Println(version, author, appName, appCompany, copyright)
	// fmt.Println(username, password, hostname)

	conf := global.GetInstance(fileName)
	fmt.Println(conf.GetConfigValue("basicinfo", "version"))

	ExampleRun()
}

func main() {
	fmt.Println("Beginning the Program!")

	RunProgram()
}
