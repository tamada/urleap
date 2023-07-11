package urleap

import (
	"os"
	"testing"
)

func TestShortenUrl(t *testing.T) {
	config := NewConfig(os.Getenv("URLEAP_TOKEN"), Shorten)
	bitly := NewBitly("")
	testdata := []struct {
		giveUrl          string
		wontShortenError bool
		wontDeleteError  bool
	}{
		{"https://tamadalab.github.io/", false, false},
	}
	for _, td := range testdata {
		result, err := bitly.Shorten(config, td.giveUrl)
		if (err == nil) == td.wontShortenError {
			t.Errorf("shorten %s wont error %t, but got %t", td.giveUrl, td.wontShortenError, !td.wontShortenError)
		}
		err = bitly.Delete(config, result.Shorten)
		if (err == nil) == td.wontDeleteError {
			t.Errorf("delete %s wont error %t, but got %t", result.Shorten, td.wontDeleteError, !td.wontDeleteError)
		}
	}
}
