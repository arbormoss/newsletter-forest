package markdown

import "regexp"

// these regex define what md elements look like for all
// platforms, regardless of the platform's md spec
var (
	Image = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	Link  = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	CheckmarkDone  = regexp.MustCompile(`- \[(x|X|\\|/)\]\s(.*)`)
	CheckmarkEmpty = regexp.MustCompile(`- \[ \]\s(.*)`)

	Code = regexp.MustCompile("`((.|\n)*?)`")

	Bold1   = regexp.MustCompile(`\*(.*?)\*`)
	Bold2   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	Italic1 = regexp.MustCompile(`\_(.*?)\_`)
	Italic2 = regexp.MustCompile(`\_\_(.*?)\_\_`)

	Bullet = regexp.MustCompile(`-\s(.*)`)
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
