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

type UpdateBookPayload struct {
	IsReading *bool `json:"isReading" validate:"required,boolean"`
}

type UpdateMemoPayload struct {
	Text     string `json:"text"`
	Category string `json:"category"`
}

type UserStatPayload struct {
	Duration  int64 `json:"duration"`
	ReadCount int   `json:"readCount"`
	MemoCount int   `json:"memoCount"`
}
