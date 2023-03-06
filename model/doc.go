package model

import (
	"brag-cli/common"
	"time"
)

type Doc struct {
	ID        string `json:"id"`
	Brag      string `json:"brag"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func NewDoc(brag string) *Doc {
	return &Doc{
		ID:        common.GenID(brag + string(time.Now().Unix())),
		Brag:      brag,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
