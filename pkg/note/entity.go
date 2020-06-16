package note

import "github.com/jinzhu/gorm"

// Note export
type Note struct {
	gorm.Model
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
