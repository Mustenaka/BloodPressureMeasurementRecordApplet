package baseuser

import (
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"

	"github.com/gin-gonic/gin"
)

// 添加我的检验指标 - BNP
func (uh *BaseUserHandler) AddTiBnps() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			Data     int    `json:"data"  binding:"required"`
			CreateAt string `json:"create_at" `
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

		// 特判是否输入了时间，没有则记录当前时间
		if param.CreateAt == "" {
			param.CreateAt = timeconvert.NowDateTimeString()
		}

		// 进行血压记录
		err = uh.tbnpSrv.AddByIdWithTime(context.TODO(), baseUser.UserId, param.Data, param.CreateAt)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.TestIndicatorErr, "BNP检验指标写入错误"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "tongue datail add successful",
		})
	}
}

// 获取我的检验指标 - BNP
func (uh *BaseUserHandler) GetTiBnps() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			// 数量限制
			Limit int `form:"limit"`
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
		records, err := uh.tbnpSrv.GetByIdLimit(context.TODO(), baseUser.UserId, param.Limit)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.TestIndicatorErr, "获取我的检验指标失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}

// 添加我的检验指标 - 肌酐
func (uh *BaseUserHandler) AddTiCreatinines() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			Data     int    `json:"data"  binding:"required"`
			CreateAt string `json:"create_at" `
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

		// 特判是否输入了时间，没有则记录当前时间
		if param.CreateAt == "" {
			param.CreateAt = timeconvert.NowDateTimeString()
		}

		// 进行血压记录
		err = uh.tcreatinineSrv.AddByIdWithTime(context.TODO(), baseUser.UserId, param.Data, param.CreateAt)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.TestIndicatorErr, "肌酐检验指标写入错误"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "tongue datail add successful",
		})
	}
}

// 获取我的检验指标 - 肌酐
func (uh *BaseUserHandler) GetTiCreatinines() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			// 数量限制
			Limit int `form:"limit"`
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
		records, err := uh.tcreatinineSrv.GetByIdLimit(context.TODO(), baseUser.UserId, param.Limit)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.TestIndicatorErr, "获取我的检验指标失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}
