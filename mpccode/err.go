package mpccode

import (
	"encoding/json"
	"errors"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type errCode struct {
	ErrCode   int
	ErrMsg    string
	ErrDetail interface{}
}

func FromNrcpErr(err error) error {
	return &errCode{}
}
func (e *errCode) Equal(err error) bool {
	if !errors.As(err, &e) {
		return false
	}
	cerr := gerror.Cause(err)
	target := cerr.(*errCode)
	return e.ErrCode == target.ErrCode
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

func (e *errCode) instance(detail ...interface{}) error {
	errcode := &errCode{}
	if len(detail) == 0 {
		errcode = &errCode{e.ErrCode, e.ErrMsg, nil}
	} else {
		errcode = &errCode{e.ErrCode, e.ErrMsg, detail}
	}
	return gerror.NewCode(errcode, errcode.Text())
}

func (e *errCode) instance_json(val interface{}) error {
	if val == nil {
		return nil
	}
	m := &m{}
	switch val.(type) {
	case string:
		json.Unmarshal([]byte(val.(string)), m)
	case []byte:
		json.Unmarshal(val.([]byte), m)
	default:
		return nil
	}
	return gerror.NewCode(&errCode{m.Code, m.Message, m.Detail})
}
func (e *errCode) Text() string {
	j, _ := json.Marshal(e)
	return string(j)
}

func (e *errCode) Error() string {
	// return errors.New(e.message)
	return e.ErrMsg
}
func (e *errCode) Message() string {
	return e.ErrMsg
}
func (e *errCode) Code() int {
	return e.ErrCode
}

type m struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

func (e *errCode) Detail() interface{} {

	return e.ErrDetail
	// m := &m{
	// 	Code:    e.code,
	// 	Message: e.message,
	// 	Detail:  e.detail,
	// }
	// v, _ := json.Marshal(m)
	// return string(v)
}

func (e *errCode) SetDetail(detail interface{}) {

	e.ErrDetail = detail

}

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
