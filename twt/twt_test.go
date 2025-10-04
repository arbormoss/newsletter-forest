package twt

import "testing"

func TestParseWithCheckbox(t *testing.T) {
	md := parse(
		`- [ ] empty checkbox
- [x] done x checkbox
- [X] done X checkbox
- [\] done \ checkbox
- [/] done / checkbox
	- [ ] empty second-teir checkbox
	- [x] done second-teir checkbox`)

	expected := `- ❎ empty checkbox
- ✅ done x checkbox
- ✅ done X checkbox
- ✅ done \ checkbox
- ✅ done / checkbox
	- ❎ empty second-teir checkbox
	- ✅ done second-teir checkbox`

	if md != expected {
		t.Errorf("Failed to parse Twt Md Checkboxes\nActual:\n%s\nExpected:\n%s", md, expected)
	}
}

func TestParseWithImages(t *testing.T) {
	md := parse(`![img](image/source.png)![img2](image/source2.png)sometext![img3](image/source3.png)![img](image/sour
ce.png)`)

	expected := `sometext![img](image/sour
ce.png)`

	if md != expected {
		t.Errorf("Failed to remove Twt Md images\nActual:\n%s\nExpected:\n%s", md, expected)
	}
}

func TestParseWithLinks(t *testing.T) {
	md := parse(`[title1](https://some.site1)[title2](https://some.site2)sometext[title3](https://some.site3)[title3](https://so
me.site3)`)

	expected := `https://some.site1https://some.site2sometexthttps://some.site3[title3](https://so
me.site3)`

	if md != expected {
		t.Errorf("Failed to parse Twt Md links\nActual:\n%s\nExpected:\n%s", md, expected)
	}
}

func TestParseWithHeadings(t *testing.T) {
	md := parse(
		`# heading 1
## heading 2
### heading 3
some text
#### heading 4
- bullet point
##### heading 5`)

	expected := `heading 1

heading 2

heading 3

some text
heading 4

- bullet point
heading 5
`

	if md != expected {
		t.Errorf("Failed to remove Twt Md headings\nActual:\n%s\nExpected:\n%s", md, expected)
	}
}
