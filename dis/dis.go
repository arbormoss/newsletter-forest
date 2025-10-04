package dis

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

type DiscordConf struct {
	Enable  bool
	Channel string
	Token   string
}

var (
	ErrorDiscordSession  = fmt.Errorf("Error: Failed to open discord session")
	ErrorDiscordShutdown = fmt.Errorf("Error: Failed graceful discord shutdown")
)

func Publish(article string, conf DiscordConf) error {
	session, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		return ErrorDiscordSession
	}

	if err = session.Open(); err != nil {
		return ErrorDiscordSession
	}

	for _, img := range parseImages(article) {
		session.ChannelMessageSend(conf.Channel, img)
	}
	session.ChannelMessageSend(conf.Channel, parse(article))

	err = session.Close()
	if err != nil {
		return ErrorDiscordShutdown
	}

	return nil
}

func parse(article string) string {
	article = removeImages(article)
	article = parseCheckboxes(article)

	return article
}

func parseImages(md string) []string {
	regex := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	images := regex.FindAllString(md, -1)

	for i, img := range images {
		images[i] = regex.ReplaceAllString(img, `[$1]($2)`)
	}

	return images
}

func removeImages(md string) string {
	regex := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	md = regex.ReplaceAllString(md, ``)

	return md
}

func parseCheckboxes(md string) string {
	regex := regexp.MustCompile(`- \[(x|X|\\|/)\]\s(.*)`)
	md = regex.ReplaceAllString(md, "- :white_check_mark: $2")

	regex = regexp.MustCompile(`- \[ \]\s(.*)`)
	md = regex.ReplaceAllString(md, "- :negative_squared_cross_mark: $1")

	return md
}
