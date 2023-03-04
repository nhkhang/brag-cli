package bragcli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bragcli",
	Short: "bragcli - simplest way to brag and performance review easier",
	Long:  "bragcli",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("Error execute root cmd: %v", err)
		os.Exit(0)
	}
}
