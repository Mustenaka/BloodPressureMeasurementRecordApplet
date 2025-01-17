package code

// 自定义错误码，通常错误由错误码和错误信息两部分组成，便于跟踪和维护错误信息
// 错误码为0表示成功
// 错误码3开头，表示第三方api调用错误
// 错误码4开头，表示业务层面的错误，比如校验等
// 错误码5开头，表示服务器错误，比如数组越界等
// ----------------------------------
// 错误码过多时，可以根据业务功能拆分到不同的文件或者包中

const (
	// Success 表示成功
	Success = 0
	// Unknown 无法预知或者未手动处理的错误
	Unknown = -1
)

const (
	// MpApiErr 小程序接口调用错误
	MpApiErr = iota + 30000
)

const (
	// ValidateErr 校验错误
	ValidateErr = iota + 40000

	// RequireAuthErr 没有权限
	RequireAuthErr

	// NotFoundErr 没有记录
	NotFoundErr

	// UserLoginErr 登录错误
	UserLoginErr

	// UserRegisterErr 注册错误
	UserRegisterErr

	// UserUpdateErr 修改错误
	UserUpdateErr

	// UserDeleteErr 删除错误
	UserDeleteErr

	// AuthTokenErr token 鉴权错误或权限不足
	AuthTokenErr

	// RecordCreateErr 创建记录，数据持久化失败
	RecordCreateErr

	// openidGetErr openid获取错误，可能是服务器到腾讯服务器链接出问题
	OpenidGetErr

	// BadRequestErr 错误响应码过多时
	BadRequestErr
)

const (
	// BusinessErr 业务错误
	BusinessErr = iota + 50000

	// BPRecordErr 写入血压记录错误
	BPRecordErr

	// TreatPlanErr 治疗方案记录错误
	TreatPlanErr

	// PatientInfoErr 病历信息记录错误
	PatientInfoErr

	// TongueDetailErr 舌苔脉象信息记录错误
	TongueDetailErr

	// TestIndicatorErr 检验指标错误
	TestIndicatorErr

	// MedicalRecordErr 检查报告错误
	MedicalRecordErr
)

const (
	// TransactionErr 事物提交失败
	TransactionErr = iota + 60000
	// DuplicateErr 记录存在重复
	DuplicateErr
)
