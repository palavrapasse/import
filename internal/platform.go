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
	p := Platform{}

	err := p.SetName(name)

	return p, err
}

func checkIfPlatformNameConstraintsAreMet(n string) error {
	size := len(strings.TrimSpace(n))

	if size == 0 {
		return errors.New("platform name can not be nil or empty")
	}

	if size > 30 {
		return errors.New("platform name exceeds 30 characters")
	}

	p.Name = n

	return nil
}
