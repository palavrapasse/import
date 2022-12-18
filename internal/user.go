package internal

import (
	"errors"
	"net/mail"
	"strings"
)

type Email string

type User struct {
	UserId AutoGenKey
	Email  Email
}

func NewUser(email string) (User, error) {
	u := User{}

	err := u.SetEmail(email)

	return u, err
}

func (u *User) SetEmail(e string) error {
	size := len(strings.TrimSpace(e))

	if size == 0 {
		return errors.New("user email can not be nil or empty")
	}

	if size > 130 {
		return errors.New("user email exceeds 130 characters")
	}

	_, err := mail.ParseAddress(e)
	if err != nil {
		return err
	}

	u.Email = Email(e)

	return nil
}
