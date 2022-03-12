package main

import (
	"fmt"

	"github.com/widuu/goini"
)

// Test hello
func TestHello() {
	var fileName = "./config/config.ini"

	conf := goini.SetConfig(fileName)

	var basicinfo = "basicinfo"
	version := conf.GetValue(basicinfo, "version")
	author := conf.GetValue(basicinfo, "author")
	appName := conf.GetValue(basicinfo, "appName")
	appCompany := conf.GetValue(basicinfo, "appCompany")
	copyright := conf.GetValue(basicinfo, "copyright")

	var dadtabase = "database"
	username := conf.GetValue(dadtabase, "username")
	password := conf.GetValue(dadtabase, "password")
	hostname := conf.GetValue(dadtabase, "hostname")

	fmt.Println(version, author, appName, appCompany, copyright)
	fmt.Println(username, password, hostname)
}

func main() {
	fmt.Println("Beginning the Program!")

	TestHello()
}
