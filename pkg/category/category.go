package category

import (
	"brag-cli/common"
	"brag-cli/model"
	"errors"
	"fmt"
)

func LoadCategory(category string) (*model.Category, error) {
	var data *model.Category
	if err := common.LoadData(category, &data); err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("load category failed")
	}

	return data, nil
}

func SaveCategory(category *model.Category) error {
	if err := common.SaveData(category.Name, category); err != nil {
		return err
	}

	return nil
}

func AddBrag(brag string, category string) error {
	dCategory, err := LoadCategory(category)
	if err != nil {
		return err
	}

	fmt.Println("cur category", dCategory)

	dCategory.Brags = append(dCategory.Brags, *model.NewBrag(brag))

	if err := SaveCategory(dCategory); err != nil {
		return err
	}

	return nil
}

func GetCategory(category string) (*model.Category, error) {
	dCategory, err := LoadCategory(category)
	if err != nil {
		return nil, err
	}

	return dCategory, nil
}
