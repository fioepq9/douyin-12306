package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode int32 = iota
	ServiceErrCode
	ParamErrCode
	LoginErrCode
	UserAlreayExistErrCode
	SonyflakeGenerateIdFailErrCode
	UserNotExistErrCode
	UserNotLoginErrCode
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{ErrCode: code, ErrMsg: msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                    = NewErrNo(SuccessCode, "Success")
	ServiceErr                 = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr                   = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	LoginErr                   = NewErrNo(LoginErrCode, "Wrong username or password")
	UserAlreayExistErr         = NewErrNo(UserAlreayExistErrCode, "User already exist")
	SonyflakeGenerateIdFailErr = NewErrNo(SonyflakeGenerateIdFailErrCode, "Sonyflake generate id fail")
	UserNotExistErr            = NewErrNo(UserNotExistErrCode, "User doesn't exist")
	UserNotLoginErr            = NewErrNo(UserNotLoginErrCode, "User not login")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
