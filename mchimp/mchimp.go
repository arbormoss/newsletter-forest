package mchimp

import "fmt"

type MchimpConf struct {
	Enable bool
}

func Publish(article string, conf MchimpConf) error {
	fmt.Print(parse(article))
	return nil
}

func parse(article string) string {
	return article
}
