package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	"github.com/tamada/urleap"
)

const VERSION = "0.2.2"

func versionString(args []string) string {
	prog := "urleap"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

/*
helpMessage prints the help message.
This function is used in the small tests, so it may be called with a zero-length slice.
*/
func helpMessage(args []string, flags *flag.FlagSet) string {
	prog := "urleap"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [URLs...]
OPTIONS
%s
ARGUMENT
    URL     specify the url for shortening. this arguments accept multiple values.
            if no arguments were specified, urleap prints the list of available shorten urls.`, prog, flags.FlagUsages())
}

type UrleapError struct {
	statusCode int
	message    string
}

func (e UrleapError) Error() string {
	return e.message
}

type flags struct {
	deleteFlag    bool
	listGroupFlag bool
	helpFlag      bool
	versionFlag   bool
}

type runOpts struct {
	token  string
	qrcode string
	config string
	group  string
}

/*
This struct holds the values of the options.
*/
type options struct {
	runOpt  *runOpts
	flagSet *flags
}

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

func (opts *options) mode(args []string) urleap.Mode {
	switch {
	case opts.flagSet.listGroupFlag:
		return urleap.ListGroup
	case len(args) == 0:
		return urleap.List
	case opts.flagSet.deleteFlag:
		return urleap.Delete
	case opts.runOpt.qrcode != "":
		return urleap.QRCode
	default:
		return urleap.Shorten
	}
}

/*
Define the options and return the pointer to the options and the pointer to the flagset.
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	completions := false
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args, flags)) }
	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "specify the token for the service. This option is mandatory.")
	flags.StringVarP(&opts.runOpt.qrcode, "qrcode", "q", "", "include QR-code of the URL in the output.")
	flags.StringVarP(&opts.runOpt.config, "config", "c", "", "specify the configuration file.")
	flags.StringVarP(&opts.runOpt.group, "group", "g", "", "specify the group name for the service.")
	flags.BoolVarP(&opts.flagSet.listGroupFlag, "list-group", "L", false, "list the groups. This is hidden option.")
	flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "delete the specified shorten URL.")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "print this mesasge and exit.")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "print the version and exit.")
	flags.BoolVarP(&completions, "generate-completions", "", false, "generate completions")
	flags.MarkHidden("generate-completions")
	return opts, flags
}

/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *UrleapError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args, flags))
		return nil, nil, &UrleapError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &UrleapError{statusCode: 0, message: ""}
	}
	if value, _ := flags.GetBool("generate-completions"); value {
		err := GenerateCompletion(flags)
		if err != nil {
			return nil, nil, &UrleapError{statusCode: 1, message: err.Error()}
		}
		return nil, nil, &UrleapError{statusCode: 0, message: "generate completions"}
	}
	if opts.runOpt.token == "" {
		return nil, nil, &UrleapError{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}

func printResult(result *urleap.ShortenUrl) {
	fmt.Printf("%s (%s)\n", result.Shorten, result.Group)
}

func shortenEach(bitly *urleap.Bitly, config *urleap.Config, url string) error {
	result, err := bitly.Shorten(config, url)
	if err != nil {
		return err
	}
	printResult(result)
	return nil
}

func deleteEach(bitly *urleap.Bitly, config *urleap.Config, url string) error {
	return bitly.Delete(config, url)
}

func listUrls(bitly *urleap.Bitly, config *urleap.Config) error {
	urls, err := bitly.List(config)
	if err != nil {
		return err
	}
	for _, url := range urls {
		fmt.Println(url)
	}
	return nil
}

func listGroups(bitly *urleap.Bitly, config *urleap.Config) error {
	groups, err := bitly.Groups(config)
	if err != nil {
		return err
	}
	for i, group := range groups {
		fmt.Printf("GUID[%d] %s\n", i, group.Guid)
	}
	return nil
}

func performImpl(args []string, executor func(url string) error) *UrleapError {
	for _, url := range args {
		err := executor(url)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

func perform(opts *options, args []string) *UrleapError {
	bitly := urleap.NewBitly(opts.runOpt.group)
	config := urleap.NewConfig(opts.runOpt.token, opts.mode(args))
	switch config.RunMode {
	case urleap.List:
		err := listUrls(bitly, config)
		return makeError(err, 1)
	case urleap.ListGroup:
		err := listGroups(bitly, config)
		return makeError(err, 2)
	case urleap.Delete:
		return performImpl(args, func(url string) error {
			return deleteEach(bitly, config, url)
		})
	case urleap.Shorten:
		return performImpl(args, func(url string) error {
			return shortenEach(bitly, config, url)
		})
	}
	return nil
}

func makeError(err error, status int) *UrleapError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*UrleapError)
	if ok {
		return ue
	}
	return &UrleapError{statusCode: status, message: err.Error()}
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
