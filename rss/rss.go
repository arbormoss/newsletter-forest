package rss

import "fmt"

func Publish(article string) {
	fmt.Print(parse(article))
}

func parse(article string) string {
	return article
}
