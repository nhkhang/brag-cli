package bragcli

import (
	"brag-cli/common"
	"brag-cli/pkg/category"
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use: common.ViewCommand,
	Run: func(cmd *cobra.Command, args []string) {
		cat, err := cmd.Flags().GetString(string(common.CategoryFlag))
		if err != nil {
			return
		}

		if cat == "" {
			cat = common.DefaultCategory
		}

		category, err := category.GetCategory(cat)
		if err != nil {
			fmt.Printf("Get category failed: %v", err)
			return
		}

		fmt.Println("Category: ", category)
	},
}
