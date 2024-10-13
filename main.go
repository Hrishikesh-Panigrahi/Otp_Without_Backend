package main

import (
	"log"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/controller"
)

func main() {
	// Handle the routes for the Frontend
	http.HandleFunc("/", controller.Emailhandler)
	http.HandleFunc("/otp", controller.OTPhandler)

	// Handle the routes for the Backend
	http.HandleFunc("/userinput", controller.UserInput)
	http.HandleFunc("/verifyotp", controller.VerifyOTP)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
