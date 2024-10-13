package main

import (
	"log"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/controller"
)

func main() {

	http.HandleFunc("/", controller.Emailhandler)
	http.HandleFunc("/otp", controller.Emailhandler)

	http.HandleFunc("/userinput", controller.UserInput)
	http.HandleFunc("/verifyotp", controller.VerifyOTP)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
