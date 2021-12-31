package repeater

import "testing"

func TestSigleARepeater(t *testing.T) {
	ret := Repeater("A", 1)
	if ret != "A" {
		t.Error("Recived:", ret)
	}
}

func TestSigleWordRepeater(t *testing.T) {
	ret := Repeater("Bab", 1)
	if ret != "Bab" {
		t.Error("Recived:", ret)
	}
}
