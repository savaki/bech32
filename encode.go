// MIT License
//
// Copyright (c) 2021 Matt Ho
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package bech32

import (
	"bytes"
	"fmt"

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

func prefixCheck(prefix string) (int, error) {
	chk := 1
	for _, c := range []byte(prefix) {
		if c < 33 || c > 126 {
			return 0, fmt.Errorf("invalid prefix: %v character out of range", c)
		}
		chk = polymodStep(c>>5, chk)
	}
	chk = polymodStep(0, chk)
	for _, c := range []byte(prefix) {
		chk = polymodStep(c&0x1f, chk)
	}
	return chk, nil
}

// Encode a human readable part (hrp) and bytes as a bech32 string
func Encode(hrp string, data []byte) (encoded string, err error) {
	if len(hrp) < 1 {
		return "", ErrInvalidLength
	}

	chk, err := prefixCheck(hrp)
	if err != nil {
		return "", err
	}

	encoded = hrp + sep

	r := bitio.NewReader(bytes.NewBuffer(data))
	remainingBits := len(data) * 8
	for remainingBits > 0 {
		var b uint64
		if remainingBits > 5 {
			b, err = r.ReadBits(5)
			if err != nil {
				return "", fmt.Errorf("error reading data: %w", err)
			}
			remainingBits -= 5
		} else {
			// Consume the remaining bits, and align them as if it were a u5
			// ending in 0s
			b, err = r.ReadBits(uint8(remainingBits))
			b = b << (5 - remainingBits)
			if err != nil {
				return "", fmt.Errorf("error reading final bits: %w", err)
			}
			remainingBits = 0
		}
		chk = polymodStep(byte(b), chk)
		encoded += string(charset[b])
	}

	for i := 0; i < 6; i++ {
		chk = polymodStep(0, chk)
	}

	plm := chk ^ 1

	checksum := []byte{}
	for p := 0; p < 6; p++ {
		c := (plm >> (5 * (5 - p))) & 0x1f
		checksum = append(checksum, charset[byte(c)])
	}
	encoded += string(checksum)
	return encoded, nil
}
