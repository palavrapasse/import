package internal

type Password string

type Credentials struct {
	CredId   AutoGenKey
	Password Password
}
