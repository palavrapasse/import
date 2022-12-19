package internal

type UserCredentials struct {
	CredId AutoGenKey
	UserId AutoGenKey
}

func NewUserCredentials(credId AutoGenKey, userId AutoGenKey) UserCredentials {
	return UserCredentials{
		CredId: credId,
		UserId: userId,
	}
}
