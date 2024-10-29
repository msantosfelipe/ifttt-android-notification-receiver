package infra

import (
	"fmt"

	"github.com/msantosfelipe/ifttt-android-notification-receiver/config"
	gomail "gopkg.in/mail.v2"
)

type mailSender struct {
	isEnabled bool
}

type MailSender interface {
	SendMail(appName, notificationText string)
}

func NewMailSender() MailSender {
	return &mailSender{
		isEnabled: config.EMAIL_ENV.ENABLE,
	}
}

func (ms *mailSender) SendMail(appName, notificationText string) {
	if !ms.isEnabled {
		return
	}

	message := gomail.NewMessage()

	message.SetHeader("From", config.EMAIL_ENV.EMAIL_FROM)
	message.SetHeader("To", config.EMAIL_ENV.EMAIL_TO)
	message.SetHeader("Subject", fmt.Sprintf("Notification from %s", appName))

	message.SetBody("text/plain", notificationText)

	dialer := gomail.NewDialer(
		config.EMAIL_ENV.EMAIL_SERVER,
		config.EMAIL_ENV.EMAIL_PORT,
		config.EMAIL_ENV.EMAIL_USERNAME,
		config.EMAIL_ENV.EMAIL_PASSWORD,
	)

	fmt.Println("Sending email...")
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
