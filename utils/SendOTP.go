package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

func SendOTP(clientName string, clientEmail string, otp string) {
	ms := mailersend.NewMailersend(Config("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := "Subject"
	text := "your OTP is" + otp
	html := "Greetings from the team, you got this message through MailerSend."

	from := mailersend.From{
		Name:  Config("MAILERSEND_FROM_NAME"),
		Email: Config("MAILERSEND_FROM_EMAIL"),
	}

	recipients := []mailersend.Recipient{
		{
			Name:  clientName,
			Email: clientEmail,
		},
	}

	tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetTags(tags)

	message.SetInReplyTo("client-id")

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return
	}
	fmt.Printf("Email sent! Message ID: %s\n", res.Header.Get("X-Message-Id"))
}
