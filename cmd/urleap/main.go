package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	"github.com/tamada/urleap"
)

const VERSION = "0.1.16"

/*
helpMessage prints the help message.
This function is used in the small tests, so it may be called with a zero-length slice.
*/
func helpMessage(args []string) string {
	prog := "urleap"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [URLs...]
OPTIONS
    -t, --token <TOKEN>      specify the token for the service. This option is mandatory.
    -q, --qrcode <FILE>      include QR-code of the URL in the output.
    -c, --config <CONFIG>    specify the configuration file.
    -g, --group <GROUP>      specify the group name for the service. Default is "urleap"
    -d, --delete             delete the specified shorten URL.
    -h, --help               print this mesasge and exit.
    -v, --version            print the version and exit.
ARGUMENT
    URL     specify the url for shortening. this arguments accept multiple values.
            if no arguments were specified, urleap prints the list of available shorten urls.`, prog)
}

type UrleapError struct {
	statusCode int
	message    string
}

func (e UrleapError) Error() string {
	return e.message
}

type flags struct {
	deleteFlag  bool
	helpFlag    bool
	versionFlag bool
}

/*
This struct holds the values of the options.
*/
type options struct {
	token   string
	qrcode  string
	config  string
	group   string
	flagSet *flags
}

func newOptions() *options {
	return &options{flagSet: &flags{}}
}

/*
Define the options and return the pointer to the options and the pointer to the flagset.
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.token, "token", "t", "", "specify the token for the service. This option is mandatory.")
	flags.StringVarP(&opts.qrcode, "qrcode", "q", "", "include QR-code of the URL in the output.")
	flags.StringVarP(&opts.config, "config", "c", "", "specify the configuration file.")
	flags.StringVarP(&opts.group, "group", "g", "", "specify the group name for the service. Default is \"urleap\"")
	flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "delete the specified shorten URL.")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "print this mesasge and exit.")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "print the version and exit.")
	return opts, flags
}

/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *UrleapError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &UrleapError{statusCode: 0, message: ""}
	}
	if opts.token == "" {
		return nil, nil, &UrleapError{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}

func performEach(bitly *urleap.Bitly, opts *options, config *urleap.Config, url string) error {
	if opts.flagSet.deleteFlag {
		return bitly.Delete(config, url)
	} else {
		result, err := bitly.Shorten(config, url)
		if err != nil {
			return err
		}
		fmt.Println(result)
	}
	return nil
}

func perform(opts *options, args []string) *UrleapError {
	bitly := urleap.NewBitly(opts.group)
	config := urleap.NewConfig(opts.token)
	for _, url := range args {
		err := performEach(bitly, opts, config, url)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if len(args) == 0 {
		urls, err := bitly.List(config)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, url := range urls {
			fmt.Println(url)
		}
	}
	return nil
}

func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
