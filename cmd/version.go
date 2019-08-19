package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version string = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of sfp",
	Long:  "Print the version number of sfp",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Small Form-factor Pluggable Inter-Integrated Circuit Bus Reader/Writer v%s\n", version)
	},
}
