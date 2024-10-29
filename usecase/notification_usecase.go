package usecase

import (
	"fmt"

	"github.com/msantosfelipe/ifttt-android-notification-receiver/config"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/domain"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/infra"
)

type usecase struct {
	mailSender             infra.MailSender
	pushNotificationSender infra.PushNotificationSender
}

func NewNotificationUsecase(
	mailSender infra.MailSender,
	pushNotificationSender infra.PushNotificationSender,
) domain.NotificationUsecase {
	return &usecase{
		mailSender:             mailSender,
		pushNotificationSender: pushNotificationSender,
	}
}

func (uc *usecase) ProcessNotification(notification domain.Notification) error {
	if !isValidApp(notification.Name) {
		fmt.Printf("Invalid app name: %s\n", notification.Name)
		return nil
	}

	uc.mailSender.SendMail(notification.Name, notification.Body)

	uc.pushNotificationSender.PushNotification(
		fmt.Sprintf("%s: - %s", notification.Name, notification.Body),
	)

	return nil
}

func isValidApp(name string) bool {
	if config.ENV.ALLOW_ALL_APPS {
		return true
	}

	for _, i := range config.ENV.ALLOWED_APPS {
		if i == name {
			return true
		}
	}
	return false
}
