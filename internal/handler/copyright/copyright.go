package copyright

import (
	"BloodPressure/pkg/log"
	"BloodPressure/pkg/version"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Copyright ping服务器状态
func Copyright() gin.HandlerFunc {
	return func(c *gin.Context) {
		// goVersion := runtime.Version()
		AppVersion := version.AppVersion

		log.Debug("result", log.WithPair("result", AppVersion))

		c.String(http.StatusOK, AppVersion)
	}
}
