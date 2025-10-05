package mchimp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// given the name of a mailchimp audience, find it's Id
func audienceIdFromName(name, apiKey, dc string) (string, error) {
	req, err := http.NewRequest("GET", "https://"+dc+".api.mailchimp.com/3.0/audiences", nil)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth("anystring", apiKey)

	res, err := http.DefaultClient.Do(req)

	// this parses the json as a map of strings to anys
	// this is hard to debug and should be replaced by library calls in the future
	var respjson map[string]interface{}
	body, err := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &respjson); err != nil {
		return "", err
	}

	audienceId, err := parseAudienceId(respjson, name)
	if err != nil {
		return "", err
	}

	return audienceId, nil
}

// parses an audience id from a json object of all audiences
func parseAudienceId(json map[string]interface{}, audienceName string) (string, error) {
	audiences := json["audiences"].([]interface{})

	// magic most horrific and vile
	//
	// looks through the audience json for a matching name and id
	for _, audience := range audiences {
		if audience.(map[string]interface{})["name"] == audienceName {
			id := audience.(map[string]interface{})["id"].(string)
			if id == "" {
				return "", ErrorAudienceId
			}

			return id, nil
		}
	}
	return "", ErrorAudienceId
}

// creates a mailchimp template, which is required to send an email.
// the template content is just some html that is the body of the email.
func createTemplate(apiKey, dc, content string) (string, error) {
	content = escapeReq(content)

	var data = strings.NewReader(`{"name" : "default" , "html" : "<div>` + content + `</div>" }`)
	req, err := http.NewRequest("POST", "https://"+dc+".api.mailchimp.com/3.0/templates", data)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var respjson map[string]interface{}

	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &respjson); err != nil {
		return "", err
	}

	if respjson["id"] == nil {
		log.Fatal(respjson)
	}

	id := int(respjson["id"].(float64))

	return strconv.Itoa(id), nil
}

func createCampeign(apiKey, audienceId, subject, previewText, from, replyTo, templateId, dc string) (string, error) {
	var data = strings.NewReader(`{"type":"regular","recipients":{"list_id":"` + audienceId + `"},"settings":{"subject_line":"` + subject + `","preview_text":"` + previewText + `","title":"` + subject + `","from_name":"` + from + `","reply_to":"` + replyTo + `", "use_conversation":false,"authenticate":false,"auto_footer":false,"inline_css":false,"auto_fb_post":[],"fb_comments":false,"template_id":` + templateId + `},"content_type":"template"}`)
	req, err := http.NewRequest("POST", "https://"+dc+".api.mailchimp.com/3.0/campaigns", data)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var respjson map[string]interface{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &respjson); err != nil {
		return "", err
	}

	id := respjson["id"].(string)
	return id, nil
}

func sendCampeign(apiKey, dc, campeignId string) error {
	req, err := http.NewRequest("POST", "https://"+dc+".api.mailchimp.com/3.0/campaigns/"+campeignId+"/actions/send", nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth("anystring", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

// prep content for sending in http req
func escapeReq(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, `/`, `\/`)
	s = strings.ReplaceAll(s, "\b", "")
	s = strings.ReplaceAll(s, "\f", "")
	s = strings.ReplaceAll(s, "\r", "<br>")
	s = strings.ReplaceAll(s, "\t", "  ")
	s = strings.ReplaceAll(s, "\n", "<br>")

	return s
}
