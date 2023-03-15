package bragcli

import (
	"brag-cli/common"
	"brag-cli/pkg/category"
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: common.AddCommand,
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

		if err := category.AddBrag(mess, cat); err != nil {
			fmt.Printf("Add brag failed: %v \n", err)
			return
		}
	},
}
