package utils

import (
	"github.com/sirupsen/logrus"
	gomail "gopkg.in/gomail.v2"
)

type Mail struct {
	sendCloser gomail.SendCloser
}

func NewMailSender(host string, port int, userName string, password string, from string) (*Mail, error) {
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
