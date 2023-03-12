package e

import (
	"fmt"
)

type CustomError struct {
	code	ErrCode
	msg		string
}

func (e *CustomError) Error() string {
	return e.msg
}

func (e *CustomError) Code() ErrCode {
	return e.code
}

func newErrCode(code ErrCode, msg string) *CustomError {
	_, ok := codes[code]
	if ok {
		panic(fmt.Sprintf("Error[%v] already exists. Error message is %v", code, msg))
	}
	codes[code] = msg
	return &CustomError{
		code: code,
		msg: msg,
	}
}