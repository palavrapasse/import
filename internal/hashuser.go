package internal

type HashUser struct {
	UserId  AutoGenKey
	HSHA256 HSHA256
}

func NewHashUser(hsha256 HSHA256) HashUser {

	return HashUser{
		HSHA256: hsha256,
	}

}
