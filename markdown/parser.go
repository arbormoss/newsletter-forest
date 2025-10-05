package markdown

import (
	"regexp"
	"strconv"
	"strings"
)

// this is not fully functional, but has the features i use most
// parses the md page files into html
func ParseMdToHtml(md string) string {
	// quick check for html characters
	md = escapeCharacters(md)
	md = strings.TrimSpace(md)

	md = parseBoldItalics(md)
	md = parseImages(md)
	md = parseHeadings(md)
	md = parseParagraphs(md)
	md = parseCode(md)

	md = parseBullets(md)

	// must be after images
	md = parseLinks(md)

	return md
}

// uses the standard md link for external pages
func parseLinks(md string) string {
	md = Link.ReplaceAllString(md, `<a href="$2">$1</a>`)

	return md
}

func parseImages(md string) string {
	md = Image.ReplaceAllString(md, "\n<img src=\"$2\" alt=\"$1\" ><em>$1</em>\n")

	return md
}

func parseHeadings(md string) string {
	for i := 5; i > 0; i-- {
		regex := regexp.MustCompile(strings.Repeat("#", i) + `\s(.*)`)
		md = regex.ReplaceAllString(md, "\n<h"+strconv.Itoa(i)+">$1</h"+strconv.Itoa(i)+">\n")
	}

	return md
}

// TODO: multi level bullet points
func parseBullets(md string) string {
	md = Bullet.ReplaceAllString(md, "<ul><li>$1</li></ul>")
	md = BulletFix.ReplaceAllString(md, "\n")

	return md
}

// triple back ticks get a seperate line while single ones are inline
func parseCode(md string) string {
	md = Code.ReplaceAllString(md, "<code>$1</code>")

	// ensure consecuteive code blocks do not have added padding
	md = CodePadding1.ReplaceAllString(md, "\n")
	md = CodePadding2.ReplaceAllString(md, " ")

	return md
}

// escape common html problems
func escapeCharacters(md string) string {
	md = strings.ReplaceAll(md, "&", "&amp;")
	md = strings.ReplaceAll(md, "<", "&lt;")
	md = strings.ReplaceAll(md, ">", "&gt;")
	md = strings.ReplaceAll(md, "\"", "&quot;")
	md = strings.ReplaceAll(md, "'", "&#39;")
	return md
}

func parseBoldItalics(md string) string {
	md = Bold1.ReplaceAllString(md, "<strong>$1</strong>")
	md = Bold2.ReplaceAllString(md, "<strong>$1</strong>")

	md = Italic1.ReplaceAllString(md, "<em>$1</em>")
	md = Italic2.ReplaceAllString(md, "<em>$1</em>")

	return md
}

func parseParagraphs(md string) string {
	regex := regexp.MustCompile(`(.)[\f\t\v]*\n[\f\t\v]*(.)`)
	md = regex.ReplaceAllString(md, "$1$2")

	regex = regexp.MustCompile(`(</h\d>)[\f\t\v]*\n[\s]*(.)`)
	md = regex.ReplaceAllString(md, "$1$2")

	regex = regexp.MustCompile(`(.)[\f\t\v]*\n([\f\t\v]*\n)+(.)`)
	md = regex.ReplaceAllString(md, "$1\n\n$3")

	return md
}
