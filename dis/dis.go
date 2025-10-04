package dis

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordConf struct {
	Enable  bool
	Channel string
	Token   string
}

var (
	ErrorDiscordSession  = fmt.Errorf("Error: Failed to open discord session")
	ErrorDiscordShutdown = fmt.Errorf("Error: Failed gracefull discord shutdown")
)

func Publish(article string, conf DiscordConf) error {
	session, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		return ErrorDiscordSession
	}

	if err = session.Open(); err != nil {
		return ErrorDiscordSession
	}

	session.ChannelMessageSend(conf.Channel, parse(article))

	err = session.Close()
	if err != nil {
		return ErrorDiscordShutdown
	}

	return nil
}

func parse(article string) string {
	return article
}
