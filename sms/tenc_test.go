package sms

import (
	"testing"
)

func Test_Tenc_foreign(t *testing.T) {

	domestic := NewTencSms(
		"AKID51zveEaotSAnIez267vjsxrnfR1eCZwG",
		"KXlv05GIC0lN2ccq1IYggZJv1CPOLKDt",
		"sms.tencentcloudapi.com",
		"HMAC-SHA256",
		"ap-guangzhou",
		"1400856433",
		"",
		"1933346",
		"1941647",
	)
	resp, stat, err := domestic.SendSms("+447862429616", "456712")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}
	t.Log(resp, stat)
	////

}

func Test_Tenc_foreign_binding(t *testing.T) {

	domestic := NewTencSms(
		"AKID51zveEaotSAnIez267vjsxrnfR1eCZwG",
		"KXlv05GIC0lN2ccq1IYggZJv1CPOLKDt",
		"sms.tencentcloudapi.com",
		"HMAC-SHA256",
		"ap-guangzhou",
		"1400856433",
		"",
		"1933346",
		"1941647",
	)

	///
	resp, stat, err := domestic.SendBinding("+447862429616")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}
	t.Log(resp, stat)
	///

}

func Test_Tenc_domestic_incorrect(t *testing.T) {
	domestic := NewTencSms(
		"AKID51zveEaotSAnIez267vjsxrnfR1eCZwG",
		"KXlv05GIC0lN2ccq1IYggZJv1CPOLKDt",
		"sms.tencentcloudapi.com",
		"HMAC-SHA256",
		"ap-guangzhou",
		"1400856433",
		"",
		"1933346",
		"1941647",
	)
	resp, stat, err := domestic.SendSms("+4478624296161", "4567")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}

	t.Log(resp, stat)
	///

}
func Test_Tenc_domestic_xinjiapo(t *testing.T) {
	domestic := NewTencSms(
		"AKID51zveEaotSAnIez267vjsxrnfR1eCZwG",
		"KXlv05GIC0lN2ccq1IYggZJv1CPOLKDt",
		"sms.tencentcloudapi.com",
		"HMAC-SHA256",
		"ap-guangzhou",
		"1400856433",
		"",
		"1933346",
		"1941647",
	)

	resp, stat, err := domestic.SendSms("+659035559", "4567")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}

	t.Log(resp, stat)
}

func Test_Tenc_foreign_xinjiapo2(t *testing.T) {
	domestic := NewTencSms(
		"AKID51zveEaotSAnIez267vjsxrnfR1eCZwG",
		"KXlv05GIC0lN2ccq1IYggZJv1CPOLKDt",
		"sms.tencentcloudapi.com",
		"HMAC-SHA256",
		"ap-guangzhou",
		"1400856433",
		"",
		"1933346",
		"1941647",
	)
	resp, stat, err := domestic.SendSms("+6588606326", "4567")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}

	t.Log(resp, stat)
}
