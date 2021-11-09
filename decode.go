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
	"encoding/base32"
	"errors"
	"io"
	"strings"

	"github.com/icza/bitio"
)

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
		if i >= len(charset_rev) {
			return "", nil, ErrInvalidCharacter
		}

		numValue := charset_rev[i]
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
