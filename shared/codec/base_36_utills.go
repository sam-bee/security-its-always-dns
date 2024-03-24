package codec

import "math/big"

func StringToBase36(data []byte) string {
	bigint := big.NewInt(0)
	bigint.SetBytes(data)
	base36 := bigint.Text(36)
	return base36
}

func Base36ToString(data string) (string, error) {
	bigint := big.NewInt(0)
	bigint.SetString(data, 36)
	decodedBytes := bigint.Bytes()
	return string(decodedBytes), nil
}
