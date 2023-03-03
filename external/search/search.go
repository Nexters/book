package search

import (
	"fmt"
	"strconv"

	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/nexters/book/config"
	"github.com/nexters/book/http/requests"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var display = 10

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
		SearchBook(query string, page string) (SearchResponse, error)
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
func (b bookSearch) SearchBook(query string, page string) (res SearchResponse, err error) {
	// make request
	requests := requests.NewHttp[SearchResponse](
		map[string]string{
			"X-Naver-Client-Id":     b.settings.External.ClientID,
			"X-Naver-Client-Secret": b.settings.External.ClientSecret,
		},
	)

	// page가 주어진 경우
	if len(page) > 0 {
		n, parse_err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			err = parse_err
			return
		}

		start := n + ((n - 1) * (int64(display)))
		uri := fmt.Sprintf(
			"%v?query=%v&display=%d&start=%d",
			b.settings.External.SearchEndpoint,
			url.QueryEscape(query),
			display,
			start,
		)

		res, err = requests.GET(uri)
		return
	}

	// LEGACY: page가 주어지지 않은 경우
	uri := fmt.Sprintf(
		"%v?query=%v",
		b.settings.External.SearchEndpoint,
		url.QueryEscape(query),
	)

	res, err = requests.GET(uri)
	return
}
