package bragcli

import (
	"brag-cli/common"
	"brag-cli/service"
	"fmt"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use: common.EditCommand,
	Run: func(cmd *cobra.Command, args []string) {
		mess, err := cmd.Flags().GetString(string(common.MessFlag))
		if err != nil || mess == "" {
			return
		}

		cat, err := cmd.Flags().GetString(string(common.CategoryFlag))
		if err != nil {
			return
		}

		if cat == "" {
			cat = common.DefaultCategory
		}

		category, err := service.GetCategory(cat)
		if err != nil {
			fmt.Printf("Get category failed: %v", err)
			return
		}

		fmt.Println("Category: ", category)
	},
}
