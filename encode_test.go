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
	"encoding/hex"
	"testing"
)

func TestEncode(t *testing.T) {
	testCases := map[string]struct {
		HRP     string
		Data    string
		Want    string
		WantErr bool
	}{
		"blank": {
			WantErr: true,
		},
		"type-00": {
			HRP:  "addr",
			Data: "019493315cd92eb5d8c4304e67b7e16ae36d61d34502694657811a2c8e337b62cfff6403a06a3acbc34f8c46003c69fe79a3628cefa9c47251",
			Want: "addr1qx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzer3n0d3vllmyqwsx5wktcd8cc3sq835lu7drv2xwl2wywfgse35a3x",
		},
		/*"type-01": {
			HRP:  "addr",
			Data: "11c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f337b62cfff6403a06a3acbc34f8c46003c69fe79a3628cefa9c47251",
			Want: "addr1z8phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gten0d3vllmyqwsx5wktcd8cc3sq835lu7drv2xwl2wywfgs9yc0hh",
		},
		"type-02": {
			HRP:  "addr",
			Data: "219493315cd92eb5d8c4304e67b7e16ae36d61d34502694657811a2c8ec37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f",
			Want: "addr1yx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzerkr0vd4msrxnuwnccdxlhdjar77j6lg0wypcc9uar5d2shs2z78ve",
		},
		"type-03": {
			HRP:  "addr",
			Data: "31c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542fc37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f",
			Want: "addr1x8phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gt7r0vd4msrxnuwnccdxlhdjar77j6lg0wypcc9uar5d2shskhj42g",
		},
		"type-04": {
			HRP:  "addr",
			Data: "419493315cd92eb5d8c4304e67b7e16ae36d61d34502694657811a2c8e8198bd431b03",
			Want: "addr1gx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzer5pnz75xxcrzqf96k",
		},
		"type-05": {
			HRP:  "addr",
			Data: "51c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f8198bd431b03",
			Want: "addr128phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gtupnz75xxcrtw79hu",
		},
		"type-06": {
			HRP:  "addr",
			Data: "619493315cd92eb5d8c4304e67b7e16ae36d61d34502694657811a2c8e",
			Want: "addr1vx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzers66hrl8",
		},
		"type-07": {
			HRP:  "addr",
			Data: "71c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f",
			Want: "addr1w8phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gtcyjy7wx",
		},
		"type-08": {
			HRP:  "stake",
			Data: "e1337b62cfff6403a06a3acbc34f8c46003c69fe79a3628cefa9c47251",
			Want: "stake1uyehkck0lajq8gr28t9uxnuvgcqrc6070x3k9r8048z8y5gh6ffgw",
		},
		"type-09": {
			HRP:  "stake",
			Data: "f1c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f",
			Want: "stake178phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gtcccycj5",
		},*/
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			bytes, _ := hex.DecodeString(tc.Data)
			got, err := Encode(tc.HRP, bytes)
			if tc.WantErr {
				if err == nil {
					t.Fatalf("expected not nil, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("got %v; want nil", err)
			}
			if want := tc.Want; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}
