package security

import (
	"fmt"
	"strings"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "E10ADC3949BA59ABBE56E057F20F883E"
	result := Md5("123456")
	fmt.Println(result)
	if str == strings.ToUpper(result) {
		t.Log(result)
	} else {
		t.Fatal("failed: ", result)
	}
}
