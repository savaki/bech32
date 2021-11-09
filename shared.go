package bech32

import "errors"

var (
	ErrInvalidCharacter = errors.New("invalid character")
	ErrInvalidFormat    = errors.New("invalid format")
	ErrInvalidLength    = errors.New("invalid length")
	ErrInvalidData      = errors.New("invalid data")
)

const (
	sep = "1"
)

var charset_rev = []int{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	15, -1, 10, 17, 21, 20, 26, 30, 7, 5, -1, -1, -1, -1, -1, -1, -1, 29, -1, 24, 13, 25, 9, 8, 23,
	-1, 18, 22, 31, 27, 19, -1, 1, 0, 3, 16, 11, 28, 12, 14, 6, 4, 2, -1, -1, -1, -1, -1, -1, 29,
	-1, 24, 13, 25, 9, 8, 23, -1, 18, 22, 31, 27, 19, -1, 1, 0, 3, 16, 11, 28, 12, 14, 6, 4, 2, -1,
	-1, -1, -1, -1,
}

var charset = []byte{
	'q', 'p', 'z', 'r', 'y', '9', 'x', '8',
	'g', 'f', '2', 't', 'v', 'd', 'w', '0',
	's', '3', 'j', 'n', '5', '4', 'k', 'h',
	'c', 'e', '6', 'm', 'u', 'a', '7', 'l',
}

var gen = []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}
