package sms

import (
	"testing"
)

func Test_domestic(t *testing.T) {

	domestic := &Huawei{
		APIAddress:        "https://smsapi.cn-south-1.myhuaweicloud.com:443/sms/batchSendSms/v1",
		ApplicationKey:    "DZcZWoIauKdHY1wD0179m4Jk5N9V",
		ApplicationSecret: "1GSPX8WyCwJbTfwEohqsBlOzbcam",
		Sender:            "8823091933902",
		TemplateID:        "000ab97205d34659a78b47c9e1a805fb",
		Signature:         "幂玛",
	}
	resp, stat, err := domestic.SendSms("+8615111226175", "4567")
	if err != nil {
		t.Error(err)
	}
	if stat != "" {
		t.Log(stat)
		t.Error(err)
	}
	t.Log(resp)
}
