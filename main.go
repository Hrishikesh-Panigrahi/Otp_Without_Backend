package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	otp := GenerateOTP(6)
	ttl := 5 * 60
	expires := time.Now().Add(time.Duration(ttl) * time.Second).Format("2006-01-02 15:04:05")
	data := fmt.Sprintf(otp + "." + expires + "." + "your-secure-key")

	fmt.Printf(data)

	hash := CreateHash(data)

	fmt.Printf("\n" + hash)

	fullHash := hash + "." + expires

	fmt.Printf("\n" + fullHash)
}

func GenerateOTP(maxDigits uint32) string {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%0*d", maxDigits, bi)
}

func CreateHash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
