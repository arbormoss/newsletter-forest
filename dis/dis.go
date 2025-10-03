package dis

import "fmt"

type DiscordConf struct {
	Enable bool
}

func Publish(article string, conf DiscordConf) {
	fmt.Print(parse(article))
}

func parse(article string) string {
	return article
}
