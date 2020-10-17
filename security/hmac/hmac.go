package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
)

func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{hmac: h}
}

type HMAC struct {
	hmac hash.Hash
}

func (h HMAC) Hash(input string) string {
	// reset is needed every time you want to hash the input with the same hash type
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}

func main() {
	key := "secret-key"
	password := "testtest"
	hmac := NewHMAC(key)
	hash := hmac.Hash(password)
	fmt.Println(hash)
}
