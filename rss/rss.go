package rss

import "fmt"

type RssConf struct {
	Enable bool
}

func Publish(article string, conf RssConf) error {
	fmt.Print(parse(article))
	return nil
}

func parse(article string) string {
	return article
}
