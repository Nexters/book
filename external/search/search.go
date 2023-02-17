package search

import (
	"fmt"
	"io"

	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/nexters/book/config"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type SearchResponse struct {
	LastBuildDate string       `json:"lastBuildDate"`
	Total         int          `json:"total"`
	Start         int          `json:"start"`
	Display       int          `json:"display"`
	Items         []SearchItem `json:"items"`
}

type SearchItem struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ShopLink    string `json:"link"`
	Image       string `json:"image"`
	Price       string `json:"discount"`
	Publisher   string `json:"publisher"`
	ISBN        string `json:"isbn"`
	Pubdate     string `json:"pubdate"`
	Description string `json:"description"`
}

type (
	BookSearch interface {
		SearchBook(query string) (SearchResponse, error)
	}
	bookSearch struct {
		settings *config.Settings
	}
)

func NewBookSearch(s *config.Settings) BookSearch {
	return bookSearch{s}
}

func (b bookSearch) SearchBook(query string) (searchRes SearchResponse, err error) {
	url := fmt.Sprintf("%v?query=%v", b.settings.External.SearchEndpoint, url.QueryEscape(query))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Add("X-Naver-Client-Id", b.settings.External.ClientID)
	req.Header.Add("X-Naver-Client-Secret", b.settings.External.ClientSecret)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	result, err := io.ReadAll(res.Body)

	err = json.Unmarshal(result, &searchRes)
	if err != nil {
		return
	}

	res.Body.Close()

	return searchRes, nil
}
