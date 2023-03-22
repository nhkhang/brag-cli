package model

import (
	"brag-cli/common"
	"time"
)

type Category struct {
	ID        string    `json:"id" csv:"id"`
	Name      string    `json:"name" csv:"name"`
	CreatedAt time.Time `json:"createdAt" csv:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" csv:"updatedAt"`
	Brags     []Brag    `json:"brags" csv:"brags"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID:        common.GenerateRandomID(6),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Brags:     make([]Brag, 0),
	}
}
