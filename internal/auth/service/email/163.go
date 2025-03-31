package myemail

import (
	"chat/internal/auth/model"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func SendMail(op *model.EmailOptions) error {
	m := gomail.NewMessage()
	m.SetHeader("From", op.MailHost)
	m.SetHeader("To", op.MailTo)
	m.SetBody("text/plain", op.Body)
	m.SetHeader("Subject", op.Subject)
	d := gomail.NewDialer(
		op.MailHost,
		op.MailPort,
		op.MailUser,
		op.MailPass,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
