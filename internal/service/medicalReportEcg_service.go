package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ MedicalReportEcgService = (*medicalReportEcgService)(nil)

// TreatmentPlanService 定义用户操作服务接口
type MedicalReportEcgService interface {
	// 添加一条记录
	AddById(ctx context.Context, uid uint, data string) error
	AddByIdWithTime(ctx context.Context, uid uint, data string, createAt string) error
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.MedicalReportEcg, error)
	GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.MedicalReportEcg, error)
}

// medicalReportEcgService 实现UserService接口
type medicalReportEcgService struct {
	tr repo.MedicalReportEcgRepo
}

// 新血压记录服务
func NewMedicalReportEcgService(_ur repo.MedicalReportEcgRepo) *medicalReportEcgService {
	return &medicalReportEcgService{
		tr: _ur,
	}
}

// 添加一条记录
func (us *medicalReportEcgService) AddById(ctx context.Context, uid uint, data string) error {
	return us.tr.AddEcg(ctx, uid, data)
}

// 添加一条带时间的记录
func (us *medicalReportEcgService) AddByIdWithTime(ctx context.Context, uid uint, data string, createAt string) error {
	return us.tr.AddEcgWithTime(ctx, uid, data, createAt)
}

// 查询
func (us *medicalReportEcgService) GetById(ctx context.Context, uid uint) ([]model.MedicalReportEcg, error) {
	return us.tr.GetEcgById(ctx, uid)
}

// 通过limit限制查询
func (us *medicalReportEcgService) GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.MedicalReportEcg, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limit == 0 {
		return us.tr.GetEcgById(ctx, uid)
	}
	return us.tr.GetEcgByIdLimit(ctx, uid, limit)
}
