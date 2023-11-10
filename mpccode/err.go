package mpccode

import (
	"errors"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type errCode struct {
	code    int
	message string
	detail  interface{}
}

func (e *errCode) Error() error {
	return errors.New(e.message)
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
func (e *errCode) SetDetail(detail interface{}) *errCode {
	e.detail = detail
	return e
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
	ErrEmpty       = errors.New("empty data")
	ErrArg         = errors.New("invalid argument")
	ErrNrpcTimeOut = errors.New("nats: timeout")
)
var (
	CodeNil            = &errCode{-1, "nil", nil}            // No error code specified.
	CodeOK             = &errCode{0, "ok", nil}              // It is OK.
	CodeSessionInvalid = &errCode{1, "Session Invalid", nil} // The token is invalid.
	CodeParamInvalid   = &errCode{2, "invalid argument", nil}
	///
	CodeTokenInvalid      = &errCode{11, "Token Invalid", nil}        // The token is invalid.
	CodeTokenNotExist     = &errCode{12, "Token NotExist", nil}       // The token does not exist.
	CodeTFANotExist       = &errCode{12, "TFA NotExist", nil}         // The token does not exist.
	CodeTFAExist          = &errCode{13, "TFA Exist", nil}            // The token does not exist.
	CodeTFASendSmsFailed  = &errCode{14, "TFA Send Sms Failed", nil}  // The token does not exist.
	CodeTFASendMailFailed = &errCode{15, "TFA Send Mail Failed", nil} // The token does not exist.
	CodeTFAPhoneExists    = &errCode{16, "TFA Phone Exists", nil}     // tfa绑定手机已存在
	CodeTFAMailExists     = &errCode{17, "TFA Mail Exists", nil}      // taf绑定邮箱已存在
	///
	// CodeRiskNeedVerification   = &errCode{21, "Risk Need a VerificationCode", nil} // The risk need verification code
	CodeRiskVerifyCodeInvalid  = &errCode{22, "Verify Code Invalid", nil} // The verify code is invalid.
	CodeRiskSerialNotExist     = &errCode{23, "Verify RiskSerial NotExist", nil}
	CodeRiskVerifyPhoneInvalid = &errCode{24, "Verify Phone Invalid", nil} //
	CodeRiskVerifyMailInvalid  = &errCode{25, "Verify Mail Invalid", nil}  //
	// CodeRiskNotExist           = &errCode{26, "Verify Risk Not Exist", nil} //
	///
	CodePerformRiskForbidden        = &errCode{31, "Perform Risk Forbidden", nil}         //
	CodePerformRiskNeedVerification = &errCode{32, "Perform Risk Need Verification", nil} //
	CodePerformRiskError            = &errCode{33, "Perform Risk Error", nil}             //
	///
	///
	CodeInternalError = &errCode{50, "Internal Error", nil} // An error occurred internally.
	//
	// ErrApiLimit           = &errCode{100, "Limit Api", nil}
	CodeApiLimit           = &errCode{100, "Limit Api", nil}                 // 接口访问频率太高
	CodeLimitSendPhoneCode = &errCode{101, "Limit Api Send Phone Code", nil} // 发送手机验证码频率太高
	CodeLimitSendMailCode  = &errCode{102, "Limit Api Send Mail Code", nil}  // 发送邮箱验证码频率太高
	////
	CodeErr = &errCode{65535, "CodeErr", nil}
)