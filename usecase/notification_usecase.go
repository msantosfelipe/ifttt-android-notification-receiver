package usecase

import (
	"fmt"

	"github.com/msantosfelipe/ifttt-android-notification-receiver/config"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/domain"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/infra"
)

type usecase struct {
	mailSender infra.MailSender
}

func NewNotificationUsecase(mailSender infra.MailSender) domain.NotificationUsecase {
	return &usecase{mailSender: mailSender}
}

func (uc *usecase) ProcessNotification(notification domain.Notification) error {
	if !isValidApp(notification.Name) {
		fmt.Printf("Invalid app name: %s\n", notification.Name)
		return nil
	}

	uc.SendEmail(notification.Name, notification.Body)

	return nil
}

func (uc *usecase) SendEmail(appName, notificationText string) {
	if config.EMAIL_ENV.ENABLE_EMAIL {
		uc.mailSender.SendMail(appName, notificationText)
	}

	fmt.Println("Email is disabled")
}

func isValidApp(name string) bool {
	for _, i := range config.ENV.VALID_APPS {
		if i == name {
			return true
		}
	}
	return false
}
