package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/utils"
)

// UserInput is a controller function to handle the user input
func UserInput(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		phone := r.FormValue("phone")
		fmt.Print(phone)

		// Generate a 6 digit OTP
		otp := utils.GenerateOTP(6)

		// Set the expiry time of the OTP
		ttl := 5 * 60
		expires := time.Now().Add(time.Duration(ttl) * time.Second).Format("2006-01-02 15:04:05")

		// Create a hash of the OTP, expiry time and a secure key
		data := fmt.Sprintf(phone + "." + otp + "." + expires)
		hash := utils.CreateHash(data)

		// Create a full hash by concatenating the hash and the expiry time
		fullHash := hash + "." + expires

		fmt.Printf("\n" + hash)
		fmt.Printf("\n" + fullHash)
	}

}
