package reg

import (
	"regexp"
)

// Date verify
func VerifyDateFormat(email string) bool {
	// YYYY-MM-DD
	pattern := `(([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))-02-29)$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// DateTime verify
func VerifyDateTimeFormat(email string) bool {
	// YYYY-MM-DD HH:MM:SS
	pattern := `^((\d{2}(([02468][048])|([13579][26]))[\-]?((((0[13578])|(1[02]))[\-]?((0[1-9])|([1-2][0-9])|(3[01])))|(((0[469])|(11))[\-]?((0[1-9])|([1-2][0-9])|(30)))|(02[\-]?((0[1-9])|([1-2][0-9])))))|(\d{2}(([02468][1235679])|([13579][01345789]))[\-]?((((0[13578])|(1[02]))[\-]?((0[1-9])|([1-2][0-9])|(3[01])))|(((0[469])|(11))[\-]?((0[1-9])|([1-2][0-9])|(30)))|(02[\-]?((0[1-9])|(1[0-9])|(2[0-8])))))) (((([0-1][0-9])|(2[0-3]))[\:]?([0-5][0-9])[\:]?((([0-5][0-9])))))$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// Time verify
func VerifyTimeFormat(email string) bool {
	// HH:MM:SS
	pattern := `([0-1][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// email verify
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
