package repeater

import "testing"

func TestSigleLetterASingleRepeater(t *testing.T) {
	ret := Repeater("A", 1)
	if ret != "A" {
		t.Error("Recived:", ret)
	}
}

func TestSigleWordSingleRepeater(t *testing.T) {
	ret := Repeater("Bab", 1)
	if ret != "Bab" {
		t.Error("Recived:", ret)
	}
}

func TestSigleWordMultiplyRepeater(t *testing.T) {
	ret := Repeater("CaD ", 2)
	if ret != "CaD CaD " {
		t.Error("Recived:", ret)
	}
}

func TestSigleWordZeroRepeater(t *testing.T) {
	ret := Repeater("CaD ", 0)
	if ret != "" {
		t.Error("Recived:", ret)
	}
}
