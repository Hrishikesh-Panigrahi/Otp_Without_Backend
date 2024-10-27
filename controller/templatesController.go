package controller

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/utils"
	"github.com/gin-gonic/gin"
)

func Emailhandler(c *gin.Context) {
	utils.ClearCookies(c)

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

func Result(c *gin.Context) {
	cookie, err := c.Cookie("result")

	if err != nil {
		fmt.Fprintf(c.Writer, "Cookie not found")
		return
	}

	var message string

	if cookie == "OTP is valid" {
		message = "safe"
	} else if cookie == "OTP is not valid" {
		message = "Unsafe"
	}

	type Data struct {
		Title  string
		Result string
	}

	data := Data{Title: "Result", Result: message}

	utils.RenderHtml(c, http.StatusOK, "base.html", data)
}

func Nextstep(c *gin.Context) {
	cookie, err := c.Cookie("result")

	if err != nil {
		fmt.Fprintf(c.Writer, "Cookie not found")
		return
	}

	if cookie == "OTP is valid" {
		fmt.Fprint(c.Writer, "This would ideally be Next step")
	} else {
		fmt.Fprint(c.Writer, "OTP is not valid")
	}

}
