package infra

import (
	"context"
	"fmt"
	"log"
	"os"

	onesignal "github.com/OneSignal/onesignal-go-api"
	"github.com/gregdel/pushover"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/config"
)

type pushNotificationSender struct {
	isOneSignalEnabled bool
	isPushOverEnabled  bool
	oneSignalClient    *onesignal.APIClient
	pushOverClient     *pushover.Pushover
}

type PushNotificationSender interface {
	PushNotification(notificationText string)
}

func NewPushNotificationSender() PushNotificationSender {
	return &pushNotificationSender{
		isOneSignalEnabled: config.PUSH_NOTIFICATION_ENV.ENABLE_ONE_SIGNAL,
		isPushOverEnabled:  config.PUSH_NOTIFICATION_ENV.ENABLE_PUSHOVER,
		oneSignalClient:    onesignal.NewAPIClient(onesignal.NewConfiguration()),
		pushOverClient:     pushover.New(config.PUSH_NOTIFICATION_ENV.PUSH_OVER_APP_TOKEN),
	}
}

func (pns *pushNotificationSender) PushNotification(notificationText string) {
	if pns.isOneSignalEnabled {
		pns.OneSignalPushNotification(notificationText)
	}
	if pns.isPushOverEnabled {
		pns.PushOverlPushNotification(notificationText)
	}
}
func (pns *pushNotificationSender) PushOverlPushNotification(notificationText string) {
	recipient := pushover.NewRecipient(config.PUSH_NOTIFICATION_ENV.PUSH_OVER_APP_RECIPIENT)
	message := pushover.NewMessage(notificationText)

	response, err := pns.pushOverClient.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	log.Println(response)
}

func (pns *pushNotificationSender) OneSignalPushNotification(notificationText string) {
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

	request := pns.oneSignalClient.DefaultApi.CreateNotification(osAuthCtx)

	_, r, err := request.Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[One Signal] Error when calling `CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "[One Signal] Full HTTP response: %v\n", r)
		return
	}

	fmt.Println("[One Signal] Push notification sent successfully!")
}
