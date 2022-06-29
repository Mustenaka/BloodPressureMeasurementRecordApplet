package baseuser

import (
	"BloodPressure/internal/model"
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/log"
	"BloodPressure/pkg/response"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 添加我的体检报告 - 24小时血压检测
func (uh *BaseUserHandler) AddMr24HoursBpr() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			DayHigh   int    `json:"day_high"`
			DayLow    int    `json:"day_low"`
			NightHigh int    `json:"night_high"`
			NightLow  int    `json:"night_low"`
			CreateAt  string `json:"create_at" `
			// 血压最大记录
			MaxHigh     int    `json:"max_high"`
			MaxHighTime string `json:"max_high_time"`
			MaxLow      int    `json:"max_low"`
			MaxLowTime  string `json:"max_low_time"`
			// 血压最低记录
			MinHigh     int    `json:"min_high"`
			MinHighTime string `json:"min_high_time"`
			MinLow      int    `json:"min_low"`
			MinLowTime  string `json:"min_low_time"`
		}

		// 检验基本结构
		var param RecordParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		log.Debugf("param: %+v", param)

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

// 获取我的检验指标 - 24小时血压检测
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

// 添加我的体检报告 - 24小时ECG心电图检测
func (uh *BaseUserHandler) AddMr24HoursEcg() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			Data     string `json:"data"`
			Average  int    `json:"average"` // 平均心率（次/分）
			CreateAt string `json:"create_at" `
		}

		// 检验基本结构
		var param RecordParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		log.Debugf("param: %+v", param)

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
		err = uh.m24ecgSrv.AddByIdWithTime(context.TODO(), baseUser.UserId, param.Data, param.Average, param.CreateAt)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.MedicalRecordErr, "24小时动态ECG写入错误"), nil)
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

// 获取我的检验指标 - 24小时ECG心电图检测
func (uh *BaseUserHandler) GetMr24HoursEcg() gin.HandlerFunc {
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
		records, err := uh.m24ecgSrv.GetByIdLimit(context.TODO(), baseUser.UserId, param.Limit)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.MedicalRecordErr, "获取我的检验报告失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}

// 添加我的体检报告 - ECG心电图检测
func (uh *BaseUserHandler) AddMrEcg() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			Data     string `json:"data"`
			CreateAt string `json:"create_at" `
		}

		// 检验基本结构
		var param RecordParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		log.Debugf("param: %+v", param)

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
		err = uh.mecgSrv.AddByIdWithTime(context.TODO(), baseUser.UserId, param.Data, param.CreateAt)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.MedicalRecordErr, "ECG写入错误"), nil)
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

// 获取我的检验指标 - ECG心电图检测
func (uh *BaseUserHandler) GetMrEcg() gin.HandlerFunc {
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
		records, err := uh.mecgSrv.GetByIdLimit(context.TODO(), baseUser.UserId, param.Limit)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.MedicalRecordErr, "获取我的检验报告失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}

// 添加我的体检报告 - 心超检测
func (uh *BaseUserHandler) AddMrechocardiographys() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			Ef       int    `json:"ef"`
			Lvidd    int    `json:"lvidd"`
			Lvids    int    `json:"lvids"`
			CreateAt string `json:"create_at" `
		}

		// 检验基本结构
		var param RecordParam
		if err := c.ShouldBind(&param); err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "存在必要信息未填写"), nil)
			return
		}

		log.Debugf("param: %+v", param)

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
		err = uh.mcreatinineSrv.AddByIdWithTime(context.TODO(), baseUser.UserId, param.Ef, param.Lvidd, param.Lvids, param.CreateAt)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.MedicalRecordErr, "ECG写入错误"), nil)
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

// 获取我的检验指标 - 心超检测
func (uh *BaseUserHandler) GetMrechocardiographys() gin.HandlerFunc {
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
		records, err := uh.mcreatinineSrv.GetByIdLimit(context.TODO(), baseUser.UserId, param.Limit)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.MedicalRecordErr, "获取我的检验报告失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}
