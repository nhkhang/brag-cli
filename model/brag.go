package model

import (
	"brag-cli/common"
	"time"
)

type Brag struct {
	ID          string  `json:"id"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	RefID       *string `json:"refId,omitempty"`
}

func NewBrag(title string) *Brag {
	return &Brag{
		ID:        common.GenerateRandomID(6),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Title:     title,
	}
}
