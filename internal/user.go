package internal

type Email string

type User struct {
	UserId AutoGenKey
	Email  Email
}
