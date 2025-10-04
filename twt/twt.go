package twt

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/arbormoss/newsletter-forest/markdown"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

type TwitterConf struct {
	Enable      bool
	Token       string
	Tokensecret string
	Key         string
	Keysecret   string
}

var (
	ErrorTwtClient      = fmt.Errorf("Failed to create twt client")
	ErrorTwtTweetCreate = fmt.Errorf("Failed to send twt tweet")

	ErrorTwtNoToken       = fmt.Errorf("Twt missing client token")
	ErrorTwtNoTokensecret = fmt.Errorf("Twt missing client token secret")

	ErrorTwtNoKey       = fmt.Errorf("Twt missing client key")
	ErrorTwtNoKeysecret = fmt.Errorf("Twt missing client key secret")
)

func Publish(article string, conf TwitterConf) error {
	if conf.Key == "" {
		return ErrorTwtNoKey
	}
	if conf.Keysecret == "" {
		return ErrorTwtNoKeysecret
	}
	if conf.Token == "" {
		return ErrorTwtNoToken
	}
	if conf.Tokensecret == "" {
		return ErrorTwtNoTokensecret
	}

	os.Setenv("GOTWI_API_KEY", conf.Key)
	os.Setenv("GOTWI_API_KEY_SECRET", conf.Keysecret)

	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           conf.Token,
		OAuthTokenSecret:     conf.Tokensecret,
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		return ErrorTwtClient
	}

	p := &types.CreateInput{
		Text: gotwi.String(parse(article)),
	}

	_, err = managetweet.Create(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return ErrorTwtTweetCreate
	}

	return nil
}

func parse(article string) string {
	article = parseCheckboxes(article)
	article = removeImages(article)
	article = removeHeadings(article)
	article = parseLinks(article)

	return article
}

func parseCheckboxes(md string) string {
	md = markdown.CheckmarkDone.ReplaceAllString(md, "- ✅ $2")
	md = markdown.CheckmarkEmpty.ReplaceAllString(md, "- ❎ $1")

	return md
}

func parseLinks(md string) string {
	md = markdown.Link.ReplaceAllString(md, "$2")
	return md
}

func removeImages(md string) string {
	md = markdown.Image.ReplaceAllString(md, ``)
	return md
}

func removeHeadings(md string) string {
	for i := 5; i > 0; i-- {
		regex := regexp.MustCompile(strings.Repeat("#", i) + `\s(.*)`)
		md = regex.ReplaceAllString(md, "$1\n")
	}

	return md
}
