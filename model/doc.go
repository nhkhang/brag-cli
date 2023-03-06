package model

type Doc struct {
	ID        string `json:"id"`
	Brag      string `json:"brag"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
