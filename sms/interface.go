package sms

type ISmsSender interface {
	SendSms(string, string, string) (bool, string, error)
}
