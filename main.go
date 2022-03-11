package main

import (
	"fmt"

	"github.com/widuu/goini"
)

// Test hello
func TestHello() {
	var fileName = "./config/config.ini"

	conf := goini.SetConfig(fileName)

	version := conf.GetValue("basicinfo", "version")

	fmt.Println(version)

}

func main() {
	fmt.Println("Beginning the Program!")

	TestHello()
}
