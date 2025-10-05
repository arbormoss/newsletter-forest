package markdown

import "regexp"

// these regex define what md elements look like for all
// platforms, regardless of the platform's md spec
var (
	Image = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	Link  = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	CheckmarkDone  = regexp.MustCompile(`- \[(x|X|\\|/)\]\s(.*)`)
	CheckmarkEmpty = regexp.MustCompile(`- \[ \]\s(.*)`)

	Code         = regexp.MustCompile("`((.|\n)*?)`")
	CodePadding1 = regexp.MustCompile(`</code>\s*\n\s*<code>`)
	CodePadding2 = regexp.MustCompile(`</code>\s*<code>`)

	Bold1   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	Bold2   = regexp.MustCompile(`\_\_(.*?)\_\_`)
	Italic1 = regexp.MustCompile(`\*(.*?)\*`)
	Italic2 = regexp.MustCompile(`\_(.*?)\_`)

	Bullet    = regexp.MustCompile(`-\s(.*)`)
	BulletFix = regexp.MustCompile(`</ul>\s*<ul>`)
)
