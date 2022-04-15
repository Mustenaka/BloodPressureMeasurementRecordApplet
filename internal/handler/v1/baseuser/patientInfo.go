package baseuser

import (
	"BloodPressure/internal/model"
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	"context"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 添加患者信息
func (uh *BaseUserHandler) AddPatientInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type InfoParam struct {
			RealName                  string `json:"real_name"  binding:"required"`                   // 真实姓名
			Sex                       string `json:"sex"  binding:"required"`                         // 性别
			Birthday                  string `json:"birthday"  binding:"required"`                    // 生日
			Tel                       string `json:"tel"  binding:"required"`                         // 电话号码
			IsMarried                 bool   `json:"is_married"  binding:"required"`                  // 0-未婚、1-已婚
			HbpYears                  int    `json:"hbp_years"  binding:"required"`                   // 高血压患病时间（年）
			Anamnesis                 string `json:"anamnesis"  binding:"required"`                   // 既往病史(对应表格1~12)
			IsSmoking                 bool   `json:"is_smoking"  binding:"required"`                  // 是否吸烟
			SmokingHistory            int    `json:"smoking_history"  binding:"required"`             // 吸烟史（年）
			SmokingDaily              int    `json:"smoking_daily"  binding:"required"`               // 日吸烟数
			IsDrink                   bool   `json:"is_drink"  binding:"required"`                    // 是否饮酒
			DrinkHistory              int    `json:"drink_history"  binding:"required"`               // 饮酒史（年）
			DrinkDaily                int    `json:"drink_daily"  binding:"required"`                 // 每日饮酒量
			PatientHeight             int    `json:"patient_height"  binding:"required"`              // 身高
			PatientWeight             int    `json:"patient_weight"  binding:"required"`              // 体重
			PatientWaistCircumference int    `json:"patient_waist_circumference"  binding:"required"` // 腰围
			PatientChestCircumference int    `json:"patient_chest_circumference"  binding:"required"` // 胸围
			PatientHipCircumference   int    `json:"patient_hip_circumference"  binding:"required"`   // 臀围
			IsTakeChineseMedicine     bool   `json:"is_take_chinese_medicine"  binding:"required"`    // 是否服用中药
			AntihypertensivePlan      string `json:"antihypertensive_plan"  binding:"required"`       // 降压方案
			IsNondrugControlPlan      bool   `json:"is_nondrug_control_plan"  binding:"required"`     // 是否非药物控制手段
			NondrugControlPlan        string `json:"nondrug_control_plan"  binding:"required"`        // 非药物控制手段内容
		}

		// 检验基本结构
		var param InfoParam
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
		var insertInfo *model.PatientInfo
		typeNameOfInfoParam := reflect.TypeOf(param)
		typeNameOfInsertInfo := reflect.TypeOf(insertInfo)

		valueOfInfoParam := reflect.ValueOf(param).Elem()
		valueOfInsertInfo := reflect.ValueOf(insertInfo).Elem()

		for i := 0; i < typeNameOfInfoParam.NumField(); i++ {
			for j := 0; j < typeNameOfInsertInfo.NumField(); j++ {
				if typeNameOfInfoParam.Field(i).Name == typeNameOfInsertInfo.Field(j).Name {
					valueOfInsertInfo.Field(j).Set(valueOfInfoParam.Field(i))
				}
			}
		}

		// 插入uid
		insertInfo.UserId = baseUser.UserId

		// 患者信息插入
		err = uh.pinfoService.Add(context.TODO(), insertInfo)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.PatientInfoErr, "添加患者信息失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "record add successful",
		})
	}
}

// 更新患者信息
func (uh *BaseUserHandler) UpdatePatientInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type InfoParam struct {
			RealName                  string `json:"real_name"  binding:"required"`                   // 真实姓名
			Sex                       string `json:"sex"  binding:"required"`                         // 性别
			Birthday                  string `json:"birthday"  binding:"required"`                    // 生日
			Tel                       string `json:"tel"  binding:"required"`                         // 电话号码
			IsMarried                 bool   `json:"is_married"  binding:"required"`                  // 0-未婚、1-已婚
			HbpYears                  int    `json:"hbp_years"  binding:"required"`                   // 高血压患病时间（年）
			Anamnesis                 string `json:"anamnesis"  binding:"required"`                   // 既往病史(对应表格1~12)
			IsSmoking                 bool   `json:"is_smoking"  binding:"required"`                  // 是否吸烟
			SmokingHistory            int    `json:"smoking_history"  binding:"required"`             // 吸烟史（年）
			SmokingDaily              int    `json:"smoking_daily"  binding:"required"`               // 日吸烟数
			IsDrink                   bool   `json:"is_drink"  binding:"required"`                    // 是否饮酒
			DrinkHistory              int    `json:"drink_history"  binding:"required"`               // 饮酒史（年）
			DrinkDaily                int    `json:"drink_daily"  binding:"required"`                 // 每日饮酒量
			PatientHeight             int    `json:"patient_height"  binding:"required"`              // 身高
			PatientWeight             int    `json:"patient_weight"  binding:"required"`              // 体重
			PatientWaistCircumference int    `json:"patient_waist_circumference"  binding:"required"` // 腰围
			PatientChestCircumference int    `json:"patient_chest_circumference"  binding:"required"` // 胸围
			PatientHipCircumference   int    `json:"patient_hip_circumference"  binding:"required"`   // 臀围
			IsTakeChineseMedicine     bool   `json:"is_take_chinese_medicine"  binding:"required"`    // 是否服用中药
			AntihypertensivePlan      string `json:"antihypertensive_plan"  binding:"required"`       // 降压方案
			IsNondrugControlPlan      bool   `json:"is_nondrug_control_plan"  binding:"required"`     // 是否非药物控制手段
			NondrugControlPlan        string `json:"nondrug_control_plan"  binding:"required"`        // 非药物控制手段内容
		}

		// 检验基本结构
		var param InfoParam
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
		var insertInfo *model.PatientInfo
		typeNameOfInfoParam := reflect.TypeOf(param)
		typeNameOfInsertInfo := reflect.TypeOf(insertInfo)

		valueOfInfoParam := reflect.ValueOf(param).Elem()
		valueOfInsertInfo := reflect.ValueOf(insertInfo).Elem()

		for i := 0; i < typeNameOfInfoParam.NumField(); i++ {
			for j := 0; j < typeNameOfInsertInfo.NumField(); j++ {
				if typeNameOfInfoParam.Field(i).Name == typeNameOfInsertInfo.Field(j).Name {
					valueOfInsertInfo.Field(j).Set(valueOfInfoParam.Field(i))
				}
			}
		}

		// 插入uid
		insertInfo.UserId = baseUser.UserId

		// 患者信息插入
		err = uh.pinfoService.Update(context.TODO(), insertInfo)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.PatientInfoErr, "更新患者信息失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "record add successful",
		})
	}
}

// 获取血压记录
func (uh *BaseUserHandler) GetPatientInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过id找到基本用户
		uid := c.GetUint(constant.UserID)
		baseUser, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
			return
		}

		// 获取血压记录
		records, err := uh.pinfoService.GetById(context.TODO(), baseUser.UserId)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.PatientInfoErr, "患者信息获取失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}
