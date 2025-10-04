package dis

import (
	"testing"
)

func TestParseCheckboxes(t *testing.T) {
	md := parse(
		`- [ ] empty checkbox
- [x] done x checkbox
- [X] done X checkbox
- [\] done \ checkbox
- [/] done / checkbox
	- [ ] empty second-teir checkbox
	- [x] done second-teir checkbox`)

	expected := `- :negative_squared_cross_mark: empty checkbox
- :white_check_mark: done x checkbox
- :white_check_mark: done X checkbox
- :white_check_mark: done \ checkbox
- :white_check_mark: done / checkbox
	- :negative_squared_cross_mark: empty second-teir checkbox
	- :white_check_mark: done second-teir checkbox`

	if md != expected {
		t.Errorf("Failed to parse Discord Md Checkboxes\nActual:\n%s\nExpected:\n%s", md, expected)
	}
}

func TestParseWithImages(t *testing.T) {
	md := parse(`![img](image/source.png)![img2](image/source2.png)sometext![img3](image/source3.png)![img](image/sour
ce.png)`)

	expected := `sometext![img](image/sour
ce.png)`

	if md != expected {
		t.Errorf("Failed to remove images in Discord Md Parsing\nActual:\n%s\nExpected:\n%s", md, expected)
	}
}
