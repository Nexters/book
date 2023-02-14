package memo

type UpdateMemoPayload struct {
	Text     string `json:"text"`
	Category string `json:"category"`
}
