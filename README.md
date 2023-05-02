# urleap

URL shortener via bit.ly, etc.

[![build](https://github.com/tamada/urleap/actions/workflows/build.yml/badge.svg)](https://github.com/tamada/urleap/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/tamada/urleap/badge.svg?branch=main)](https://coveralls.io/github/tamada/urleap?branch=main)
[![codebeat badge](https://codebeat.co/badges/d63e3c67-fc5d-4f27-9e81-d80861d60c20)](https://codebeat.co/projects/github-com-tamada-urleap-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/tamada/urleap)](https://goreportcard.com/report/github.com/tamada/urleap)

![MIT License](https://img.shields.io/badge/Licnese-MIT%20License-informational)
![Version](https://img.shields.io/badge/Version-0.1.14-informational)

## :speaking_head: Overview

There are several URL shortening services such as [bit.ly](https://bit.ly) and [TinyURL](https://tinyurl.com/app). To use these services, users need to open a browser and enter the URL. We would like to make it easier to shorten the url from the CLI.

## :runner: Usage

```sh
urleap [OPTIONS] [URL...]
OPTIONS
  -t, --token <TOKEN>      specify the token for the service. This option is mandatory.
  -q, --qrcode <FILE>      include QR-code of the URL in the output.
  -c, --config <CONFIG>    specify the configuration file.
  -h, --help               print this mesasge and exit.
  -v, --version            print the version and exit.
ARGUMENT
  URL     specify the url for shortening. this arguments accept multiple values.
          if no arguments were specified, urleap prints the list of available shorten urls.
```

### :gear: Configuration File

#### Location

`urleap` reads the following list of files in order and overwrites the settings. 
If the file does not exist, it is simply ignored.

* `/opt/homebrew/opt/urleap/config.json`

* `/usr/local/opt/urleap/config.json`

* `$URLEAP_HOME/config.json`
* `~/.config/urleap/config.json`
* `./.urleap.json`
* files specified by `--config` option

#### Format

```json
{
  "provider": {
    "api": "bit.ly",
    "api_version": "v4",
  }
}
```

## :anchor: Installation

### :beer: Homebrew

```sh
brew tap tamada/brew
brew install urleap
```

### :whale: Docker

```sh
docker run -it --rm tamada/urleap:latest -t <token> <url...>
```

#### tags

* `0.1.2`, `latest`

## :smile: About

### :scroll: License

![MIT License](https://img.shields.io/badge/Licnese-MIT%20License-informational)

* Permitted
  * 🙆‍♀️ Commercial use
  * 🙆‍♀️ Modification
  * 🙆‍♀️ Distribution
  * 🙆‍♀️ Private use
* Limitations
  * 🙅‍♂️ Liability
  * 🙅‍♂️ Warranty

### :man_office_worker: Developers:woman_office_worker:

* Haruaki Tamada [:octocat:](https://github.com/tamada)

### :jack_o_lantern: Icon

![Icon](docs/static/images/urleap.svg)
