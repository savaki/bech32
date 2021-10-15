package bech32

import (
	"bytes"
	"encoding/base32"
	"errors"
	"io"
	"strings"

	"github.com/icza/bitio"
)

var (
	ErrInvalidCharacter = errors.New("invalid character")
	ErrInvalidFormat    = errors.New("invalid format")
	ErrInvalidLength    = errors.New("invalid length")
)

const (
	sep = "1"
)

var charset = []int{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	15, -1, 10, 17, 21, 20, 26, 30, 7, 5, -1, -1, -1, -1, -1, -1, -1, 29, -1, 24, 13, 25, 9, 8, 23,
	-1, 18, 22, 31, 27, 19, -1, 1, 0, 3, 16, 11, 28, 12, 14, 6, 4, 2, -1, -1, -1, -1, -1, -1, 29,
	-1, 24, 13, 25, 9, 8, 23, -1, 18, 22, 31, 27, 19, -1, 1, 0, 3, 16, 11, 28, 12, 14, 6, 4, 2, -1,
	-1, -1, -1, -1,
}

// Decode bech32 string into its human-readable part (hrp) and its associated data
func Decode(addr string) (hrp string, data []byte, err error) {
	if len(addr) < 8 {
		return "", nil, ErrInvalidLength
	}

	index := strings.Index(addr, sep)
	if index == -1 {
		return "", nil, ErrInvalidFormat
	}

	//s := strings.ToUpper(addr[index+1:])
	s := addr[index+1:]
	if data, err := base32.HexEncoding.DecodeString(s); err == nil {
		return addr[:index], data, nil
	}

	buf := bytes.NewBuffer(nil)
	w := bitio.NewWriter(buf)
	for _, r := range addr[index+1 : len(addr)-6] {
		i := int(r)
		if i >= len(charset) {
			return "", nil, ErrInvalidCharacter
		}

		numValue := charset[i]
		if numValue > 31 || numValue < 0 {
			return "", nil, ErrInvalidCharacter
		}

		if err := w.WriteBits(uint64(numValue), 5); err != nil {
			return "", nil, ErrInvalidCharacter
		}
	}

	r := bitio.NewReader(buf)
	for {
		v, err := r.ReadBits(8)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", nil, ErrInvalidCharacter
		}
		data = append(data, byte(v))
	}

	return addr[0:index], data, nil
}
