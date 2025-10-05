package mchimp

import (
	"fmt"
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
)

func Publish(article string, conf MchimpConf) error {
	audienceId, err := audienceIdFromName(conf.Audience, conf.Key, conf.Dc)
	if err != nil {
		return ErrorAudienceId
	}

	templateId, err := createTemplate(conf.Key, conf.Dc, parse(article))
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

// todo, parse md to html
func parse(article string) string {
	return article
}
