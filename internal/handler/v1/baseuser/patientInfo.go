package baseuser

import (
	"BloodPressure/internal/model"
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/log"
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
			RealName                  string `json:"real_name,omitempty"`                     // 真实姓名
			Sex                       string `json:"sex,omitempty"`                           // 性别
			Birthday                  string `json:"birthday,omitempty" `                     // 生日
			Tel                       string `json:"tel,omitempty" `                          // 电话号码
			IsMarried                 bool   `json:"is_married,omitempty"  `                  // 0-未婚、1-已婚
			HbpYears                  int    `json:"hbp_years,omitempty"  `                   // 高血压患病时间（年）
			Anamnesis                 string `json:"anamnesis,omitempty"  `                   // 既往病史(对应表格1~12)
			IsSmoking                 bool   `json:"is_smoking,omitempty"  `                  // 是否吸烟
			SmokingHistory            int    `json:"smoking_history,omitempty" `              // 吸烟史（年）
			SmokingDaily              int    `json:"smoking_daily,omitempty"  `               // 日吸烟数
			IsDrink                   bool   `json:"is_drink,omitempty"  `                    // 是否饮酒
			DrinkHistory              int    `json:"drink_history,omitempty"  `               // 饮酒史（年）
			DrinkDaily                int    `json:"drink_daily,omitempty"`                   // 每日饮酒量
			PatientHeight             int    `json:"patient_height,omitempty" `               // 身高
			PatientWeight             int    `json:"patient_weight,omitempty" `               // 体重
			PatientWaistCircumference int    `json:"patient_waist_circumference,omitempty" `  // 腰围
			PatientChestCircumference int    `json:"patient_chest_circumference,omitempty"  ` // 胸围
			PatientHipCircumference   int    `json:"patient_hip_circumference,omitempty"  `   // 臀围
			IsTakeChineseMedicine     bool   `json:"is_take_chinese_medicine,omitempty" `     // 是否服用中药
			AntihypertensivePlan      string `json:"antihypertensive_plan,omitempty" `        // 降压方案
			IsNondrugControlPlan      bool   `json:"is_nondrug_control_plan,omitempty"  `     // 是否非药物控制手段
			NondrugControlPlan        string `json:"nondrug_control_plan,omitempty"  `        // 非药物控制手段内容
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

		log.Debug("获取InfoParam", log.WithPair("IsMarried", param.IsMarried))

		// 利用reflect反射构造插入结构（数据字段太多了）
		objOfInfo := &model.PatientInfo{}
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

		// 患者信息插入
		err = uh.pinfoService.Add(context.TODO(), objOfInfo)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.PatientInfoErr, "添加患者信息失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "info add successful",
		})
	}
}

// 更新患者信息
func (uh *BaseUserHandler) UpdatePatientInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type InfoParam struct {
			RealName                  string `json:"real_name,omitempty"`                     // 真实姓名
			Sex                       string `json:"sex,omitempty"`                           // 性别
			Birthday                  string `json:"birthday,omitempty" `                     // 生日
			Tel                       string `json:"tel,omitempty" `                          // 电话号码
			IsMarried                 bool   `json:"is_married,omitempty"  `                  // 0-未婚、1-已婚
			HbpYears                  int    `json:"hbp_years,omitempty"  `                   // 高血压患病时间（年）
			Anamnesis                 string `json:"anamnesis,omitempty"  `                   // 既往病史(对应表格1~12)
			IsSmoking                 bool   `json:"is_smoking,omitempty"  `                  // 是否吸烟
			SmokingHistory            int    `json:"smoking_history,omitempty" `              // 吸烟史（年）
			SmokingDaily              int    `json:"smoking_daily,omitempty"  `               // 日吸烟数
			IsDrink                   bool   `json:"is_drink,omitempty"  `                    // 是否饮酒
			DrinkHistory              int    `json:"drink_history,omitempty"  `               // 饮酒史（年）
			DrinkDaily                int    `json:"drink_daily,omitempty"`                   // 每日饮酒量
			PatientHeight             int    `json:"patient_height,omitempty" `               // 身高
			PatientWeight             int    `json:"patient_weight,omitempty" `               // 体重
			PatientWaistCircumference int    `json:"patient_waist_circumference,omitempty" `  // 腰围
			PatientChestCircumference int    `json:"patient_chest_circumference,omitempty"  ` // 胸围
			PatientHipCircumference   int    `json:"patient_hip_circumference,omitempty"  `   // 臀围
			IsTakeChineseMedicine     bool   `json:"is_take_chinese_medicine,omitempty" `     // 是否服用中药
			AntihypertensivePlan      string `json:"antihypertensive_plan,omitempty" `        // 降压方案
			IsNondrugControlPlan      bool   `json:"is_nondrug_control_plan,omitempty"  `     // 是否非药物控制手段
			NondrugControlPlan        string `json:"nondrug_control_plan,omitempty"  `        // 非药物控制手段内容
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

		log.Debug("获取InfoParam", log.WithPair("IsMarried", param.IsMarried))

		// 利用reflect反射构造插入结构（数据字段太多了）
		objOfInfo := &model.PatientInfo{}
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

		// 患者信息插入
		err = uh.pinfoService.Update(context.TODO(), objOfInfo)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.PatientInfoErr, "添加患者信息失败"), nil)
			return
		}

		// 返回这个结果
		response.JSON(c, nil, struct {
			Result string `json:"result"`
		}{
			Result: "info update successful",
		})
	}
}

// 获取患者信息
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
