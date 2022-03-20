package server

import "BloodPressure/pkg/global"

// 当前服务器实例
type HttpServer struct {
	mode         string
	port         string
	maxPingCount int
	jwtSecret    string
	f            func()
}

// NewHttpServer 创建server实例
func NewHttpServer() *HttpServer {
	conf := global.GetInstance()
	return &HttpServer{
		mode:         conf.GetConfigValue("server", "mode"),
		port:         conf.GetConfigValue("server", "port"),
		maxPingCount: conf.GetConfigValueInt("server", "max-ping"),
		jwtSecret:    conf.GetConfigValue("server", "jwt-secret"),
	}
}
