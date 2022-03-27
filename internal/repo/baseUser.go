package repo

import (
	"BloodPressure/internal/model"
	"context"
)

// baseUser 用户repo接口
type BaseUserRepo interface {
	GetBaseUserByName(ctx context.Context, name string) (*model.BaseUser, error)
	GetBaseUserById(ctx context.Context, id uint) (*model.BaseUser, error)
	GetBaseUserByOpenId(ctx context.Context, openid string) (*model.BaseUser, error)
	AddBaseUserByAdmin(ctx context.Context, name, password string) (*model.BaseUser, error)
	AddBaseUserByDetail(ctx context.Context, name, openid, realname, telephone, email, brithday, sex string) (*model.BaseUser, error)
}
