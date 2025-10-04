package dis

import (
	"fmt"

	"github.com/arbormoss/newsletter-forest/markdown"
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

	// images need to be parsed as seperate messages because discord does
	// not nicely allow inline images
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

// discord has it's own md support but is missing a lot
// of features. This parses in checkboxes and removes images.
// The images have to be sent as seperate messages.
func parse(article string) string {
	article = removeImages(article)
	article = parseCheckboxes(article)

	return article
}

func parseImages(md string) []string {
	images := markdown.Image.FindAllString(md, -1)

	for i, img := range images {
		images[i] = markdown.Image.ReplaceAllString(img, `[$1]($2)`)
	}

	return images
}

func removeImages(md string) string {
	md = markdown.Image.ReplaceAllString(md, ``)

	return md
}

func parseCheckboxes(md string) string {
	md = markdown.CheckmarkDone.ReplaceAllString(md, "- :white_check_mark: $2")
	md = markdown.CheckmarkEmpty.ReplaceAllString(md, "- :negative_squared_cross_mark: $1")

	return md
}
