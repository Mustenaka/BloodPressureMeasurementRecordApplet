package middleware

import (
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/log"
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 记录每次请求的请求信息和响应信息
func Logger(c *gin.Context) {
	// 请求前
	now := time.Now()
	reqPath := c.Request.URL.Path
	reqId := c.GetString(constant.RequestId)           //
	method := c.Request.Method                         // 请求方法GET POST
	ip := c.ClientIP()                                 // 客户端IP
	requestBody, err := ioutil.ReadAll(c.Request.Body) // 请求body
	if err != nil {
		requestBody = []byte{}
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	log.Info("New request start",
		log.WithPair(constant.RequestId, reqId),
		log.WithPair("host", ip),
		log.WithPair("path", reqPath),
		log.WithPair("method", method),
		log.WithPair("body", string(requestBody)))

	c.Next()
	// 请求后
	latency := time.Since(now)
	log.Info("New request end",
		log.WithPair(constant.RequestId, reqId),
		log.WithPair("host", ip),
		log.WithPair("path", reqPath),
		log.WithPair("cost", latency))
}
