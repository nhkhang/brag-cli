package bragcli

import (
	"fmt"
	"os"

	"brag-cli/common"
	"brag-cli/pkg/category"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brag",
	Short: "bragcli - simple way to brag and get performance review easier",
	Long:  "bragcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to brag-cli")
	},
}

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

var viewCmd = &cobra.Command{
	Use: common.ViewCommand,
	Run: func(cmd *cobra.Command, args []string) {
		// message, _ := cmd.Flags().GetString("m")
		category, err := category.GetCategory("default")
		if err != nil {
			fmt.Printf("Get category failed: %v", err)
			return
		}

		fmt.Println("Category: ", category)
	},
}

func Execute() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(viewCmd)

	rootCmd.PersistentFlags().StringVarP(nil, string(common.MessFlag), string(common.MessShortFlag), "", "Input brag message")
	rootCmd.PersistentFlags().StringVarP(nil, string(common.CategoryFlag), string(common.CategoryShortFlag), "", "Input type of brag")

	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("Error execute root cmd: %v", err)
		os.Exit(0)
	}
}
