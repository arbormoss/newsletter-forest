package markdown

import (
	"regexp"
	"strings"
)

type MdFormat struct {
	BoldFormat   string
	ItalicFormat string

	ImageFormat string
	LinkFormat  string

	CodeFormat string

	BulletFormat     string
	BulletListPrefix string
	BulletListSuffix string

	DoneBulletFormat      string
	UncheckedBulletFormat string

	HeadingMaker MdHeadingMaker
}

// this is not fully functional, but has the features i use most
func ParseMdToHtml(md string, format MdFormat) string {
	// quick check for html characters
	md = escapeCharacters(md)
	md = strings.TrimSpace(md)

	md = parseBoldItalics(md, format.BoldFormat, format.ItalicFormat)
	md = parseImages(md, format.ImageFormat)
	md = parseHeadings(md, format.HeadingMaker)
	md = parseCode(md, format.CodeFormat)

	// TODO: seperate bullet fmt
	md = parseBullets(md, format.BulletFormat, format.DoneBulletFormat, format.UncheckedBulletFormat, format.BulletListPrefix, format.BulletListSuffix)

	md = parseLinks(md, format.LinkFormat)

	return md
}

// uses the standard md link for external pages
func parseLinks(md, format string) string {
	md = Link.ReplaceAllString(md, format)

	return md
}

func parseImages(md, format string) string {
	md = Image.ReplaceAllString(md, format)

	return md
}

type MdHeadingMaker func(int) string

func parseHeadings(md string, maker MdHeadingMaker) string {
	for i := 5; i > 0; i-- {
		regex := regexp.MustCompile(strings.Repeat("#", i) + `\s(.*)`)
		md = regex.ReplaceAllString(md, maker(i))
	}

	return md
}

// TODO: multi level bullet points
func parseBullets(md, bulletFmt, doneFmt, uncheckFmt, rmPrefix, rmSuffix string) string {
	md = CheckmarkDone.ReplaceAllString(md, doneFmt)
	md = CheckmarkEmpty.ReplaceAllString(md, uncheckFmt)
	md = Bullet.ReplaceAllString(md, bulletFmt)
	var BulletFix = regexp.MustCompile(rmSuffix + `\s*` + rmPrefix)
	md = BulletFix.ReplaceAllString(md, "\n")

	return md
}

// only supports single tics
func parseCode(md, format string) string {
	md = Code.ReplaceAllString(md, format)

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

func parseBoldItalics(md, boldFormat, italicFormat string) string {
	md = Bold1.ReplaceAllString(md, boldFormat)
	md = Bold2.ReplaceAllString(md, boldFormat)

	md = Italic1.ReplaceAllString(md, italicFormat)
	md = Italic2.ReplaceAllString(md, italicFormat)

	return md
}
