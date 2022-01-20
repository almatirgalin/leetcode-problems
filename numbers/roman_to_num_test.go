package numbers

import "testing"

func TestRomanToint(t *testing.T) {
	roman := "III"

	if romanToInt(roman) != 3 {
		t.Error("III is 3, got ", romanToInt(roman))
	}

	roman = "LVIII"

	if romanToInt(roman) != 58 {
		t.Error("LVIII is 58, got ", romanToInt(roman))
	}

	roman = "MCMXCIV"

	if romanToInt(roman) != 1994 {
		t.Error("III is 1994, got ", romanToInt(roman))
	}
}
