package dis

import (
	"fmt"
	"strings"

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

	content := markdown.ParseMdToHtml(article, markdown.MdFormat{
		BoldFormat:            "**$1**",
		ItalicFormat:          "*$1*",
		ImageFormat:           "",
		LinkFormat:            "[$1]($2)\n",
		CodeFormat:            "`$1`",
		BulletFormat:          "- $1",
		BulletListPrefix:      "<ul>",
		BulletListSuffix:      "</ul>",
		DoneBulletFormat:      "- \u2705 $2",
		UncheckedBulletFormat: "- \u274E $1",
		HeadingMaker:          headingMaker,
	})

	session.ChannelMessageSend(conf.Channel, content)

	err = session.Close()
	if err != nil {
		return ErrorDiscordShutdown
	}

	return nil
}

// images must be parsed as sperate messages to support discord md
func parseImages(md string) []string {
	images := markdown.Image.FindAllString(md, -1)

	for i, img := range images {
		images[i] = markdown.Image.ReplaceAllString(img, `[$1]($2)`)
	}

	return images
}

func headingMaker(level int) string {
	return strings.Repeat("#", level) + " $1"
}
