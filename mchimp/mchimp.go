package mchimp

import "fmt"

type MchimpConf struct {
	Enable bool
}

func Publish(article string, conf MchimpConf) {
	fmt.Print(parse(article))
}

func parse(article string) string {
	return article
}
