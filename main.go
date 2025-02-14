package main

import (
	"net/http"
	"time"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/controller"
	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Clean up rate limits periodically to free up memory
	go func() {
		cleanupInterval := 5 * time.Minute

		for {
			time.Sleep(cleanupInterval)
			utils.CleanRateLimits()
		}
	}()

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")

	router.GET("/", controller.Emailhandler)
	router.GET("/otp", controller.OTPhandler)
	router.GET("/result", controller.Result)
	router.GET("next-step", controller.Nextstep)

	router.POST("/userinput", controller.UserInput)
	router.POST("/verifyotp", controller.VerifyOTP)

	router.NoRoute(func(c *gin.Context) {
		type Data struct {
			Title  string
			Result string
		}
		data := Data{
			Title:  "Result",
			Result: "Page not found",
		}

		utils.RenderHtml(c, http.StatusNotFound, "base.html", data)
	})

	router.Run(":8080")
}
