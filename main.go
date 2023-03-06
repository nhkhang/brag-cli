package main

import (
	"brag-cli/cmd/bragcli"
)

func main() {
	// if err := config.LoadUserConfig(); err != nil {
	// 	log.Fatal(err)
	// }

	bragcli.Execute()
}
