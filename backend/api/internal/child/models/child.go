package models

import (
	"time"
)

type Child struct {
	ID        string    `json:"id"`
	ParentID  string    `json:"parent_id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
