package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ TreatmentPlanService = (*treatmentPlanService)(nil)

// TreatmentPlanService 定义用户操作服务接口
type TreatmentPlanService interface {
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.TreatmentPlan, error)
	GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.TreatmentPlan, error)
	// 添加一条记录
	AddById(ctx context.Context, uid uint, plan, note string) error
}

// treatmentPlanService 实现UserService接口
type treatmentPlanService struct {
	ur repo.TreatmentPlanRepo
}

// 新血压记录服务
func NewTreatmentPlanService(_ur repo.TreatmentPlanRepo) *treatmentPlanService {
	return &treatmentPlanService{
		ur: _ur,
	}
}

// 获取全部的测量数据
func (us *treatmentPlanService) GetById(ctx context.Context, uid uint) ([]model.TreatmentPlan, error) {
	return us.ur.GetPlanById(ctx, uid)
}

// 根据限制天数获取数据
func (us *treatmentPlanService) GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.TreatmentPlan, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limit == 0 {
		return us.ur.GetPlanById(ctx, uid)
	}
	return us.ur.GetPlanByIdLimit(ctx, uid, limit)
}

// 添加一条血压记录
func (us *treatmentPlanService) AddById(ctx context.Context, uid uint, plan, note string) error {
	return us.ur.AddPlan(ctx, uid, plan, note)
}
