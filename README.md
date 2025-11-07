# Newsletter Forest

This is a program for distributing a Markdown newsletter to multiple social media feeds.
It is intended to be used by [ACM UMN](https://acm.umn.edu/) for distributing their newsletter more easily.

## Usage

### Build

1. Install [golang](https://go.dev/).
2. Run `go build`

### Test

Run `./test.sh` in the project root to test all go packages.
This has no additional dependencies other than golang and it's go libraries.

### Publish

To publish an article just point the program at it.

`./newsletter-forest <md article>`

Flags:
- `-h`: help
- `-c`: path to yaml configuration file (required, defaults to `./conf.yaml`)

### Configuration

To publish an article to a specific target, add it to a yaml config file.
Make sure to set `enable` to true to enable publishing to the target.
For most targets, you need to set a token or key of some kind for authentication.
To enable to disable a target without fully removing it from the configuration, set `enable`.

Example configuration:
```yaml
---
discord:
  enable: false
  channel: "<discord channel id>" # get this by enabling dev mode
  token: "<discord announcement bot token>" # this bot will need to be in the server
twitter: # the twitter api makes no sense anymore
  enable: false
  token: "<twitter access token>"
  tokensecret: "<twitter access token secret>" # this *should* match what it's called in the dev console
  key: "<twitter api key>"
  keysecret: "<twitter api key secret>"
mchimp: # MailChimp
  enable: false
  key: "<mailchimp api key>"
  audience: "<mailchimp audience name>" # you can find this on your MailChimp dashboard
  dc: "<mailchimp region>" # you can easily find this in your admin dashboard url: https://<region>.admin.mailchimp.com/
  subject: "<email subject>"
  from: "<sender name>" # a name not an email
  replyto: "<sender email>" # this must be a verified email on your mailchimp
```

## Markdown Support by Platform

See [MD SUPPORT](MD_SUPPORT.md).

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## Licence

[MIT](LICENSE)

[Credits](CREDITS)
