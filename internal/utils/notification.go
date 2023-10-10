package utils

import "fmt"

type NotificationType string

const (
	NotificationTypeEmail NotificationType = "email"
	NotificationTypeSms   NotificationType = "sms"
)

type Notification struct {
	Type    NotificationType
	Subject string
	Body    string
	Target  string
}

func SendNotification(notification Notification) error {
	if notification.Type == NotificationTypeEmail {
		// Set up the email sender
		email := "youremail@gmail.com"
		password := "yourpassword"

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", email)
		mailer.SetHeader("To", "recipient@example.com")
		mailer.SetHeader("Subject", "Hello from Golang")
		mailer.SetBody("text/plain", "This is the email body.")

		// Create a new SMTP client session
		dialer := gomail.NewDialer("smtp.gmail.com", 587, email, password)

		// Send the email
		if err := dialer.DialAndSend(mailer); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Email sent successfully!")

	} else if notification.Type == NotificationTypeSms {
		// do action if have env for send sms
	}
	return nil
}
