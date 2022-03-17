package errors

import (
	code2 "BloodPressure/pkg/errors/code"
	"errors"
	"testing"
)

func TestDecodeErr(t *testing.T) {
	err := WithCode(code2.Success, "success")
	code, msg := DecodeErr(err)
	if code != code2.Success {
		t.Error("unexpected errcode")
	}
	if msg != "success" {
		t.Error("unexpected msg")
	}
	t.Logf("errcode = %d, message=%s \r\n", code, msg)
}

func TestWrap(t *testing.T) {
	err := errors.New("top1")
	err2 := WithCode(code2.Unknown, "unknown")

	err3 := Wrap(err, code2.Unknown, "unknow")
	if !errors.Is(err3, err) {
		t.Error("expected value is err")
	}
	err4 := Wrap(err2, code2.ValidateErr, "validate err")

	var err5 *bizErrWithCode
	if !errors.As(err4, &err5) {
		t.Error("expected type *bizErrWithCode")
	}
	t.Logf("%s\r\n", err5)
}
