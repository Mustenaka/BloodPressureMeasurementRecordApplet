package strtools

// 如果src为空则返回dst，如果不是则返回src本身
func UpdateNotNullStirng(src, dst string) string {
	if src == "" {
		return dst
	}
	return src
}
