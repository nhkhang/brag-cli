package bragcli

import (
	"brag-cli/common"
	"brag-cli/model"
	"brag-cli/service"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func PrintBrags(input []model.Brag) {
	data := [][]string{
		{"ID", "TITLE", "CATEGORY", "CREATED_AT"},
	}

	for _, brag := range input {
		data = append(data, []string{brag.ID, brag.Title, brag.Category, brag.CreatedAt.String()})
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, ' ', 0)
	for _, row := range data {
		fmt.Fprintln(w, row[0]+"\t"+row[1]+"\t"+row[2]+"\t"+row[3])
	}
}

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

		category, err := service.GetCategory(cat)
		if err != nil {
			fmt.Printf("Get category failed: %v", err)
			return
		}

		fmt.Println("Category: ", category)
		PrintBrags(category.Brags)
	},
}
