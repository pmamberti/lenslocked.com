package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// NewHMAC creates and returns an HMAC object
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

// HMAC is a wrapper around crypto/hmac pkg making it easier
// to use in our code.
type HMAC struct {
	hmac hash.Hash
}

// Hash will hash the provided input using HMAC with
// the secret key provided when the HMAC type is initialised
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}
