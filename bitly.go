package urleap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Bitly struct {
	url   string
	group string
}

type group struct {
	guid     string `json:"guid"`
	isActive bool   `json:"is_active"`
}

func (bitly *Bitly) groups(config *Config) (*group, error) {
	request, err := http.NewRequest("GET", bitly.buildUrl("groups"), nil)
	if err != nil {
		return nil, err
	}
	data, err := sendRequest(request, config)
	if err != nil {
		return nil, err
	}
	result := []*group{}
	err = json.Unmarshal(data, &result)
	for _, g := range result {
		if g.isActive {
			return g, err
		}

	}
	return nil, fmt.Errorf("no active group found")
}

func NewBitly(group string) *Bitly {
	return &Bitly{url: "https://api-ssl.bitly.com/v4/", group: group}
}

func (bitly *Bitly) buildUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", bitly.url, endpoint)
}

func (bitly *Bitly) List(config *Config) ([]*ShortenUrl, error) {
	if bitly.group == "" {
		gs, err := bitly.groups(config)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("list subcommand requires a group guid: %s", gs.guid)
	}
	request, err := http.NewRequest("GET", bitly.buildUrl(fmt.Sprintf("/groups/%s/bitlinks?size=20", bitly.group)), nil)
	if err != nil {
		return nil, err
	}
	data, err := sendRequest(request, config)
	if err != nil {
		return nil, err
	}
	result := struct {
		Links []*ShortenUrl `json:"links"`
	}{}
	err = json.Unmarshal(data, &result)
	return removeDeletedLinks(result.Links, bitly.group), err
}

func removeDeletedLinks(links []*ShortenUrl, group string) []*ShortenUrl {
	result := []*ShortenUrl{}
	for _, link := range links {
		if !link.IsDeleted {
			link.Group = group
			result = append(result, link)
		}
	}
	return result
}

func (bitly *Bitly) Shorten(config *Config, url string) (*ShortenUrl, error) {
	reader := strings.NewReader(fmt.Sprintf(`{"long_url": "%s", "group_guid": "%s"}`, url, bitly.group))
	request, err := http.NewRequest("POST", bitly.buildUrl("shorten"), reader)
	if err != nil {
		return nil, err
	}
	data, err := sendRequest(request, config)
	if err != nil {
		return nil, err
	}
	result := &ShortenUrl{}
	fmt.Println("result:", string(data))
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	result.Group, err = findGroup(data)
	return result, err
}

func findGroup(data []byte) (string, error) {
	r := struct {
		References struct {
			Group string `json:"group"`
		} `json:"references"`
	}{}
	err := json.Unmarshal(data, &r)
	uri := r.References.Group
	index := strings.LastIndex(uri, "/")
	return uri[index+1:], err
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
		data, _ := io.ReadAll(response.Body)
		fmt.Println("response body:", string(data))
		return []byte{}, fmt.Errorf("response status code %d", response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
