package baseuser

import (
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"

	"github.com/gin-gonic/gin"
)

func (uh *BaseUserHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RegisterParam struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		// 检验基本结构
		var param RegisterParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "用户名和密码不能为空"), nil)
			return
		}

	}
}
