package entity

type HashUser struct {
	UserId  AutoGenKey
	HSHA256 HSHA256
}

func NewHashUser(userId AutoGenKey, hsha256 HSHA256) HashUser {
	return HashUser{
		UserId:  userId,
		HSHA256: hsha256,
	}

}
