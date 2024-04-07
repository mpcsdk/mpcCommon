package mpccode

import (
	"errors"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type errCode struct {
	code    int
	message string
	detail  interface{}
}

func (e *errCode) Equal(err error) bool {
	if !errors.As(err, &e) {
		return false
	}
	cerr := gerror.Cause(err)
	target := cerr.(*errCode)
	return e.code == target.code
}

func Equal(err error, target error) bool {
	c := gerror.Code(err)
	if c == gcode.CodeNil {
		return false
	}
	tc := gerror.Code(target)
	if tc == gcode.CodeNil {
		return false
	}
	///
	if tc.Code() == c.Code() {
		return true
	}
	return false
}

//
//	func (e *errCode) errCode() gcode.Code {
//		return gcode.New(e.code, e.message, e.detail)
//	}
// func (e *errCode) CodeErr() error {
// 	return gerror.WrapCode(gcode.New(e.code, e.message, e.detail), e)
// }

func (e *errCode) instance(detail ...interface{}) error {
	if len(detail) == 0 {
		return gerror.NewCode(&errCode{e.code, e.message, nil})
	} else {
		return gerror.NewCode(&errCode{e.code, e.message, detail})
	}
}
func (e *errCode) Error() string {
	// return errors.New(e.message)
	return e.message
}
func (e *errCode) Message() string {
	return e.message
}
func (e *errCode) Code() int {
	return e.code
}
func (e *errCode) Detail() interface{} {
	return e.detail
}

// func (e *errCode) SetDetail(detail any) error {
// 	e.detail = detail
// 	return e.CodeErr()
// }

type errDetail struct {
	K string
	V interface{}
}

func ErrDetail(k string, v interface{}) *errDetail {
	return &errDetail{k, v}
}

func errData(data map[string]interface{}, kvs ...*errDetail) {
	if len(kvs) >= 1 {
		d := kvs[0]
		data[d.K] = d.V
		errData(data, kvs[1:]...)
	}
}
func ErrDetails(kvs ...*errDetail) string {
	data := map[string]interface{}{}
	errData(data, kvs...)
	str, err := gjson.EncodeString(data)
	if err != nil {
		return err.Error()
	}
	return str
}

var (
	ErrNrpcTimeOut = errors.New("nats: timeout")
)
