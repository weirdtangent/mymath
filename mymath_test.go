package mymath

import "testing"

func TestFloatPrec(t *testing.T) {
	testFloats := [8]float64{100.0, 100.10, 0.0, 0.105, 1.105, 100.105, 100.1055, 100.1006257189}
	expected := [8]string{"100.00", "100.10", "0.00", "0.105", "1.105", "100.105", "100.1055", "100.100625"}

	for x, number := range testFloats {
		floatStr, err := FloatPrec(number, 2, 6)
		if err != nil {
			t.Errorf("FloatPrec failed for %f: expected \"%s\" but got back \"%s\"", number, expected[x], err)
		} else if floatStr != expected[x] {
			t.Errorf("FloatPrec failed for %f: expected \"%s\" but got back \"%s\"", number, expected[x], floatStr)
		}
	}
}
