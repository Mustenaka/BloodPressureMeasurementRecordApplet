package copyright

import (
	"BloodPressure/pkg/copyright"
	"BloodPressure/pkg/response"

	"github.com/gin-gonic/gin"
)

// Copyright 服务器版权信息展示状态
func Copyright() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.JSON(c, nil, copyright.GetInstance().GetCopyright())
	}
}
