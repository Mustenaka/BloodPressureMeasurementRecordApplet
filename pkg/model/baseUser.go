package model

// 基础用户信息表
type BaseUser struct {
	UserId     uint   // 用户id
	OpenId     string // 用户的微信openid
	UserName   string // 用户名称-真实姓名
	Tel        string // 用户手机号
	Email      string // 用户邮件
	Permission int    // 用户权限
	LastTime   string // 用户的最后一次上线时间
	Sex        string // 用户性别，可选项"男"，"女"，"其他"
	Status     string // 用户状态，可选项"开启"，"关闭"
}

// 后端管理权限表
type PasswordUser struct {
	AdminId  uint   // 管理权限id
	UserId   uint   // 用户id
	Password string // 用户密码(需要加密)
}
