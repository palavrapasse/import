package internal

import "crypto/sha256"

type HSHA256 string

func NewHSHA256(plainText string) (HSHA256, error) {
	var h HSHA256

	hashArray := []byte(plainText)

	sum := sha256.Sum256(hashArray)

	sumString := string(sum[:])

	h = HSHA256(sumString)

	return h, nil
}
