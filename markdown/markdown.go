package markdown

import "regexp"

var (
	Image          = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	CheckmarkDone  = regexp.MustCompile(`- \[(x|X|\\|/)\]\s(.*)`)
	CheckmarkEmpty = regexp.MustCompile(`- \[ \]\s(.*)`)
)
