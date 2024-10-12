package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/utils"
)

// global variable for hash
var fullHash string

// global variable for email
var email string

// UserInput is a controller function to handle the user input
func UserInput(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		email = r.FormValue("email")
		fmt.Print(email)

		// Generate a 6 digit OTP
		otp := utils.GenerateOTP(6)

		// Send the OTP to the user
		utils.SendOTP("User", email, otp)

		// Set the expiry time of the OTP
		ttl := 5 * 60
		expires := time.Now().Add(time.Duration(ttl) * time.Second).Format("2006-01-02 15:04:05")

		// Create a hash of the OTP, expiry time and a secure key
		data := fmt.Sprintf(email + "." + otp + "." + expires)
		hash := utils.CreateHash(data)

		// Create a full hash by concatenating the hash and the expiry time
		fullHash = hash + "." + expires

		fmt.Printf("\n" + hash)
		fmt.Printf("\n" + fullHash)
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

		// Split the full hash to get the hash and the expiry time
		Storedhash := fullHash[:64]
		expires := fullHash[65:]

		data := fmt.Sprintf(email + "." + otp + "." + expires)
		hash := utils.CreateHash(data)

		// Check if the hash is correct
		if hash == Storedhash {
			fmt.Fprintf(w, "OTP is valid")
		} else {
			fmt.Fprintf(w, "OTP is invalid")
		}

	}
}
