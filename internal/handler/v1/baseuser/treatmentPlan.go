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
func (uh *BaseUserHandler) AddPlan() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			Plan string `json:"plan"  binding:"required"`
			Note string `json:"note"`
		}

		// 检验基本结构
		var param RecordParam
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
		err = uh.trplanSrc.AddById(context.TODO(), baseUser.UserId, param.Plan, param.Note)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.TreatPlanErr, "添加治疗记录失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "plan add successful",
		})
	}
}

// 获取治疗方案记录
func (uh *BaseUserHandler) GetPlans() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			// 数量限制
			LimitCount int `form:"limit_count"`
		}

		// 检验基本结构
		var param RecordParam
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

		// 获取治疗方案记录
		records, err := uh.trplanSrc.GetByIdLimit(context.TODO(), baseUser.UserId, param.LimitCount)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.TreatPlanErr, "治疗方案记录获取失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}
