package internal

import (
	"errors"
	"strings"
)

type Platform struct {
	PlatId AutoGenKey
	Name   string
}

func NewPlatform(name string) (Platform, error) {
	var p Platform

	err := checkIfPlatformNameConstraintsAreMet(name)

	if err == nil {
		p = Platform{
			Name: name,
		}
	}

	return p, err
}

func checkIfPlatformNameConstraintsAreMet(n string) error {
	size := len(strings.TrimSpace(n))

	if size == 0 {
		return errors.New("platform name can not be nil or empty")
	}

	if size > 30 {
		return errors.New("platform name constraints are not met (max 30 characters)")
	}

	return nil
}
