package urleap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Bitly struct {
	url string
}

func NewBitly() *Bitly {
	return &Bitly{url: "https://api-ssl.bitly.com/v4/"}
}

func (bitly *Bitly) buildUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", bitly.url, endpoint)
}

func (bitly *Bitly) List(config *Config) []*ShortenUrl {
	return []*ShortenUrl{}
}

func (bitly *Bitly) Shorten(config *Config, url string) (*ShortenUrl, error) {
	reader := strings.NewReader(fmt.Sprintf(`{"long_url": "%s"}`, url))
	request, err := http.NewRequest("POST", bitly.buildUrl("shorten"), reader)
	if err != nil {
		return nil, err
	}
	data, err := sendRequest(request, config)
	if err != nil {
		return nil, err
	}
	result := &ShortenUrl{}
	err = json.Unmarshal(data, result)
	return result, err
}

func (bitly *Bitly) Delete(config *Config, shortenURL string) error {
	request, err := http.NewRequest("DELETE", bitly.buildUrl("bitlinks/"+strings.TrimPrefix(shortenURL, "https://")), nil)
	if err != nil {
		return err
	}
	_, err = sendRequest(request, config)
	return err
}

func (bitly *Bitly) QRCode(config *Config, shortenURL string) ([]byte, error) {
	return []byte{}, fmt.Errorf("not implement yet")
}

func sendRequest(request *http.Request, config *Config) ([]byte, error) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	if response.StatusCode/100 != 2 {
		return []byte{}, fmt.Errorf("response status code %d", response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
