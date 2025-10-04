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
	help := flag.Bool("h", false, "display this help text")
	confPath := flag.String("c", DEFAULT_CONF, "set the path to your configuration file")
	flag.Parse()

	if *help {
		printUsage()
		return
	}

	// read config from yaml
	conf := Config{
		Rss:     rss.RssConf{},
		Mchimp:  mchimp.MchimpConf{},
		Twitter: twt.TwitterConf{},
		Discord: dis.DiscordConf{},
	}
	confContents, err := os.ReadFile(*confPath)
	if err != nil {
		fmt.Print("Failed to read configuration file.\n")
		return
	}
	if err := yml.Unmarshal(confContents, &conf); err != nil {
		fmt.Print("Failed to parse configuration file.\n")
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		return
	}
	if len(args) > 1 {
		fmt.Print("Too many arguments: Use '-h' for help\n")
		return
	}

	articleContents, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Print("Failed to read article contents: Does the article file exist?\n")
		return
	}

	if conf.Rss.Enable {
		if err = rss.Publish(string(articleContents), conf.Rss); err != nil {
			fmt.Print("Failed to publish to RSS\n")
			fmt.Print(err.Error() + "\n")
		} else {
			fmt.Print("Published to RSS\n")
		}
	}

	if conf.Twitter.Enable {
		if err = twt.Publish(string(articleContents), conf.Twitter); err != nil {
			fmt.Print("Failed to publish to Twt\n")
			fmt.Print(err.Error() + "\n")
		} else {
			fmt.Print("Published to Twt\n")
		}
	}

	if conf.Mchimp.Enable {
		if err = mchimp.Publish(string(articleContents), conf.Mchimp); err != nil {
			fmt.Print("Failed to publish to MailChimp\n")
			fmt.Print(err.Error() + "\n")
		} else {
			fmt.Print("Published to MailChimp\n")
		}
	}

	if conf.Discord.Enable {
		if err = dis.Publish(string(articleContents), conf.Discord); err != nil {
			fmt.Print("Failed to publish to Discord\n")
			fmt.Print(err.Error() + "\n")
		} else {
			fmt.Print("Published to Discord\n")
		}
	}

	fmt.Print("=====================================================================\n")
	fmt.Print("FINISHED\n")
}

func printUsage() {
	//
	// NEWSLETTER FOREST USAGE
	// =====================================================================
	// A yaml configuration file is required to manage publishing targets.
	// See https://github.com/arbormoss/newsletter-forest for config
	// options.
	//
	// Current targets: Twitter, Discord
	// =====================================================================
	// Flags:

	fmt.Print("\n")
	fmt.Print("NEWSLETTER FOREST USAGE\n")
	fmt.Print("=====================================================================\n")
	fmt.Print("A yaml configuration file is required to manage publishing targets.\n")
	fmt.Print("See https://github.com/arbormoss/newsletter-forest for config\n")
	fmt.Print("options.\n")
	fmt.Print("\n")
	fmt.Print("Current targets: Twitter, Discord\n")
	fmt.Print("=====================================================================\n")
	fmt.Print("Flags:\n")
	flag.PrintDefaults()
	fmt.Print("\n")
}
