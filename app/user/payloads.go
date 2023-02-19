package user

type UserStatPayload struct {
	Duration  int64 `json:"duration"`
	ReadCount int   `json:"readCount"`
	MemoCount int   `json:"memoCount"`
}
