package e

import (
	"fmt"
)
type ErrInterface interface {
	error
	Code()	uint32
}

type CustomError struct {
	code	uint32
	msg		string
}

func (e *CustomError) Error() string {
	return e.msg
}

func (e *CustomError) Code() uint32 {
	return e.code
}

func newErrCode(code uint32, msg string) ErrInterface {
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