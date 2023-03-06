package category

import (
	"brag-cli/common"
	"brag-cli/model"
	"errors"
	"fmt"
)

var CurCategory *model.Category

func LoadCategory(category string) error {
	if err := common.LoadData(category, &CurCategory); err != nil {
		return err
	}

	if CurCategory == nil {
		return errors.New("load category failed")
	}

	return nil
}

func SaveCategory(category string) error {
	if err := common.SaveData(category, CurCategory); err != nil {
		return err
	}

	return nil
}

func AddBrag(brag string, category string) error {
	if err := LoadCategory(category); err != nil {
		fmt.Println("vl")
		return err
	}

	fmt.Println("cur category", CurCategory)

	CurCategory.Docs = append(CurCategory.Docs, *model.NewDoc(brag))

	fmt.Println("cur category", CurCategory)
	if err := SaveCategory(category); err != nil {
		return err
	}

	return nil
}
