package entity

type HashCredentials struct {
	CredId  AutoGenKey
	HSHA256 HSHA256
}

func NewHashCredentials(credId AutoGenKey, hsha256 HSHA256) HashCredentials {
	return HashCredentials{
		CredId:  credId,
		HSHA256: hsha256,
	}

}
