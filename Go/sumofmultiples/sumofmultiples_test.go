package sumofmultiples

import "testing"

func TestValuesLessThan4(t *testing.T) {
	ret := Multiple3And5(3)
	if ret != 0 {
		t.Error("Get:", ret)
	}
}


func Test4(t *testing.T) {
	ret := Multiple3And5(4)
	if ret != 3 {
		t.Error("Get:", ret)
	}
}
