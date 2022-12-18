package internal

import (
	"errors"
	"strings"
)

type HSHA256 string

func NewHSHA256(hash string) (HSHA256, error) {
	var h HSHA256

	err := checkIfConstraintsAreMet(hash)

	if err == nil {
		h = HSHA256(hash)
	}

	return h, err

}

func checkIfConstraintsAreMet(hash string) error {
	size := len(strings.TrimSpace(hash))

	if size == 0 {
		return errors.New("HSHA256 can not be empty")
	}

	if size > 64 {
		return errors.New("HSHA256 constraints are not met (max 64 characters)")
	}

	return nil
}
