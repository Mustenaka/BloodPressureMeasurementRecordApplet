// Package constant 应用常量包，放置项目所需要的常量
// 建议不同的模块常量分开放，便于维护，也可以避免不同模块开发人员修改同一文件带来冲突
// 目前只有一个项目，并未拆分模块，所以集中放在这里
package constant

const (
	// RequestId 请求id名称
	RequestId = "request_id"
	// DateTimeLayout 时间格式
	DateTimeLayout      = "2006-01-02 15:04:05"
	MysqlDataTimeLayout = "2006-01-02T00:00:00+08:00"
	DateTimeLayoutMs    = "2006-01-02 15:04:05.000"
	// DateLayout 日期格式
	DateLayout = "2006-01-02"
	// TimeLayout 时间格式
	TimeLayout   = "15:04:05"
	TimeLayoutMs = "15:04:05.000"
	// UserID 用户id key
	UserID = "user_id"
)

const (
	// 用户上传的图片保存的地址
	UserUploadedPhotosPath = "./upload/images/"
)
