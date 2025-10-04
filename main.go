package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arbormoss/newsletter-forest/dis"
	"github.com/arbormoss/newsletter-forest/mchimp"
	"github.com/arbormoss/newsletter-forest/rss"
	"github.com/arbormoss/newsletter-forest/twt"
	yml "github.com/goccy/go-yaml"
)

type Config struct {
	Rss     rss.RssConf
	Mchimp  mchimp.MchimpConf
	Twitter twt.TwitterConf
	Discord dis.DiscordConf
}

const DEFAULT_CONF = "./conf.yaml"

func main() {
	// set up command line parsing
	confPath := flag.String("c", DEFAULT_CONF, "path to the config file")
	flag.Parse()

	// read config from yaml
	conf := Config{
		Rss:     rss.RssConf{},
		Mchimp:  mchimp.MchimpConf{},
		Twitter: twt.TwitterConf{},
		Discord: dis.DiscordConf{},
	}
	confContents, err := os.ReadFile(*confPath)
	if err != nil {
		fmt.Print("Failed to read configuration file.")
		return
	}
	if err := yml.Unmarshal(confContents, &conf); err != nil {
		fmt.Print("Failed to parse configuration file.")
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Print("Not enough arguments: Use '-h' for help")
		return
	}
	if len(args) > 1 {
		fmt.Print("Too many arguments: Use '-h' for help")
		return
	}

	articleContents, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Print("Failed to read article contents: Does the article file exist?")
		return
	}

	if conf.Rss.Enable {
		if err = rss.Publish(string(articleContents), conf.Rss); err != nil {
			fmt.Print("Failed to publish to RSS")
			fmt.Print(err)
		}
		fmt.Print("Published to RSS")
	}

	if conf.Twitter.Enable {
		if err = twt.Publish(string(articleContents), conf.Twitter); err != nil {
			fmt.Print("Failed to publish to Twt")
			fmt.Print(err)
		}
		fmt.Print("Published to Twt")
	}

	if conf.Mchimp.Enable {
		if err = mchimp.Publish(string(articleContents), conf.Mchimp); err != nil {
			fmt.Print("Failed to publish to MailChimp")
			fmt.Print(err)
		}
		fmt.Print("Published to MailChimp")
	}

	if conf.Discord.Enable {
		if err = dis.Publish(string(articleContents), conf.Discord); err != nil {
			fmt.Print("Failed to publish to Discord")
			fmt.Print(err)
		}
		fmt.Print("Published to Discord")
	}
}
