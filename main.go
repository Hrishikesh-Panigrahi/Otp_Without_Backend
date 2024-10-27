package main

import (
	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*.html")

	router.GET("/", controller.Emailhandler)
	router.GET("/otp", controller.OTPhandler)

	// Handle the routes for the Backend
	router.GET("/userinput", controller.UserInput)
	router.GET("/verifyotp", controller.VerifyOTP)

	router.Run(":8080")
}
