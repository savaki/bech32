package bech32

import (
	"encoding/hex"
	"testing"
)

func TestDecode(t *testing.T) {
	testCases := map[string]struct {
		Input   string
		HRP     string
		Want    string
		WantErr bool
	}{
		"blank": {
			WantErr: true,
		},
		"type-00": {
			Input: "addr1qx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzer3n0d3vllmyqwsx5wktcd8cc3sq835lu7drv2xwl2wywfgse35a3x",
			HRP:   "addr",
			Want:  "019493315cd92eb5d8c4304e67b7e16ae36d61d34502694657811a2c8e337b62cfff6403a06a3acbc34f8c46003c69fe79a3628cefa9c47251",
		},
		"type-01": {
			Input: "addr1z8phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gten0d3vllmyqwsx5wktcd8cc3sq835lu7drv2xwl2wywfgs9yc0hh",
			HRP:   "addr",
			Want:  "11c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f337b62cfff6403a06a3acbc34f8c46003c69fe79a3628cefa9c47251",
		},
		"type-02": {
			Input: "addr1yx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzerkr0vd4msrxnuwnccdxlhdjar77j6lg0wypcc9uar5d2shs2z78ve",
			HRP:   "addr",
			Want:  "219493315cd92eb5d8c4304e67b7e16ae36d61d34502694657811a2c8ec37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f",
		},
		"type-03": {
			Input: "addr1x8phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gt7r0vd4msrxnuwnccdxlhdjar77j6lg0wypcc9uar5d2shskhj42g",
			HRP:   "addr",
			Want:  "31c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542fc37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f",
		},
		"type-04": {
			Input: "addr1gx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzer5pnz75xxcrzqf96k",
			HRP:   "addr",
			Want:  "419493315cd92eb5d8c4304e67b7e16ae36d61d34502694657811a2c8e8198bd431b03",
		},
		"type-05": {
			Input: "addr128phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gtupnz75xxcrtw79hu",
			HRP:   "addr",
			Want:  "51c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f8198bd431b03",
		},
		"type-06": {
			Input: "addr1vx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzers66hrl8",
			HRP:   "addr",
			Want:  "619493315cd92eb5d8c4304e67b7e16ae36d61d34502694657811a2c8e",
		},
		"type-07": {
			Input: "addr1w8phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gtcyjy7wx",
			HRP:   "addr",
			Want:  "71c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f",
		},
		"type-08": {
			Input: "stake1uyehkck0lajq8gr28t9uxnuvgcqrc6070x3k9r8048z8y5gh6ffgw",
			HRP:   "stake",
			Want:  "e1337b62cfff6403a06a3acbc34f8c46003c69fe79a3628cefa9c47251",
		},
		"type-09": {
			Input: "stake178phkx6acpnf78fuvxn0mkew3l0fd058hzquvz7w36x4gtcccycj5",
			HRP:   "stake",
			Want:  "f1c37b1b5dc0669f1d3c61a6fddb2e8fde96be87b881c60bce8e8d542f",
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			hrp, data, err := Decode(tc.Input)
			if tc.WantErr {
				if err == nil {
					t.Fatalf("expected not nil, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("got %v; want nil", err)
			}
			if got, want := hrp, tc.HRP; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
			if got, want := hex.EncodeToString(data), tc.Want; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

var AntiCompilerOptimization []byte

func BenchmarkDecode(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_, data, err := Decode("addr1qx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzer3n0d3vllmyqwsx5wktcd8cc3sq835lu7drv2xwl2wywfgse35a3x")
		if err != nil {
			t.Fatalf("got %v; want nil", err)
		}
		AntiCompilerOptimization = data
	}
}
