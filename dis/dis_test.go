package dis

import (
	"strings"
	"testing"
)

func TestParseImagesWithNoImages(t *testing.T) {
	md := `Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.
Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.`

	res := parseImages(md)

	if len(res) != 0 {
		t.Errorf("Parsed nonexistent images from md: %s", strings.Join(res, ", "))
	}
}

func TestParseImagesWithTouchingImages(t *testing.T) {
	md := `![img1](img1link)![img2](img2link)![img3](img3link)`

	res := parseImages(md)

	expected := []string{
		"[img1](img1link)",
		"[img2](img2link)",
		"[img3](img3link)",
	}

	if len(res) != 3 ||
		res[0] != expected[0] ||
		res[1] != expected[1] ||
		res[2] != expected[2] {
		t.Errorf("Incorrectly parsed touching images.\nExpected: %s\nActual: %s", strings.Join(expected, ", "), strings.Join(res, ", "))
	}
}

func TestHeadingMaker(t *testing.T) {
	res := headingMaker(3)

	if res != "### $1" {
		t.Errorf("Incorrectly make level 3 heading: %s", res)
	}

	res = headingMaker(1)

	if res != "# $1" {
		t.Errorf("Incorrectly make level 1 heading: %s", res)
	}
}
