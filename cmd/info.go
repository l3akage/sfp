package cmd

import (
	"fmt"
	"log"

	sfp "github.com/l3akage/sfp/libsfp"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Read SFP",
	Long:  "Read data from SFP and print them out",
	Run: func(cmd *cobra.Command, args []string) {
		readFirmware()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func readFirmware() {
	printLog(fmt.Sprintf("Trying to open device %s", device))
	device, err := sfp.NewDevice(device)
	if err != nil {
		log.Fatal(err)
	}
	defer device.Close()

	printLog("Read Data")
	t := device.GetTransceiver()
	t.Print()

	if device.HasDDM() {
		printLog("Read DDM")
		d, _ := device.GetDDM()
		d.Print()
	} else {
		printLog("No DDM available")
	}
}
