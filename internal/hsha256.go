package internal

type HSHA256 string

func NewHSHA256(hash string) (HSHA256, error) {
	var h HSHA256

	h = HSHA256(hash)

	return h, nil

}
