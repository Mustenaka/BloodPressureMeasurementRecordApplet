package server

import (
	"BloodPressure/pkg/log"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// Ping 用来检查是否程序正常启动
func Ping(baseUrl, port string, maxCount int) error {
	seconds := 1
	url := baseUrl + port + "/ping"
	for i := 0; i < maxCount; i++ {
		resp, err := http.Get(url)
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Infof("等待服务在线, 已等待 %d 秒，最多等待 %d 秒", seconds, maxCount)
		time.Sleep(time.Second * 1)
		seconds++
	}
	return fmt.Errorf("服务启动失败，端口 %s", port)
}

// Ping 用来检查是否程序正常启动
func HttpsPing(baseUrl, port string, maxCount int) error {
	seconds := 1

	// 忽略客户端证书
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	url := baseUrl + port + "/ping"
	// url := "https://www.lyhxxcx.cn/ping"

	for i := 0; i < maxCount; i++ {
		resp, err := client.Get(url)
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Infof("等待服务在线, 已等待 %d 秒，最多等待 %d 秒", seconds, maxCount)
		time.Sleep(time.Second * 1)
		seconds++
	}
	return fmt.Errorf("服务启动失败，端口 %s", port)
}
