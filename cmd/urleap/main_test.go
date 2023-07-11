package main

import "testing"

func Example_Help() {
	goMain([]string{"./urleap", "--help"})
	// Output:
	// urleap [OPTIONS] [URLs...]
	// OPTIONS
	//   -c, --config string   specify the configuration file.
	//   -d, --delete          delete the specified shorten URL.
	//   -g, --group string    specify the group name for the service.
	//   -h, --help            print this mesasge and exit.
	//   -L, --list-group      list the groups. This is hidden option.
	//   -q, --qrcode string   include QR-code of the URL in the output.
	//   -t, --token string    specify the token for the service. This option is mandatory.
	//   -v, --version         print the version and exit.
	//
	// ARGUMENT
	//     URL     specify the url for shortening. this arguments accept multiple values.
	//             if no arguments were specified, urleap prints the list of available shorten urls.
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./urleap", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
