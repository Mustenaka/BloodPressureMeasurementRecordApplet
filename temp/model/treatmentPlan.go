package model

// 治疗计划表
type TreatmentPlan struct {
	TreatmentId uint   // 治疗方案id
	UserId      uint   // 用户id
	Plan        string // 治疗计划
	Note        string // 治疗计划备注（如禁忌症）
	CreateDate  string // 创建日期
	CreateTime  string // 创建时间
	Status      string // 状态{“生效”，“失效”}
	EndDate     string // 结束日期
	EndTime     string // 结束时间
}
