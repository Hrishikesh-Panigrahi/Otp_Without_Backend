package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/utils"
	"github.com/gin-gonic/gin"
)

// global variable for email
var email string

// UserInput is a controller function to handle the user input
// and send the OTP to the user
func UserInput(c *gin.Context) {

	if c.Request.Method == "POST" {
		if err := c.Request.ParseForm(); err != nil {
			fmt.Fprintf(c.Writer, "ParseForm() err: %v", err)
			return
		}

		email = c.PostForm("email")
		fmt.Println("Email received:", email)

		otp := utils.GenerateOTP(6)

		utils.SendOTP("User", email, otp)

		ttl := 5 * 60
		expires := time.Now().Add(time.Duration(ttl) * time.Second)

		// Create a hash of the OTP, expiry time and a secure key
		data := fmt.Sprintf(email + "." + otp + "." + expires.Format("2006-01-02 15:04:05"))
		hash := utils.CreateHash(data)

		// Create a full hash by concatenating the hash and the expiry time
		fullHash := hash + "." + expires.Format("2006-01-02 15:04:05")

		fmt.Printf("\n" + hash)
		fmt.Printf("\n" + fullHash)

		//set cookies
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "OTP_HASH",
			Value:   fmt.Sprintf("%s.%s", hash, expires.Format("2006-01-02 15:04:05")),
			Expires: expires,
			Path:    "/",
		})

		// Redirect to the OTP page
		utils.Redirect(c, "/otp", http.StatusFound)
	}

}

// VerifyOTP is a controller function to verify the OTP
// entered by the user. It compares the hash of the OTP
// with the stored hash in the cookie and stores the result
// in another cookie
func VerifyOTP(c *gin.Context) {
	if c.Request.Method == "POST" {
		if err := c.Request.ParseForm(); err != nil {
			fmt.Fprintf(c.Writer, "ParseForm() err: %v", err)
			return
		}

		otp := c.PostForm("otp")
		fmt.Print(otp)
		cookie, err := c.Cookie("OTP_HASH")

		if err != nil {
			fmt.Fprintf(c.Writer, "Cookie not found")
			return
		}

		parts := strings.Split(cookie, ".")
		if len(parts) < 2 {
			http.Error(c.Writer, "Invalid OTP data", http.StatusUnauthorized)
			return
		}

		storedHash := parts[0]
		expiryStr := parts[1]

		// Parse the expiry timestamp
		expiry, err := time.Parse("2006-01-02 15:04:05", expiryStr)
		if err != nil || time.Now().After(expiry) {
			http.Error(c.Writer, "OTP expired", http.StatusUnauthorized)
			return
		}

		data := fmt.Sprintf(email + "." + otp + "." + expiryStr)
		hash := utils.CreateHash(data)

		if hash == storedHash {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:  "result",
				Value: "OTP is valid",
			})
		} else {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:  "result",
				Value: "OTP is invalid",
			})
		}
		
		utils.Redirect(c, "/result", http.StatusFound)
	}
}
