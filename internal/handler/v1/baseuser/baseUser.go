package baseuser

import (
	"BloodPressure/internal/service"
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	"context"

	"github.com/gin-gonic/gin"
)

// BaseUserHandler 用户业务handler
type BaseUserHandler struct {
	userSrv service.BaseUserService
}

// 新建一个handler
func NewBaseUserHandler(_userSrv service.BaseUserService) *BaseUserHandler {
	return &BaseUserHandler{
		userSrv: _userSrv,
	}
}

func (uh *BaseUserHandler) GetBaseUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetUint(constant.UserID)
		user, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
		} else {
			response.JSON(c, nil, user)
		}
	}
}
