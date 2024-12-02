package domain

type Config struct {
	API_PREFIX     string
	PORT           string
	VALID_API_KEY  string
	ALLOW_ALL_APPS bool
	ALLOWED_APPS   []string
	ALLOWED_TITLES []string
}

type Email struct {
	ENABLE         bool
	EMAIL_SERVER   string
	EMAIL_PORT     int
	EMAIL_FROM     string
	EMAIL_TO       string
	EMAIL_USERNAME string
	EMAIL_PASSWORD string
}

type PushNotification struct {
	ENABLE_ONE_SIGNAL       bool
	ENABLE_PUSHOVER         bool
	ONE_SIGNAL_APP_ID       string
	ONE_SIGNAL_REST_API_KEY string
	ONE_SIGNAL_SEGMENT      string
	PUSH_OVER_APP_TOKEN     string
	PUSH_OVER_APP_RECIPIENT string
}
