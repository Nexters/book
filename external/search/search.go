package search

import (
	"fmt"

	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/nexters/book/config"
	"github.com/nexters/book/http/requests"
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
	// BookSearch 외부 API를 이용한 책 검색
	BookSearch interface {
		SearchBook(query string) (SearchResponse, error)
	}
	bookSearch struct {
		settings *config.Settings
	}
)

// NewBookSearch 생성자
func NewBookSearch(s *config.Settings) BookSearch {
	return bookSearch{s}
}

// SearchBook search book description via external api
func (b bookSearch) SearchBook(query string) (res SearchResponse, err error) {
	url := fmt.Sprintf("%v?query=%v", b.settings.External.SearchEndpoint, url.QueryEscape(query))

	requests := requests.NewHttp[SearchResponse](
		map[string]string{
			"X-Naver-Client-Id":     b.settings.External.ClientID,
			"X-Naver-Client-Secret": b.settings.External.ClientSecret,
		},
	)

	// make request
	res, err = requests.GET(url)

	return
}
