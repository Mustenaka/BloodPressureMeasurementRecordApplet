package server

import (
	"BloodPressure/pkg/config"
	"BloodPressure/pkg/log"
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// HttpServer 代表当前服务端实例
type HttpServer struct {
	config *config.Config
	f      func()
}

// NewHttpServer 创建server实例
func NewHttpServer(config *config.Config) *HttpServer {
	return &HttpServer{
		config: config,
	}
}

// Router 加载路由，使用侧提供接口，实现侧需要实现该接口
type Router interface {
	Load(engine *gin.Engine)
}

// AppOptions 用来接收应用启动时指定的参数
type AppOptions struct {
	PrintVersion   bool   // 打印版本
	ConfigFilePath string // 配置文件路径
}

// ResolveAppOptions 解析启动参数
func ResolveAppOptions(opt *AppOptions) {
	var printVersion bool
	var configFilePath string
	flag.BoolVar(&printVersion,
		"v",
		false,
		"-v 选项用于控制是否打印当前项目的版本",
	)
	flag.StringVar(&configFilePath,
		"c", "",
		"-c 选项用于指定要使用的配置文件")
	flag.Parse()

	opt.PrintVersion = printVersion
	// opt.ConfigFilePath = configFilePath
	opt.ConfigFilePath = "./config/config.ini"
}

// Run server的启动入口
// 加载路由, 启动HTTP服务
func (s HttpServer) Run(rs ...Router) {
	var wg sync.WaitGroup
	wg.Add(1)

	// 设置gin启动模式，必须在创建gin实例之前
	gin.SetMode(s.config.ServerConfig.Mode)
	g := gin.New()
	s.routerLoad(g, rs...)

	// health check
	log.Debug("进行health check")
	go func() {
		if err := Ping(s.config.ServerConfig.Url, s.config.ServerConfig.Port, s.config.ServerConfig.MaxPingCount); err != nil {
			log.Fatal("server no response")
		}
		log.Infof("server started success! port: %s", s.config.ServerConfig.Port)
	}()

	// 创建http服务
	srv := http.Server{
		Addr:    s.config.ServerConfig.Port,
		Handler: g,
	}
	if s.f != nil {
		srv.RegisterOnShutdown(s.f)
	}

	// graceful shutdown
	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	// 检测服务是否shutdown了
	go func() {
		<-sgn
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Errorf("server shutdown err %v \n", err)
		}
		wg.Done()
	}()

	// 监听服务
	err := srv.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Errorf("server start failed on port %s", s.config.ServerConfig.Port)
			return
		}
	}
	wg.Wait()
	log.Infof("server stop on port %s", s.config.ServerConfig.Port)
}

// Run server的启动入口
// 加载路由, 启动HTTPS服务
func (s HttpServer) RunTLS(rs ...Router) {
	var wg sync.WaitGroup
	wg.Add(1)

	// 设置gin启动模式，必须在创建gin实例之前
	gin.SetMode(s.config.ServerConfig.Mode)
	g := gin.New()
	s.routerLoad(g, rs...)

	// health check
	log.Debug("进行health check")
	go func() {
		if err := HttpsPing(s.config.ServerConfig.Url, s.config.ServerConfig.Port, s.config.ServerConfig.MaxPingCount); err != nil {
			log.Fatal("server no response")
		}
		log.Infof("server started success! port: %s", s.config.ServerConfig.Port)
	}()

	// 创建http服务
	srv := http.Server{
		Addr:    s.config.ServerConfig.Port,
		Handler: g,
		// TLSConfig: &tls.Config{},
	}

	if s.f != nil {
		srv.RegisterOnShutdown(s.f)
	}

	// graceful shutdown
	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT)

	// 检测服务是否shutdown了
	go func() {
		<-sgn
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Errorf("server shutdown err %v \n", err)
		}
		wg.Done()
	}()

	// 监听服务 https
	err := srv.ListenAndServeTLS("ssl/www.lyhxxcx.cn_public.crt", "ssl/www.lyhxxcx.cn.key")
	// err := srv.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Errorf("server start failed on port %s", s.config.ServerConfig.Port)
			return
		}
	}
	wg.Wait()
	log.Infof("server stop on port %s", s.config.ServerConfig.Port)
}

// RouterLoad 加载自定义路由
func (s *HttpServer) routerLoad(g *gin.Engine, rs ...Router) *HttpServer {
	for _, r := range rs {
		r.Load(g)
	}
	return s
}

// RegisterOnShutdown 注册shutdown后的回调处理函数，用于清理资源
func (s *HttpServer) RegisterOnShutdown(_f func()) {
	s.f = _f
}
