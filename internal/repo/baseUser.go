package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// baseUser 用户repo接口
type BaseUserRepo interface {
	// 获取
	GetBaseUserByName(ctx context.Context, name string) (*model.BaseUser, error)
	GetBaseUserById(ctx context.Context, id uint) (*model.BaseUser, error)
	GetBaseUserByOpenId(ctx context.Context, openid string) (*model.BaseUser, error)
	// 添加
	AddBaseUserByNamePassword(ctx context.Context, name, password string) error
	AddBaseUserByDetail(ctx context.Context, name, openid, sex, avatarUrl string) error
	// 修改
	UpdateBaseUserDetail(ctx context.Context, srcUser *model.BaseUser, realname, telephone, email, brithday, sex, avatarUrl string) error
	UpdateBaseUserPassword(ctx context.Context, srcUser *model.BaseUser, password string) error
}
