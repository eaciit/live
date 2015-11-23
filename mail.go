package live

import (
	"gopkg.in/gomail.v2"
)

type EmailSetting struct {
	// Setting variable for email connection
	SenderEmail   string
	HostEmail     string
	PortEmail     int
	UserEmail     string
	PasswordEmail string
}

/*
Build email connection and send the email, can use for multiple receiver
*/
func (e *EmailSetting) sendEmail(subj string, msg string, sendTo []string) error {
	var err error

	if sendTo != nil {
		m := gomail.NewMessage()

		m.SetHeader("From", e.SenderEmail)
		m.SetHeader("To", sendTo...)
		//	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", subj)
		m.SetBody("text/html", msg)

		d := gomail.NewPlainDialer(e.HostEmail, e.PortEmail, e.UserEmail, e.PasswordEmail)

		err = d.DialAndSend(m)
	}

	return err
}
