package mysql

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"BloodPressure/pkg/db"
	timeconvert "BloodPressure/tools/timeConvert"
	"context"
)

var _ repo.TreatmentPlanRepo = (*treatmentPlanRepo)(nil)

type treatmentPlanRepo struct {
	ds db.IDataSource
}

// 创建一个新的UserRepo
func NewTreatmentPlanRepo(_ds db.IDataSource) *treatmentPlanRepo {
	return &treatmentPlanRepo{
		ds: _ds,
	}
}

// 添加计划
func (ur *treatmentPlanRepo) AddPlan(ctx context.Context, id uint, plan, note string) error {
	createDatetime := timeconvert.NowDateTimeString()
	plans := &model.TreatmentPlan{
		UserId:         id,
		Plan:           plan,
		Note:           note,
		CreateDatetime: createDatetime,
		Status:         "生效",
	}
	err := ur.ds.Master().Create(plans).Error
	return err
}

// 查询全部计划
func (ur *treatmentPlanRepo) GetPlanById(ctx context.Context, id uint) ([]model.TreatmentPlan, error) {
	plans := []model.TreatmentPlan{}
	err := ur.ds.Master().Where("user_id = ?", id).Find(&plans).Error
	return plans, err
}

// 通过limit查询限制计划
func (ur *treatmentPlanRepo) GetPlanByIdLimit(ctx context.Context, id uint, limit int) ([]model.TreatmentPlan, error) {
	plans := []model.TreatmentPlan{}
	err := ur.ds.Master().Where("user_id = ?", id).Limit(limit).Find(&plans).Error
	return plans, err
}
