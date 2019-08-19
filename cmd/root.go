package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	device string
	debug  bool
)

var rootCmd = &cobra.Command{
	Use:   "sfp",
	Short: "sfp is a sfp i2c reader/writer",
	Long:  "Small Form-factor Pluggable Inter-Integrated Circuit Bus Reader/Writer",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVar(&device, "device", "/dev/i2c-1", "path to i2c device (default is /dev/i2c-1)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debugging messages")
}

func printLog(text string) {
	if debug == true {
		fmt.Println(text)
	}
}
