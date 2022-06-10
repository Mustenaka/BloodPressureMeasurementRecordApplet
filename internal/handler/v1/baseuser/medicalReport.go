package baseuser

import (
	"BloodPressure/internal/model"
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 添加我的检验指标 - BNP
func (uh *BaseUserHandler) AddMr24HoursBpr() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			DayHigh   int    `form:"day_high"`
			DayLow    int    `form:"day_low"`
			NightHigh int    `form:"night_high"`
			NightLow  int    `form:"night_low"`
			CreateAt  string `json:"create_at" `
			// 血压最大记录
			MaxHigh     int    `form:"max_high"`
			MaxHighTime string `form:"max_high_time"`
			MaxLow      int    `form:"max_low"`
			MaxLowTime  string `form:"max_low_time"`

			MinHigh     int    `form:"min_high"`
			MinHighTime string `form:"min_high_time"`
			MinLow      int    `form:"min_low"`
			MinLowTime  string `form:"min_low_time"`
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

		// 利用reflect反射构造插入结构（数据字段太多了）
		objOfInfo := &model.MedicalReport24hoursbpr{}
		objOfParam := &param

		typeNameOfInsertInfo := reflect.TypeOf(*objOfInfo)
		typeNameOfInfoParam := reflect.TypeOf(*objOfParam) // 获取type

		valueOfInsertInfo := reflect.ValueOf(objOfInfo).Elem()
		valueOfInfoParam := reflect.ValueOf(objOfParam).Elem() // 获取value

		for i := 0; i < typeNameOfInfoParam.NumField(); i++ {
			for j := 0; j < typeNameOfInsertInfo.NumField(); j++ {
				if typeNameOfInfoParam.Field(i).Name == typeNameOfInsertInfo.Field(j).Name {
					valueOfInsertInfo.Field(j).Set(valueOfInfoParam.Field(i))
				}
			}
		}

		// 设置uid
		objOfInfo.UserId = baseUser.UserId

		// 特判是否输入了时间，没有则记录当前时间
		if param.CreateAt == "" {
			param.CreateAt = timeconvert.NowDateTimeString()
		}

		// 进行血压记录
		err = uh.m24bprSrv.AddById(context.TODO(), baseUser.UserId, objOfInfo)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.MedicalRecordErr, "24小时动态血压写入错误"), nil)
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
func (uh *BaseUserHandler) GetMr24HoursBpr() gin.HandlerFunc {
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
		records, err := uh.m24bprSrv.GetByIdLimit(context.TODO(), baseUser.UserId, param.Limit)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.MedicalRecordErr, "获取我的检验报告失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}
