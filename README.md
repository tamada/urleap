# urleap

URL shortener via bit.ly, etc.

![MIT License](https://img.shields.io/badge/Licnese-MIT%20License-informational)

![0.1.0](https://img.shields.io/badge/Version-0.1.0-informational)

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
