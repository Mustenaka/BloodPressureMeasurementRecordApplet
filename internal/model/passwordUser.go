package model

// 后端管理权限表
type PasswordUser struct {
	AdminId  uint   // 管理权限id
	UserId   uint   // 用户id
	Password string // 用户密码(需要加密保存)
}
