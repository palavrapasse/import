package internal

type Password string

type Credentials struct {
	CredId   AutoGenKey
	Password Password
}

func NewPassword(password string) Password {
	return Password(password)
}

func NewCredentials(password Password) Credentials {
	return Credentials{
		Password: password,
	}
}
