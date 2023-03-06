package model

import (
	"brag-cli/common"
	"time"
)

type Category struct {
	ID        string `json:"id" csv:"id"`
	Name      string `json:"name" csv:"name"`
	CreatedAt int64  `json:"createdAt" csv:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" csv:"updatedAt"`
	Docs      []Doc  `json:"docs" csv:"docs"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID:        common.GenID(name + time.Now().String()),
		Name:      name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Docs:      make([]Doc, 0),
	}
}
