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
	userSrv   service.BaseUserService        // 用户服务
	bprSrv    service.PatientBpRecordService // 血压相关记录服务
	trplanSrc service.TreatmentPlanService   // 治疗方案服务
}

// 新建一个handler
func NewBaseUserHandler(_userSrv service.BaseUserService, _bprSrv service.PatientBpRecordService) *BaseUserHandler {
	return &BaseUserHandler{
		userSrv: _userSrv,
		bprSrv:  _bprSrv,
	}
}

// 获取用户信息
func (uh *BaseUserHandler) GetBaseUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetUint(constant.UserID)
		// 通过id找到基本用户
		baseUser, err := uh.userSrv.GetById(context.TODO(), uid)

		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
		} else {
			response.JSON(c, nil, baseUser)
		}
	}
}
