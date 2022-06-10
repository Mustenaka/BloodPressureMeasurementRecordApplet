package service

import (
	"BloodPressure/internal/model"
	"BloodPressure/internal/repo"
	"context"
)

var _ TestIndicatorCreatinineService = (*testIndicatorCreatinineService)(nil)

// TreatmentPlanService 定义用户操作服务接口
type TestIndicatorCreatinineService interface {
	// 添加一条记录
	AddById(ctx context.Context, uid uint, data int) error
	AddByIdWithTime(ctx context.Context, uid uint, data int, createAt string) error
	// 查询
	GetById(ctx context.Context, uid uint) ([]model.TestIndicatorCreatinine, error)
	GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.TestIndicatorCreatinine, error)
}

// TestIndicatorCreatinineService 实现UserService接口
type testIndicatorCreatinineService struct {
	tr repo.TestIndicatorCreatinineRepo
}

// 新血压记录服务
func NewTestIndicatorCreatinineService(_ur repo.TestIndicatorCreatinineRepo) *testIndicatorCreatinineService {
	return &testIndicatorCreatinineService{
		tr: _ur,
	}
}

// 添加
func (us *testIndicatorCreatinineService) AddById(ctx context.Context, uid uint, data int) error {
	return us.tr.AddBnp(ctx, uid, data)
}

// 添加一条带时间的肌酐记录
func (us *testIndicatorCreatinineService) AddByIdWithTime(ctx context.Context, uid uint, data int, createAt string) error {
	return us.tr.AddBnpWithTime(ctx, uid, data, createAt)
}

// 查询
func (us *testIndicatorCreatinineService) GetById(ctx context.Context, uid uint) ([]model.TestIndicatorCreatinine, error) {
	return us.tr.GetBnpById(ctx, uid)
}

// 通过limit限制查询
func (us *testIndicatorCreatinineService) GetByIdLimit(ctx context.Context, uid uint, limit int) ([]model.TestIndicatorCreatinine, error) {
	// 如果限制天数为0，则降级为获取全部数据
	if limit == 0 {
		return us.tr.GetBnpById(ctx, uid)
	}
	return us.tr.GetBnpByIdLimit(ctx, uid, limit)
}
