package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ MedicalReportEchocardiographyService = (*medicalReportEchocardiographyService)(nil)

// TreatmentPlanService 定义用户操作服务接口
type MedicalReportEchocardiographyService interface {
	// 添加一条记录
	AddById(ctx context.Context, uid uint, ef, lvidd, lvids int) error
	AddByIdWithTime(ctx context.Context, uid uint, ef, lvidd, lvids int, createAt string) error
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.MedicalReportEchocardiography, error)
	GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.MedicalReportEchocardiography, error)
}

// treatmentPlanService 实现UserService接口
type medicalReportEchocardiographyService struct {
	tr repo.MedicalReportEchocardiographyRepo
}

// 新血压记录服务
func NewMedicalReportEchocardiographyService(_ur repo.MedicalReportEchocardiographyRepo) *medicalReportEchocardiographyService {
	return &medicalReportEchocardiographyService{
		tr: _ur,
	}
}

// 添加一条记录
func (us *medicalReportEchocardiographyService) AddById(ctx context.Context, uid uint, ef, lvidd, lvids int) error {
	return us.tr.AddEchocardiography(ctx, uid, ef, lvidd, lvids)
}

// 添加一条带时间的记录
func (us *medicalReportEchocardiographyService) AddByIdWithTime(ctx context.Context, uid uint, ef, lvidd, lvids int, createAt string) error {
	return us.tr.AddEchocardiographyWithTime(ctx, uid, ef, lvidd, lvids, createAt)
}

// 查询
func (us *medicalReportEchocardiographyService) GetById(ctx context.Context, uid uint) ([]model.MedicalReportEchocardiography, error) {
	return us.tr.GetEchocardiographyById(ctx, uid)
}

// 通过limit限制查询
func (us *medicalReportEchocardiographyService) GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.MedicalReportEchocardiography, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limit == 0 {
		return us.tr.GetEchocardiographyById(ctx, uid)
	}
	return us.tr.GetEchocardiographyByIdLimit(ctx, uid, limit)
}
