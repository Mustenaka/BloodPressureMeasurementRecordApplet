package global

type IGloblal interface {
	GetConfigValue(selection string, key string) string
	SetConfigValue(selection string, key string, value string) bool
	DeleteConfigValue(selection string, key string)
	ReadConfigList()
}
