package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// PatientBpRecordRepo 血压记录repo接口
type TongueDetailRepo interface {
	// 添加舌苔记录
	AddTongue(ctx context.Context, id uint, tongue, tongueCoating, pulse string) error
	AddTongueWithTime(ctx context.Context, id uint, tongue, tongueCoating, pulse, createAt string) error

	// 获取记录组
	GetTongueById(ctx context.Context, id uint) ([]model.TongueDetail, error)
	GetTongueByIdLimit(ctx context.Context, id uint, limit int) ([]model.TongueDetail, error)

	// 删除
	DeleteTongueByUserId(ctx context.Context, userId uint) error
	DeleteTongueById(ctx context.Context, id uint) error
}
