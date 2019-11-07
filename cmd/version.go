package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Run: func(cmd *cobra.Command, args []string) {
		version := `
	Mail - Cli utility
	Version: 1.0.0
	Built:   06/11/2019`
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
