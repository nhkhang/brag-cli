package model

import (
	"brag-cli/common"
	"time"
)

type Brag struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Description *string   `json:"description,omitempty"`
	RefID       *string   `json:"refId,omitempty"`
}

func NewBrag(title string) *Brag {
	return &Brag{
		ID:        common.GenerateRandomID(6),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     title,
	}
}
