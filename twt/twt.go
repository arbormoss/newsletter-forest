package twt

import "fmt"

type TwitterConf struct {
	Enable bool
}

func Publish(article string, conf TwitterConf) {
	fmt.Print(parse(article))
}

func parse(article string) string {
	return article
}
