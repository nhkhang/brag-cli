package config

import (
	"brag-cli/common"
	"brag-cli/model"
)

var UserConfig *model.UserConfig
var UserConfigName string = "user-config.json"

func LoadUserConfig() error {
	if err := common.LoadData(UserConfigName, UserConfig); err != nil {
		return err
	}

	return nil
}
