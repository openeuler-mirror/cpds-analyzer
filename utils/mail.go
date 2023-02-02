package utils

import (
	"github.com/sirupsen/logrus"
	gomail "gopkg.in/gomail.v2"
)

type Mail struct {
	sendCloser gomail.SendCloser
}

func NewMailSender(host string, port int, userName string, password string) (*Mail, error) {
	logrus.Debug("create new mail sender")
	logrus.Debugf("create new mail sender with host: %s", host)
	logrus.Debugf("create new mail sender with port: %s", port)
	logrus.Debugf("create new mail sender with user name: %s", userName)
	logrus.Debugf("create new mail sender with password: %s", password)
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

	logrus.Infof("send email to %s", address)
	logrus.Debugf("subject: %s", subject)
	logrus.Debugf("message: %s", message)

	return gomail.Send(m.sendCloser, msg)
}
