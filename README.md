# Newsletter Forest

This is a program for distributing a Markdown newsletter to multiple social media feeds.
It is intended to be used by [ACM UMN](https://acm.umn.edu/) for distributing their newsletter more easily.

## UNDER CONSTRUCTION

See [TODO](TODO.md).

## Usage

### Build

1. Install [golang](https://go.dev/).
2. Run `go build`

### Test

Either run `./test.sh` or run `go test ./...` in the project root to test all go packages.

### Publish

`./newsletter-forest <md article>`

Flags:
- `-h`: help
- `-c`: path to yaml configuration file (defaults to `./conf.yaml`)

### Configuration

To publish an article to a specific target, add it to a yaml config file.
Make sure to set `enable` to true to enable publishing to the target.
For most targets, you need to set a token or key of some kind for authenticaiton.
To enable to disable a target without fully removing it from the configuration, set `enable`.

Example configuration:
```yaml
---
discord:
  enable: false
  channel: "<discord channel id>"
  token: "<discord announcement bot token>"
twitter:
  enable: false
  token: "<twitter access token>"
  tokensecret: "<twitter access token secret>"
  key: "<twitter api key>"
  keysecret: "<twitter api key secret>"
```

## Markdown Support by Platform

See [MD SUPPORT](MD_SUPPORT.md).

## Contributing

Pull requests are welcome.
Either open an issue or chat with me before opening one.
