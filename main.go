package main

import (
	"flag"
	"log"
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
		log.Fatal(err)
	}
	if err := yml.Unmarshal(confContents, &conf); err != nil {
		log.Fatal(err)
	}

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Not enough arguments: Use '-h' for help")
	}
	if len(args) > 1 {
		log.Fatal("Too many arguments: Use '-h' for help")
	}

	articleContents, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal("Failed to read article contents: Does it exist?")
	}

	if conf.Rss.Enable {
		rss.Publish(string(articleContents), conf.Rss)
	}
	if conf.Twitter.Enable {
		if err = twt.Publish(string(articleContents), conf.Twitter); err != nil {
			log.Fatal(err)
		}

	}
	if conf.Mchimp.Enable {
		mchimp.Publish(string(articleContents), conf.Mchimp)
	}
	if conf.Discord.Enable {
		dis.Publish(string(articleContents), conf.Discord)
	}
}
