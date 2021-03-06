package model

// Todo represents a Todo in database
type Todo struct {
	ID     string `json:"id" pg:"type:bigserial,pk"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user_id"`
}
