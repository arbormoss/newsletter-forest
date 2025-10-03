package main

import (
	"flag"
	"log"
	"os"

	"github.com/arbormoss/newsletter-forest/rss"
	yml "github.com/goccy/go-yaml"
)

type Config struct {
	Rss bool
}

const DEFAULT_CONF = "./conf.yaml"

func main() {
	// set up command line parsing
	confPath := flag.String("c", DEFAULT_CONF, "path to the config file")
	flag.Parse()

	// read config from yaml
	conf := Config{
		Rss: false,
	}
	confContents, err := os.ReadFile(*confPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := yml.Unmarshal(confContents, &conf); err != nil {
		log.Fatal(err)
	}

	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Too many arguments: Use '-h' for help")
	}

	articleContents, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal("Failed to read article contents: Does it exist?")
	}

	if conf.Rss {
		rss.Publish(string(articleContents))
	}
}
