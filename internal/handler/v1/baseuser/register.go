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
			Username  string `json:"username" binding:"required"`
			Password  string `json:"password" binding:"required"`
			OpenId    string `json:"openid"   binding:"required"`
			RealName  string `json:"realname"   binding:"required"`
			Telephone string `json:"telephone"   binding:"required"`
			Email     string `json:"email"   binding:"required"`
		}

		// 检验基本结构
		var param RegisterParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

	}
}
