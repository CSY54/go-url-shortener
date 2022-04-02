package url

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
)

func B64ToUint32(s string) (uint32, error) {
	bytes, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return 0, err
	}

	if len(bytes) > 4 {
		return 0, errors.New("Invalid ID")
	}

	if len(bytes) < 4 {
		pad := make([]byte, 4 - len(bytes))
		bytes = append(bytes, pad...)
	}

	res := binary.LittleEndian.Uint32(bytes)
	return res, nil
}

func Uint32ToB64(v uint32) (string, error) {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, v)

	for len(bs) > 1 && bs[len(bs) - 1] == 0 {
		bs = bs[:len(bs) - 1]
	}

	res := base64.RawURLEncoding.EncodeToString(bs)
	return res, nil
}

// TODO
func IsValidUrl(url string) bool {
	return true
}
