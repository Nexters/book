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
	ID        uint `json:"bookId" validate:"required,number"`
	IsReading bool `json:"isReading" validate:"required,boolean"`
}

type UpdateMemoPayload struct {
	ID       uint   `json:"memoId" validate:"required,number"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

type UserStatPayload struct {
	Duration  int64 `json:"duration"`
	ReadCount int   `json:"readCount"`
	MemoCount int   `json:"memoCount"`
}
