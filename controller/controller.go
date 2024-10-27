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

		otp, err := utils.GenerateOTP(6)
		if err != nil {
			http.Error(c.Writer, "Failed to generate OTP", http.StatusInternalServerError)
			return
		}

		utils.SendOTP("User", email, otp)

		ttl := 5 * 60
		expires := time.Now().Add(time.Duration(ttl) * time.Second)

		// Create a hash of the OTP, expiry time and a secure key
		data := fmt.Sprintf(email + "." + otp + "." + expires.Format("2006-01-02 15:04:05"))
		hash, err := utils.CreateHash(data)

		if err != nil {
			http.Error(c.Writer, "Failed to create hash", http.StatusInternalServerError)
			return
		}

		// Create a full hash by concatenating the hash and the expiry time
		fullHash := hash + "." + expires.Format("2006-01-02 15:04:05")

		fmt.Printf("\n" + hash)
		fmt.Printf("\n" + fullHash)

		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "OTP_HASH",
			Value:   fmt.Sprintf("%s.%s", hash, expires.Format("2006-01-02 15:04:05")),
			Expires: expires,
			Path:    "/",
		})

		utils.Redirect(c, "/otp", http.StatusFound)
	}

}

// VerifyOTP is a controller function to verify the OTP
// entered by the user. It compares the hash of the OTP
// with the stored hash in the cookie and stores the result
// in another cookie
func VerifyOTP(c *gin.Context) {
	clientIP := c.ClientIP()

	if !utils.CheckRateLimit(clientIP) {
		http.Error(c.Writer, "Too many attempts, please try again later", http.StatusTooManyRequests)
		return
	}

	if c.Request.Method == "POST" {
		if err := c.Request.ParseForm(); err != nil {
			http.Error(c.Writer, fmt.Sprintf("ParseForm() err: %v", err), http.StatusBadRequest)
			return
		}

		otp := c.PostForm("otp")
		fmt.Println("Received OTP:", otp)

		cookie, err := c.Cookie("OTP_HASH")
		if err != nil {
			http.Error(c.Writer, "Cookie not found", http.StatusUnauthorized)
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
		hash, err := utils.CreateHash(data)

		if err != nil {
			http.Error(c.Writer, "Failed to create hash", http.StatusInternalServerError)
			return
		}

		resultMessage := "OTP is invalid"
		if hash == storedHash {
			resultMessage = "OTP is valid"
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "result",
			Value:    resultMessage,
			HttpOnly: true,
			Secure:   true,
			Path:     "/result",
		})

		utils.Redirect(c, "/result", http.StatusFound)
	}
}
