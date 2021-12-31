package minmax

import "testing"

func TestSingleNumber(t *testing.T) {
	if HighAndLow("1") != "1 1" {
		t.Error()
	}

	if HighAndLow("2") != "2 2" {
		t.Error()
	}
}

func TestMultipleNumbers(t *testing.T) {
	result := HighAndLow("1 2")
	if result != "2 1" {
		t.Error("Returned:", result)
	}
}
