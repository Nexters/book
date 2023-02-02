package payloads

import "github.com/nexters/book/app/entity"

type FindAllBooksPayload struct {
	Count int               `json:"count"`
	Books []FindBookPayload `json:"books"`
}

type FindBookPayload struct {
	entity.Book
	MemoCount int `json:"memoCount"`
}
