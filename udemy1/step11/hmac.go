package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const apiKey = "User1Key"
const apiSecret = "User1Secret"

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	exepectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == exepectedHMAC)
}

func main() {
	h := hmac.New(sha256.New, []byte(apiSecret))
	fmt.Println(h)
	data := []byte("data-blabla")

	h.Write((data))
	fmt.Println(h)

	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)

	Server(apiKey, sign, data)
}
