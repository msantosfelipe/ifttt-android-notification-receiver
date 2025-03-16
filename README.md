# ifttt-android-notification-receiver

> Deprecating project - New version: [notification-receiver-lambda](https://github.com/msantosfelipe/notification-receiver-lambda )
<br>
<br>

Project made to centralize notification propagation workflow

![Service draw](<draw.png>)

- Using an automation app, I receive the notifications and send it to this API
This is made using toll like IFTTT or Macrodroid

- The service filters the notifications of interest and re-send it via Notification
    - Pushover: https://github.com/gregdel/pushover
    - One Signal: Google Android (FCM) Configuration
    - Gen Firebase https://documentation.onesignal.com/docs/android-firebase-credentials
    - Configure Android SDK https://documentation.onesignal.com/docs/android-sdk-setup


- TBD: The notification are going to be stored in a database for future analysis
