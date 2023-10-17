package exmail

import (
	"testing"
)

func Test_Tenc_Mail(t *testing.T) {

	From := "mixmarvel-sdk@mixmarvel-sdk.com"
	SecretId := "AKID51zveEaotSAnIez267vjsxrnfR1eCZwG"
	SecretKey := "KXlv05GIC0lN2ccq1IYggZJv1CPOLKDt"
	VerificationTemplateID := uint64(26732)
	BindingCompletionTemplateID := uint64(26731)
	Subject := "MixMarver"

	m := NewTencMailClient(SecretId, SecretKey,
		VerificationTemplateID, BindingCompletionTemplateID,
		From, Subject)
	///
	stat, err := m.SendMail("xinwei.li@mixmarvel.com", "123456")
	if err != nil {
		t.Error(err)
	}

	t.Log(stat)
	////
	stat, err = m.SendBindingMail("xinwei.li@mixmarvel.com")
	if err != nil {
		t.Error(err)
	}

	t.Log(stat)
	////
}
