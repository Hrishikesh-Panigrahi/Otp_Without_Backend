package controller

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/utils"
	"github.com/gin-gonic/gin"
)

func Emailhandler(c *gin.Context) {
	type Data struct {
		Title string
	}

	data := Data{
		Title: "SignUp",
	}

	utils.RenderHtml(c, http.StatusOK, "base.html", data)

}

func OTPhandler(c *gin.Context) {
	type Data struct {
		Title string
	}

	data := Data{
		Title: "OTP",
	}

	utils.RenderHtml(c, http.StatusOK, "base.html", data)

}
