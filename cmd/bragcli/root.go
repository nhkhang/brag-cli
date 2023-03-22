package bragcli

import (
	"fmt"
	"os"

	"brag-cli/common"

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

func Execute() {
	var mFlag string
	var cFlag string

	// Add commands
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(editCmd)

	// Add flags
	rootCmd.PersistentFlags().StringVarP(&mFlag, string(common.MessFlag), string(common.MessShortFlag), "", "Input brag message")
	rootCmd.PersistentFlags().StringVarP(&cFlag, string(common.CategoryFlag), string(common.CategoryShortFlag), "", "Input type of brag")

	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("Error execute root cmd: %v", err)
		os.Exit(0)
	}
}
