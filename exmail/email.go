package exmail

import (
	"context"
	"crypto/tls"

	"github.com/mpcsdk/mpcCommon/rand"
	"gopkg.in/gomail.v2"
)

type sMailCode struct {
	From    string
	Passwd  string
	Host    string
	Port    int
	Subject string
	Body    string
	////

	d *gomail.Dialer
}

func (s *sMailCode) SendMailCode(ctx context.Context, to string) (string, error) {

	code := rand.RandomDigits(6)
	///
	m := gomail.NewMessage()
	m.SetHeader("From", s.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/html", s.Body+code)
	return code, s.d.DialAndSend(m)

}

func new() *sMailCode {

	s := &sMailCode{
		From:    "xinwei.li@mixmarvel.com",
		Passwd:  "Kkj7pJAdUpLpjgYE",
		Host:    "smtp.exmail.qq.com",
		Port:    465,
		Subject: "Subject",
		Body:    "Body",
	}
	d := gomail.NewDialer(s.Host, s.Port, s.From, s.Passwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	s.d = d
	return s
}
