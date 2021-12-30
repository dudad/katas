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

func TestTwo(t *testing.T) {
	if Decode("II") != 2 {
		t.Error()
	}
}

func TestFour(t *testing.T) {
	if Decode("IV") != 4 {
		t.Error()
	}
}

func Test1666(t *testing.T) {
	if Decode("MDCLXVI") != 1666 {
		t.Error()
	}
}
