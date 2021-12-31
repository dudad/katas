package uniquenumber

import "testing"

func TestWhenFirstIstUniqe(t *testing.T) {
	ret := FindUniq([]float32{1, 2, 2})
	if ret != 1 {
		t.Error("Returned:", ret)
	}
}

func TestWhenLastIstUniqe(t *testing.T) {
	ret := FindUniq([]float32{2, 2, 2, 3})
	if ret != 3 {
		t.Error("Returned:", ret)
	}
}

func TestWhenSecondIstUniqe(t *testing.T) {
	ret := FindUniq([]float32{2, 4, 2, 2})
	if ret != 4 {
		t.Error("Returned:", ret)
	}
}
