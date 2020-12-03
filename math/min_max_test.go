package math

import "testing"

func TestMinMax(t *testing.T) {
	x, y := 10, 20

	if MinInt(x, y) != 10 {
		t.Errorf("Invalid minimum value returned.")
	}

	if MaxInt(x, y) != 20 {
		t.Errorf("Invalid maximum value returned.")
	}

	x, y = -10, -20

	if MinInt(x, y) != -20 {
		t.Errorf("Invalid minimum value returned.")
	}

	if MaxInt(x, y) != -10 {
		t.Errorf("Invalid maximum value returned.")
	}
}
