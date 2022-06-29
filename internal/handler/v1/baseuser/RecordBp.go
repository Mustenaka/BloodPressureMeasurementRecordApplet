package baseuser

import (
	"BloodPressure/internal/model"
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	strtools "BloodPressure/tools/strTools"
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 记录血压
func (uh *BaseUserHandler) RecordBp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义基本结构
		type RecordParam struct {
			RecordDateTime string `json:"record_date_time" binding:"required"`
			// RecordDate string `form:"record_date"  binding:"required"`
			// RecordTime string `form:"record_time"  binding:"required"`
			Low       int `json:"low"  binding:"required"`
			High      int `json:"high"  binding:"required"`
			HeartRate int `json:"heart_rate"  binding:"required"`
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

		// 处理日期时间为指定格式
		date, time := strtools.SplitDateTime(param.RecordDateTime)

		// 进行血压记录
		err = uh.bprSrv.AddByIdWithDateTime(context.TODO(), date, time, baseUser.UserId, param.Low, param.High, param.HeartRate)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.BPRecordErr, "添加血压记录失败"), nil)
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

// 自定义排序设计
type SortBy []model.PatientBpRecord

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	d1, _ := time.Parse(constant.MysqlDataTimeLayout, a[i].RecordDate)
	d2, _ := time.Parse(constant.MysqlDataTimeLayout, a[j].RecordDate)

	t1, _ := time.Parse(constant.TimeLayout, a[i].RecordTime)
	t2, _ := time.Parse(constant.TimeLayout, a[j].RecordTime)

	if d1.Equal(d2) {
		return t1.Before(t2)
	} else {
		return d1.Before(d2)
	}

}

// 获取血压记录
func (uh *BaseUserHandler) GetRecordBp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 字符串转换数字
		limitDays, err := strconv.Atoi(c.DefaultQuery("limit_days", "0"))
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.ValidateErr, "数据校验出错"), nil)
			return
		}

		// 通过id找到基本用户
		uid := c.GetUint(constant.UserID)
		baseUser, err := uh.userSrv.GetById(context.TODO(), uid)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
			return
		}

		// 获取血压记录
		records, err := uh.bprSrv.GetByIdLimitDay(context.TODO(), baseUser.UserId, limitDays)
		if err != nil {
			response.JSON(c, errors.Wrap(err, code.BPRecordErr, "血压记录获取失败"), nil)
			return
		}

		// 自定义排序这个血压记录
		sort.Sort(SortBy(records))

		// 返回这个结果
		response.JSON(c, nil, records)
	}
}
