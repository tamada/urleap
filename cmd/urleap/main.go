package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
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

/*
This struct holds the values of the options.
*/
type options struct {
	token   string
	qrcode  string
	config  string
	help    bool
	version bool
}

/*
Define the options and return the pointer to the options and the pointer to the flagset.
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.token, "token", "t", "", "specify the token for the service. This option is mandatory.")
	flags.StringVarP(&opts.qrcode, "qrcode", "q", "", "include QR-code of the URL in the output.")
	flags.StringVarP(&opts.config, "config", "c", "", "specify the configuration file.")
	flags.BoolVarP(&opts.help, "help", "h", false, "print this mesasge and exit.")
	flags.BoolVarP(&opts.version, "version", "v", false, "print the version and exit.")
	return opts, flags
}

/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *UrleapError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.help {
		fmt.Println(helpMessage(args))
		return nil, nil, &UrleapError{statusCode: 0, message: ""}
	}
	if opts.token == "" {
		return nil, nil, &UrleapError{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}

func perform(opts *options, args []string) *UrleapError {
	fmt.Println("Hello World")
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
