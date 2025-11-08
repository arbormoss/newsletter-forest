package mchimp

import (
	"fmt"
	"strconv"

	"github.com/arbormoss/newsletter-forest/markdown"
)

type MchimpConf struct {
	Enable   bool
	Key      string
	Audience string
	Dc       string
	Subject  string
	Preview  string
	From     string
	Replyto  string
}

var (
	ErrorAudienceId     = fmt.Errorf("Failed to get MailChimp Audience ID from name: Is the audience ID correct?")
	ErrorTemplateCreate = fmt.Errorf("Failed to create Mailchimp Template")
	ErrorCampeignCreate = fmt.Errorf("Failed to create Mailchimp Campeign")
	ErrorCampeignSend   = fmt.Errorf("Failed to send Mailchimp Campeign")

	mchimpMdFormat = markdown.MdFormat{
		BoldFormat:            "<strong>$1</strong>",
		ItalicFormat:          "<em>$1</em>",
		ImageFormat:           "\n<img src=\"$2\" alt=\"$1\" ><em>$1</em>\n",
		LinkFormat:            `<a href="$2">$1</a>`,
		CodeFormat:            "<code>$1</code>",
		BulletFormat:          "<ul><li>$1</li></ul>",
		BulletListPrefix:      "<ul>",
		BulletListSuffix:      "</ul>",
		DoneBulletFormat:      "- \u2705 $2",
		UncheckedBulletFormat: "- \u274E $1",
		HeadingMaker:          headingMaker,
	}
)

func Publish(article string, conf MchimpConf) error {
	audienceId, err := audienceIdFromName(conf.Audience, conf.Key, conf.Dc)
	if err != nil {
		return ErrorAudienceId
	}

	content, err := markdown.ParseMdToHtml(article, mchimpMdFormat)
	if err != nil {
		return err
	}

	templateId, err := createTemplate(conf.Key, conf.Dc, content)
	if err != nil {
		return ErrorTemplateCreate
	}

	campeignId, err := createCampeign(conf.Key, audienceId, conf.Subject, conf.Preview, conf.From, conf.Replyto, templateId, conf.Dc)
	if err != nil {
		return ErrorCampeignCreate
	}

	if err = sendCampeign(conf.Key, conf.Dc, campeignId); err != nil {
		return ErrorCampeignSend
	}

	return nil
}

func headingMaker(level int) string {
	return "\n<h" + strconv.Itoa(level) + ">$1</h" + strconv.Itoa(level) + ">\n"
}
