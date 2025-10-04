package markdown

import "regexp"

// these regex define what md elements look like for all
// platforms, regardless of the platform's md spec
var (
	Image          = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	CheckmarkDone  = regexp.MustCompile(`- \[(x|X|\\|/)\]\s(.*)`)
	CheckmarkEmpty = regexp.MustCompile(`- \[ \]\s(.*)`)
)
