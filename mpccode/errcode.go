package mpccode

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func TraceId(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	traceId := spanCtx.TraceID()
	return traceId.String()

}

var (
	CodeUnknown        = (&errCode{-1, "unknown", nil}).instance_json
	CodeNil            = (&errCode{-1, "nil", nil}).instance            // No error code specified.
	CodeOK             = (&errCode{0, "ok", nil}).instance              // It is OK.
	CodeSessionInvalid = (&errCode{1, "Session Invalid", nil}).instance // The token is invalid.
	CodeParamInvalid   = (&errCode{2, "Invalid Argument", nil}).instance
	///
	CodeTokenInvalid      = (&errCode{11, "Token Invalid", nil}).instance        // The token does not exist.
	CodeTFANotExist       = (&errCode{12, "TFA NotExist", nil}).instance         // The token does not exist.
	CodeTFAExist          = (&errCode{13, "TFA Exist", nil}).instance            // The token does not exist.
	CodeTFASendSmsFailed  = (&errCode{14, "TFA Send Sms Failed", nil}).instance  // The token does not exist.
	CodeTFASendMailFailed = (&errCode{15, "TFA Send Mail Failed", nil}).instance // The token does not exist.
	CodeTFAPhoneExists    = (&errCode{16, "TFA Phone Exists", nil}).instance     // tfa绑定手机已存在
	CodeTFAMailExists     = (&errCode{17, "TFA Mail Exists", nil}).instance      // taf绑定邮箱已存在
	///
	// CodeRiskNeedVerification   = (&errCode{21, "Risk Need a VerificationCode", nil}).instance.Instance()// The risk need verification code
	CodeRiskVerifyCodeInvalid  = (&errCode{22, "Verify Code Invalid", nil}).instance // The verify code is invalid.
	CodeRiskSerialNotExist     = (&errCode{23, "Verify RiskSerial NotExist", nil}).instance
	CodeRiskVerifyPhoneInvalid = (&errCode{24, "Verify Phone Invalid", nil}).instance //
	CodeRiskVerifyMailInvalid  = (&errCode{25, "Verify Mail Invalid", nil}).instance  //
	// CodeRiskNotExist           = (&errCode{26, "Verify Risk Not Exist", nil}).instance //
	///
	CodePerformRiskForbidden        = (&errCode{31, "Perform Risk Forbidden", nil}).instance         //
	CodePerformRiskNeedVerification = (&errCode{32, "Perform Risk Need Verification", nil}).instance //
	CodePerformRiskInternalError    = (&errCode{33, "Perform Risk Error", nil}).instance             //
	CodePerformRiskTimeOut          = (&errCode{34, "Perform Risk TimeOut", nil}).instance           //
	///
	///
	CodeInternalError = (&errCode{50, "Internal Error", nil}).instance // An error occurred internally.
	//
	// ErrApiLimit           = (&errCode{100, "Limit Api", nil}
	CodeApiLimit             = (&errCode{100, "Limit Api", nil}).instance                    // 接口访问频率太高
	CodeLimitSendPhoneCode   = (&errCode{101, "Limit Api Send Phone Code", nil}).instance    // 发送手机验证码频率太高
	CodeLimitSendMailCode    = (&errCode{102, "Limit Api Send Mail Code", nil}).instance     // 发送邮箱验证码频率太高
	CodeLimitSendMailService = (&errCode{103, "Limit Send Mail Code Service", nil}).instance // 邮箱服务验证码频率太高
	////
	CodeDataNotExists = (&errCode{201, "Data Not Exists", nil}).instance
	CodeDataExists    = (&errCode{202, "Data Exists", nil}).instance
	//tx
	CodeTxsInvalid           = (&errCode{302, "Txs Invalid", nil}).instance
	CodeTxContractAbiInvalid = (&errCode{302, "Tx Contract Abi Invalid", nil}).instance

	///
	CodeErr = (&errCode{65535, "CodeErr", nil}).instance
)

//////
//////
