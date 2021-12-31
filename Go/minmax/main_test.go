package minmax

import "testing"

func TestSingleNumber(t *testing.T) {
	if HighAndLow("1") != "1 1" {
		t.Error()
	}
}
