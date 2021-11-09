package bech32

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/icza/bitio"
)

func polymodStep(v byte, chk int) int {
	b := byte(chk >> 25)
	chk = (chk&0x01ff_ffff)<<5 ^ int(v)
	for i, g := range gen {
		if (b>>i)&1 == 1 {
			chk ^= g
		}
	}
	return chk
}

// Encode a human readable part (hrp) and bytes as a bech32 string
func Encode(hrp string, data []byte) (encoded string, err error) {
	if len(hrp) < 1 {
		return "", ErrInvalidLength
	}

	hrpBytes := []byte(hrp)
	chk := 1

	encoded = fmt.Sprintf("%s%s", hrp, sep)
	for _, v := range hrpBytes {
		chk = polymodStep(v>>5, chk)
	}
	chk = polymodStep(0, chk)
	for _, v := range hrpBytes {
		chk = polymodStep(v&0x1f, chk)
	}

	r := bitio.NewReader(bytes.NewBuffer(data))
	for {
		b, err := r.ReadBits(5)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Printf("Aborting %v\n", b)
				break
			}
			return "", ErrInvalidCharacter
		}
		chk = polymodStep(byte(b), chk)
		fmt.Printf("%v: %v (%v)\n", b, charset[b], string(charset[b]))
		encoded += string(charset[b])
	}
	fmt.Printf("%v\n", encoded)

	for i := 0; i < 6; i++ {
		chk = polymodStep(0, chk)
	}

	plm := chk ^ 1

	checksum := []byte{}
	for p := 0; p < 6; p++ {
		c := (plm >> (5 * (5 - p))) & 0x1f
		fmt.Printf("%v: %v (%v)\n", byte(c), charset[c], string(charset[c]))
		checksum = append(checksum, charset[byte(c)])
	}
	encoded += string(checksum)
	fmt.Printf("%s", encoded)
	return encoded, nil
}
