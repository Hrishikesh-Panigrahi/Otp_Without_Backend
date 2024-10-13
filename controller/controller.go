package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/utils"
)

// global variable for email
var email string

// UserInput is a controller function to handle the user input
// and send the OTP to the user
func UserInput(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		email = r.FormValue("email")
		fmt.Println("Email received:", email)

		// Generate a 6 digit OTP
		otp := utils.GenerateOTP(6)

		// Send the OTP to the user
		utils.SendOTP("User", email, otp)

		// Set the expiry time of the OTP
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
		http.SetCookie(w, &http.Cookie{
			Name:    "OTP_HASH",
			Value:   fmt.Sprintf("%s.%s", hash, expires.Format("2006-01-02 15:04:05")),
			Expires: expires,
			Path:    "/",
		})

		// Redirect to the OTP page
		http.Redirect(w, r, "/otp", http.StatusFound)
	}

}

// VerifyOTP is a controller function to verify the OTP
func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		otp := r.FormValue("otp")
		fmt.Print(otp)
		cookie, err := r.Cookie("OTP_HASH")

		if err != nil {
			fmt.Fprintf(w, "Cookie not found")
			return
		}

		parts := strings.Split(cookie.Value, ".")
		if len(parts) < 2 {
			http.Error(w, "Invalid OTP data", http.StatusUnauthorized)
			return
		}

		storedHash := parts[0]
		expiryStr := parts[1]

		// Parse the expiry timestamp
		expiry, err := time.Parse("2006-01-02 15:04:05", expiryStr)
		if err != nil || time.Now().After(expiry) {
			http.Error(w, "OTP expired", http.StatusUnauthorized)
			return
		}

		data := fmt.Sprintf(email + "." + otp + "." + expiryStr)
		hash := utils.CreateHash(data)

		// Check if the hash is correct
		if hash == storedHash {
			fmt.Fprintf(w, "OTP is valid")
		} else {
			fmt.Fprintf(w, "OTP is invalid")
		}
	}
}
