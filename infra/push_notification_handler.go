package infra

import (
	"context"
	"fmt"
	"os"

	onesignal "github.com/OneSignal/onesignal-go-api"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/config"
)

type pushNotificationSender struct {
	isEnabled bool
	apiClient *onesignal.APIClient
}

type PushNotificationSender interface {
	PushNotification(notificationText string)
}

func NewPushNotificationSender() PushNotificationSender {
	return &pushNotificationSender{
		isEnabled: config.PUSH_NOTIFICATION_ENV.ENABLE,
		apiClient: onesignal.NewAPIClient(onesignal.NewConfiguration()),
	}
}

func (pns *pushNotificationSender) PushNotification(notificationText string) {
	if !pns.isEnabled {
		return
	}

	appId := config.PUSH_NOTIFICATION_ENV.ONE_SIGNAL_APP_ID
	restApiKey := config.PUSH_NOTIFICATION_ENV.ONE_SIGNAL_REST_API_KEY

	osAuthCtx := context.WithValue(
		context.Background(),
		onesignal.AppAuth,
		restApiKey,
	)

	notification := *onesignal.NewNotification(appId)
	notification.SetIncludedSegments([]string{config.PUSH_NOTIFICATION_ENV.ONE_SIGNAL_SEGMENT})
	notification.SetIsIos(false)
	stringMap := onesignal.StringMap{En: &notificationText}
	notification.Contents = *onesignal.NewNullableStringMap(&stringMap)

	request := pns.apiClient.DefaultApi.CreateNotification(osAuthCtx)

	_, r, err := request.Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	fmt.Println("Push notification sent successfully!")
}
