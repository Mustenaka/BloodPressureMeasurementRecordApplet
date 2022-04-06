package baseuser

import (
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	"context"

	"github.com/gin-gonic/gin"
)

// 记录血压
func (uh *BaseUserHandler) RecordBp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RegisterParam struct {
			Low  int `json:"Low"  binding:"required"`
			High int `json:"High"  binding:"required"`
		}

		// 检验基本结构
		var param RegisterParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		// 通过id找到基本用户
		uid := c.GetUint(constant.UserID)
		baseUser, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
			return
		}

		// 进行血压记录
		err = uh.bprSrv.AddById(context.TODO(), baseUser.UserId, param.Low, param.High)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.BPRecordErr, "添加血压记录失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "record successful",
		})
	}
}
