package utils

import (
	"github.com/sirupsen/logrus"
	gomail "gopkg.in/gomail.v2"
)

type Mail struct {
	sendCloser gomail.SendCloser
}

func NewMailSender(host string, port int, userName string, password string) (*Mail, error) {
	d := gomail.NewDialer(host, port, userName, password)
	s, err := d.Dial()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &Mail{
		sendCloser: s,
	}, nil
}

func (m *Mail) Send(name, address, subject, message string) error {
	msg := gomail.NewMessage()

	msg.SetHeader("From", "no-reply@linx-info.com")
	msg.SetAddressHeader("To", address, name)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", message)

	return gomail.Send(m.sendCloser, msg)
}
