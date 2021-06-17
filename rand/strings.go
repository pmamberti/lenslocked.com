package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const RememberTokenBytes = 32

// Bytes help generate n random bytes or return error if
// there is one. This uses the crypto/rand package which
// makes is safe to use for remember tokens.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NBytes(base64String string) (int, error) {
	b, err := base64.URLEncoding.DecodeString(base64String)
	if err != nil {
		return -1, err
	}
	return len(b), nil
}

// String will generate a byte slice of nByte size, and then
// return a stringthat is a base64 encoded version ot that
// slice.
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken is an helper function that generates a remember
// token of RememberTokenBytes
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}
