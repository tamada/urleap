package urleap

import "fmt"

type ShortenUrl struct {
	Shorten  string `json:"link"`
	Original string `json:"long_url"`
}

func (surl *ShortenUrl) String() string {
	return fmt.Sprintf("%s: %s", surl.Original, surl.Shorten)
}

type URLShortener interface {
	List(config *Config) []*ShortenUrl
	Shorten(config *Config, url string) (*ShortenUrl, error)
	Delete(config *Config, shortenURL string) error
	QRCode(config *Config, shortenURL string) ([]byte, error)
}
