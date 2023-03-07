package bragcli

import (
	"fmt"
	"os"

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
	Use: "add",
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("m")

		fmt.Println("Add brag: " + message)
		if err := category.AddBrag(message, "default"); err != nil {
			fmt.Printf("Add brag failed: %v", err)
		}
	},
}

var getCmd = &cobra.Command{
	Use: "get",
	// Args: cobra.MinimumNArgs(1),
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

	addCmd.PersistentFlags().String("m", "", "String message")

	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("Error execute root cmd: %v", err)
		os.Exit(0)
	}
}
