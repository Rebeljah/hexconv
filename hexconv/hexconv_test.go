package hexconv_test

import (
	"testing"

	"github.com/rebeljah/hexconv"
)

var hexToDecCases map[string]int

func TestMain(m *testing.M) {
	hexToDecCases = map[string]int{
		"0":      0,
		"1":      1,
		"-2":     -2,
		"3":      3,
		"4":      4,
		"5":      5,
		"6":      6,
		"7":      7,
		"8":      8,
		"9":      9,
		"A":      10,
		"B":      11,
		"C":      12,
		"D":      13,
		"E":      14,
		"-F":     -15,
		"FF":     255,
		"AB":     171,
		"F0":     240,
		"B0":     176,
		"10":     16,
		"11":     17,
		"9F":     159,
		"FFF":    4095,
		"FFFF":   65535,
		"-F0A0":  -61600,
		"-AF8B1": -719025,
	}
	m.Run()
}

func TestFromHex(t *testing.T) {
	for hex, dec := range hexToDecCases {
		d, err := hexconv.FromHex(hex)

		if err != nil {
			t.Error(err)
			continue
		}

		if d != dec {
			t.Errorf("expected: %v but got: %v", dec, d)
			continue
		}
	}
}

func TestFromDec(t *testing.T) {
	for hex, dec := range hexToDecCases {
		h, err := hexconv.FromDecimal(dec)

		if err != nil {
			t.Error(err)
			continue
		}

		if h != hex {
			t.Errorf("expected: %v but got: %v", hex, h)
			continue
		}
	}
}

func TestInterConversion(t *testing.T) {
	for hex, dec := range hexToDecCases {
		h, err := hexconv.FromDecimal(dec)

		if err != nil {
			t.Error(err)
			continue
		}

		if h != hex {
			t.Errorf("expected: %v but got: %v", hex, h)
			continue
		}

		d, err := hexconv.FromHex(h)
		if err != nil {
			t.Error(err)
			continue
		}

		if d != dec {
			t.Errorf("expected: %v but got: %v", dec, d)
			continue
		}
	}
}
