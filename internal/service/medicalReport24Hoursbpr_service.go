package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ MedicalReport24HoursbprService = (*medicalReport24HoursbprService)(nil)

// TreatmentPlanService 定义用户操作服务接口
type MedicalReport24HoursbprService interface {
	// 添加一条记录
	AddById(ctx context.Context, uid uint, report *model.MedicalReport24hoursbpr) error
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.MedicalReport24hoursbpr, error)
	GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.MedicalReport24hoursbpr, error)
}

// treatmentPlanService 实现UserService接口
type medicalReport24HoursbprService struct {
	tr repo.MedicalReport24HoursbprRepo
}

// 新血压记录服务
func NewMedicalReport24HoursbprService(_ur repo.MedicalReport24HoursbprRepo) *medicalReport24HoursbprService {
	return &medicalReport24HoursbprService{
		tr: _ur,
	}
}

// 添加一条记录
func (us *medicalReport24HoursbprService) AddById(ctx context.Context, uid uint, report *model.MedicalReport24hoursbpr) error {
	return us.tr.Addbpr(ctx, uid, report)
}

// 查询
func (us *medicalReport24HoursbprService) GetById(ctx context.Context, uid uint) ([]model.MedicalReport24hoursbpr, error) {
	return us.tr.GetbprById(ctx, uid)
}

// 通过limit限制查询
func (us *medicalReport24HoursbprService) GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.MedicalReport24hoursbpr, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limit == 0 {
		return us.tr.GetbprById(ctx, uid)
	}
	return us.tr.GetbprByIdLimit(ctx, uid, limit)
}
