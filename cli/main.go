package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"uppies/cli/commands"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "uppies",
		Short: "Uppies CLI tool",
	}

	commands.LoadConfig()

	rootCmd.AddCommand(commands.PlzCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
