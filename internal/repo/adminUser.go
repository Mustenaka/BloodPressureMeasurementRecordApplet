package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// baseUser 用户repo接口
type AdminUserRepo interface {
	// 获取
	GetAdminUserByName(ctx context.Context, name string) (*model.AdminUser, error)
	GetAdminUserById(ctx context.Context, id uint) (*model.AdminUser, error)
	// 添加
	AddAdminUserByNamePassword(ctx context.Context, name, password string) error
	// 修改
	UpdateAdminUserDetail(ctx context.Context, srcUser *model.AdminUser, realname, telephone, email, brithday, sex string) error
	UpdateAdminUserPassword(ctx context.Context, srcUser *model.AdminUser, password string) error
	// 冻结/解冻
	UpdateAdminUserStatus(ctx context.Context, id uint, status string) error
}
