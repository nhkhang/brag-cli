package main

import (
	"brag-cli/cmd/bragcli"
	"brag-cli/common"
	"brag-cli/model"
	"brag-cli/service"
	"fmt"
	"os"
)

func CreateDefaultCategories() {
	for i := 0; i < len(common.SystemListCategories); i++ {
		categoryName := common.SystemListCategories[i]
		filePath := common.GetDataFilePath(categoryName)

		if _, err := os.Open(filePath); err != nil {
			if os.IsNotExist(err) {
				_, err := os.Create(filePath)
				if err != nil {
					return
				}

				if err := service.SaveCategory(model.NewCategory(categoryName)); err != nil {
					fmt.Printf("Error saving category %s: %s\n", categoryName, err)
				}
			}
		}
	}
}

func main() {
	// Setup default categories
	CreateDefaultCategories()

	// root cmd execution
	bragcli.Execute()
}
