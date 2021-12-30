package roman

import "testing"

func TestOne(t *testing.T) {
	if Decode("I") != 1 {
		t.Error()
	}
}

func TestFive(t *testing.T) {
	if Decode("V") != 5 {
		t.Error()
	}
}
