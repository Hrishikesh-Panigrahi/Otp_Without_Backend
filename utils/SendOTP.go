package utils

import (
	"fmt"
	"log"

	gomail "gopkg.in/mail.v2"
)

func SendOTP(clientName string, clientEmail string, otp string) {

	FROM_EMAIL := Config("FROM_EMAIL")
	API_HOST := Config("API_HOST")
	API_Username := Config("API_Username")
	API_Password := Config("API_Password")

	fmt.Print(FROM_EMAIL+" "+" "+API_HOST+" "+API_Username+" "+API_Password, "\n")

	// Create a new message
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", FROM_EMAIL)
	message.SetHeader("To", clientEmail)
	message.SetHeader("Subject", "Your OTP Code")

	// Set email body
	message.SetBody("text/plain", fmt.Sprintf("Hello %s, your OTP code is: %s", clientName, otp))

	// Set up the SMTP dialer
	dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, API_Username, API_Password)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("Failed to send email to %s: %v", clientEmail, err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}

}
