package config

import (
	"log"
	"testing"
)

func TestConfig(t *testing.T) {
	c := Load("./config_back.ini")

	// Basicinfo
	log.Println("Load Basicinfo:", c.BasicinfoConfig.AppName)
	log.Println("Load Basicinfo:", c.BasicinfoConfig.Author)
	log.Println("Load Basicinfo:", c.BasicinfoConfig.AppCompany)
	log.Println("Load Basicinfo:", c.BasicinfoConfig.Version)
	log.Println("Load Basicinfo:", c.BasicinfoConfig.Copyright)

	// Server Config
	log.Println("Load ServerConfig:", c.ServerConfig.Mode)
	log.Println("Load ServerConfig:", c.ServerConfig.Port)
	log.Println("Load ServerConfig:", c.ServerConfig.Url)
	log.Println("Load ServerConfig:", c.ServerConfig.MaxPingCount)
	log.Println("Load ServerConfig:", c.ServerConfig.JwtSecret)

	// DBConfig
	log.Println("Load DBConfig:", c.DBConfig.Dbname)
	log.Println("Load DBConfig:", c.DBConfig.Host)
	log.Println("Load DBConfig:", c.DBConfig.Port)
	log.Println("Load DBConfig:", c.DBConfig.Username)
	log.Println("Load DBConfig:", c.DBConfig.Password)
	log.Println("Load DBConfig:", c.DBConfig.MaximumPoolSize)
	log.Println("Load DBConfig:", c.DBConfig.MaximumIdleSize)
	log.Println("Load DBConfig:", c.DBConfig.LogMode)

	// Log Config
	log.Println("Load LogConfig:", c.LogConfig.Level)
	log.Println("Load LogConfig:", c.LogConfig.FileName)
	log.Println("Load LogConfig:", c.LogConfig.TimeFormat)
	log.Println("Load LogConfig:", c.LogConfig.MaxSize)
	log.Println("Load LogConfig:", c.LogConfig.MaxBackups)
	log.Println("Load LogConfig:", c.LogConfig.MaxAge)
	log.Println("Load LogConfig:", c.LogConfig.Compress)
	log.Println("Load LogConfig:", c.LogConfig.LocalTime)
	log.Println("Load LogConfig:", c.LogConfig.Console)
	t.Log(c.BasicinfoConfig.AppName)
}
