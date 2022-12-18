package internal

type HashCredentials struct {
	CredId  AutoGenKey
	HSHA256 HSHA256
}

func NewHashCredentials(hsha256 HSHA256) HashCredentials {

	return HashCredentials{
		HSHA256: hsha256,
	}

}
