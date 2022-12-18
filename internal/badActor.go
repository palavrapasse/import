package internal

import (
	"errors"
	"strings"
)

type BadActor struct {
	BaId       AutoGenKey
	Identifier string
}

func NewBadActor(identifier string) (BadActor, error) {
	b := BadActor{}
	err := b.SetIdentifier(identifier)

	return b, err
}

func (b *BadActor) SetIdentifier(identifier string) error {
	size := len(strings.TrimSpace(identifier))

	if size == 0 {
		return errors.New("badActor identifier can not be nil or empty")
	}

	if size > 30 {
		return errors.New("badActor identifier exceeds 30 characters")
	}

	b.Identifier = identifier

	return nil
}
