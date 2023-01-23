package entity

import "time"

type UserBooks struct {
	UserID    uint64    `gorm:"primaryKey" json:"userId"`
	BookID    uint64    `gorm:"primaryKey" json:"bookId"`
	StartedAt time.Time `gorm:"started_at" json:"startedAt"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updatedAt"`
	DoneAt    time.Time `gorm:"done_at" json:"doneAt"`
	DeletedAt time.Time `gorm:"deleted_at" json:"deletedAt"`
	// Memos     []Memo    `gorm:"memos" json:"memos"`
	Done bool `gorm:"done" json:"done"`
}
