package repeater

import "testing"

func TestSigleARepeater(t *testing.T) {
	ret := Repeater("A", 1)
	if ret != "A" {
		t.Error("Recived:", ret)
	}
}
