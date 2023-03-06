package model

type UserConfig struct {
	CurCategory string `json:"curCategory"`
}

func NewUserConfig() *UserConfig {
	return &UserConfig{
		CurCategory: "default",
	}
}
