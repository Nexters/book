package search

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/nexters/book/app/config"
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

func (b bookSearch) SearchBook(query string) (SearchResponse, error) {
	url := fmt.Sprintf("%v?query=%v", b.settings.External.SearchEndpoint, url.QueryEscape(query))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Naver-Client-Id", b.settings.External.ClientID)
	req.Header.Add("X-Naver-Client-Secret", b.settings.External.ClientSecret)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	result, err := io.ReadAll(res.Body)

	searchResponse := SearchResponse{}

	err = json.Unmarshal(result, &searchResponse)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	return searchResponse, nil
}
