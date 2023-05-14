package urleap

import "fmt"

type ShortenUrl struct {
	Shorten   string `json:"link"`
	Original  string `json:"long_url"`
	IsDeleted bool   `json:"is_deleted"`
	Group     string
}

func (surl *ShortenUrl) String() string {
	return fmt.Sprintf("%s: %s (%s)", surl.Original, surl.Shorten, surl.Group)
}

type URLShortener interface {
	List(config *Config) ([]*ShortenUrl, error)
	Shorten(config *Config, url string) (*ShortenUrl, error)
	Delete(config *Config, shortenURL string) error
	QRCode(config *Config, shortenURL string) ([]byte, error)
}
