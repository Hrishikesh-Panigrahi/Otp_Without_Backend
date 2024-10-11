package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Hrishikesh-Panigrahi/Otp_Without_Backend/utils"
)

func UserInput(w http.ResponseWriter, r *http.Request) {
	otp := utils.GenerateOTP(6)
	ttl := 5 * 60
	expires := time.Now().Add(time.Duration(ttl) * time.Second).Format("2006-01-02 15:04:05")
	data := fmt.Sprintf(otp + "." + expires + "." + "your-secure-key")

	hash := utils.CreateHash(data)

	fmt.Printf("\n" + hash)

	fullHash := hash + "." + expires

	fmt.Printf("\n" + fullHash)

	w.Write([]byte(fullHash))
}
