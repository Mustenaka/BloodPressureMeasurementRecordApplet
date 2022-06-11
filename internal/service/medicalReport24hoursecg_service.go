package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ MedicalReport24hoursecgService = (*medicalReport24hoursecgService)(nil)

// TreatmentPlanService 定义用户操作服务接口
type MedicalReport24hoursecgService interface {
	// 添加一条记录
	AddById(ctx context.Context, uid uint, data string, average int) error
	AddByIdWithTime(ctx context.Context, uid uint, data string, average int, createAt string) error
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.MedicalReport24hoursecg, error)
	GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.MedicalReport24hoursecg, error)
}

// treatmentPlanService 实现UserService接口
type medicalReport24hoursecgService struct {
	tr repo.MedicalReport24hoursecgRepo
}

// 新血压记录服务
func NewMedicalReport24hoursecgService(_ur repo.MedicalReport24hoursecgRepo) *medicalReport24hoursecgService {
	return &medicalReport24hoursecgService{
		tr: _ur,
	}
}

// 添加一条记录
func (us *medicalReport24hoursecgService) AddById(ctx context.Context, uid uint, data string, average int) error {
	return us.tr.AddEcg(ctx, uid, data, average)
}

// 添加一条带时间的记录
func (us *medicalReport24hoursecgService) AddByIdWithTime(ctx context.Context, uid uint, data string, average int, createAt string) error {
	return us.tr.AddEcgWithTime(ctx, uid, data, average, createAt)
}

// 查询
func (us *medicalReport24hoursecgService) GetById(ctx context.Context, uid uint) ([]model.MedicalReport24hoursecg, error) {
	return us.tr.GetEcgById(ctx, uid)
}

// 通过limit限制查询
func (us *medicalReport24hoursecgService) GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.MedicalReport24hoursecg, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limit == 0 {
		return us.tr.GetEcgById(ctx, uid)
	}
	return us.tr.GetEcgByIdLimit(ctx, uid, limit)
}
