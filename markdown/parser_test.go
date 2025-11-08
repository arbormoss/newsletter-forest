package markdown

import (
	"strconv"
	"testing"
)

func TestParseLinks(t *testing.T) {
	md := `[img1](img1link)[img2](img2link) hiii[img3](img3link)`

	res := parseLinks(md, "<begin link>$1<between link>$2<end link>")

	expected := `<begin link>img1<between link>img1link<end link><begin link>img2<between link>img2link<end link> hiii<begin link>img3<between link>img3link<end link>`

	if res != expected {
		t.Errorf("Failed to parse links.\nExpected: %s\nActual: %s\n", expected, res)
	}
}

func TestParseImages(t *testing.T) {
	md := `![img1](img1link)![img2](img2link) hiii![img3](img3link)`

	res := parseImages(md, "<begin image>$1<between image>$2<end image>")

	expected := `<begin image>img1<between image>img1link<end image><begin image>img2<between image>img2link<end image> hiii<begin image>img3<between image>img3link<end image>`

	if res != expected {
		t.Errorf("Failed to parse images.\nExpected: %s\nActual: %s\n", expected, res)
	}
}

func TestParseHeadings(t *testing.T) {
	md := `### third level
## second
space
##### fifth
hi
# first`

	res := parseHeadings(md, func(level int) string {
		return "<begin heading " + strconv.Itoa(level) + ">$1<end heading " + strconv.Itoa(level) + ">"
	})

	expected := `<begin heading 3>third level<end heading 3>
<begin heading 2>second<end heading 2>
space
<begin heading 5>fifth<end heading 5>
hi
<begin heading 1>first<end heading 1>`

	if res != expected {
		t.Errorf("Failed to parse headings.\nExpected: %s\nActual: %s\n", expected, res)
	}
}

func TestParseBullets(t *testing.T) {
	md := `-not a bullet
- this is a bullet
- a second
	- multi level
- back down`

	res := parseBullets(md, "<start bullet>$1<end bullet>", "<start done bullet>$1<end done bullet>", "<start uncheck bullet>$1<end uncheck bullet>", "<ul>", "</ul>")

	expected := `-not a bullet
<start bullet>this is a bullet<end bullet>
<start bullet>a second<end bullet>
	<start bullet>multi level<end bullet>
<start bullet>back down<end bullet>`

	if res != expected {
		t.Errorf("Failed to parse bullet points.\nExpected: %s\nActual: %s\n", expected, res)
	}
}

func TestParseCode(t *testing.T) {
	md := "this is some `code right here` isnt that cool"

	res := parseCode(md, "<start code>$1<end code>")

	expected := "this is some <start code>code right here<end code> isnt that cool"

	if res != expected {
		t.Errorf("Failed to parse inline code.\nExpected: %s\nActual: %s\n", expected, res)
	}
}

func TestEscapeCharacters(t *testing.T) {
	md := `<><%&FDJKA'"`

	res := escapeCharacters(md)

	expected := `&lt;&gt;&lt;%&amp;FDJKA&#39;&quot;`

	if res != expected {
		t.Errorf("Failed to escape html.\nExpected: %s\nActual: %s\n", expected, res)
	}
}

func TestUnescapeCharacters(t *testing.T) {
	md := `&quot;&lt;&gt;%&amp;FDJKA&#39;&lt;`

	res := unescapeCharacters(md)

	expected := `"<>%&FDJKA'<`

	if res != expected {
		t.Errorf("Failed to unescape html.\nExpected: %s\nActual: %s\n", expected, res)
	}
}

func TestParseBoldItalics(t *testing.T) {
	md := `md **bold** _italic_ *bold* __italic__
_italic_**bold and this is _italic_ bold**`

	res := parseBoldItalics(md, "<start bold>$1<end bold>", "<start italic>$1<end italic>")

	expected := `md <start bold>bold<end bold> <start italic>italic<end italic> <start bold>bold<end bold> <start italic>italic<end italic>
<start italic>italic<end italic><start bold>bold and this is <start italic>italic<end italic> bold<end bold>`

	if res != expected {
		t.Errorf("Failed to parse bold and italics.\nExpected: %s\nActual: %s\n", expected, res)
	}
}
