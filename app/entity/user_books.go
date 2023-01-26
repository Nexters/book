package entity

import "time"

type UserBooks struct {
	UserID    uint      `gorm:"primaryKey" json:"userId"`
	BookID    uint      `gorm:"primaryKey" json:"bookId"`
	StartedAt time.Time `gorm:"column:started_at" json:"startedAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
	DoneAt    time.Time `gorm:"column:done_at" json:"doneAt"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"`
	// Memos     []Memo    `gorm:"memos" json:"memos"`
	Done bool `gorm:"done;default:false" json:"done"`
}
