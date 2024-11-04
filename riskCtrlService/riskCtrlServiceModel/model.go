package riskctrlservicemodel

type SendPhoneCodeReq struct {
	UserId     string
	RiskSerial string
	Phone      string
}
type SendMailCodeReq struct {
	UserId     string
	RiskSerial string
	Mail       string
}
type VerifyCodeReq struct {
	UserId     string
	RiskSerial string
	PhoneCode  string
	MailCode   string
}
type RiskTxsReq struct {
	UserId   string
	SignData string
}
type RiskTxsRes struct {
	Ok         int32
	RiskKind   []string
	RiskSerial string
	Msg        string
}
