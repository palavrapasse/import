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

