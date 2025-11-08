package mchimp

import "testing"

func TestHeadingMaker(t *testing.T) {
	res := headingMaker(3)

	if res != "\n<h3>$1</h3>\n" {
		t.Errorf("Incorrectly make level 3 heading: %s", res)
	}

	res = headingMaker(1)

	if res != "\n<h1>$1</h1>\n" {
		t.Errorf("Incorrectly make level 1 heading: %s", res)
	}
}
