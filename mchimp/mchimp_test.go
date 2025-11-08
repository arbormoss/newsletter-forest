package mchimp

import "testing"

func TestHeadingMaker(t *testing.T) {
	res := headingMaker(3)

	if res != "<h3>$1</h3>" {
		t.Errorf("Incorrectly make level 3 heading: %s", res)
	}

	res = headingMaker(1)

	if res != "<h1>$1</h1>" {
		t.Errorf("Incorrectly make level 1 heading: %s", res)
	}
}
