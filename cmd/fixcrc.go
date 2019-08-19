package cmd

import (
	"fmt"
	"log"

	sfp "github.com/l3akage/sfp/libsfp"
	"github.com/spf13/cobra"
)

// fixCRCCmd represents the fixcrc command
var fixCRCCmd = &cobra.Command{
	Use:   "fixcrc",
	Short: "Fix Checksum",
	Long:  "Fix base and extended checksum",
	Run: func(cmd *cobra.Command, args []string) {
		fixCRC()
	},
}

func init() {
	rootCmd.AddCommand(fixCRCCmd)
}

func fixCRC() {
	printLog(fmt.Sprintf("Trying to open device %s", device))
	device, err := sfp.NewDevice(device)
	if err != nil {
		log.Fatal(err)
	}
	defer device.Close()

	printLog("Checking base CRC")
	baseValid := device.HasValidBaseChecksum()
	if baseValid == false {
		printLog("Fixing invalid base checksum")
		device.FixBaseChecksum()
	} else {
		printLog("Base checksum is valid")
	}
	printLog("Checking extended CRC")
	extendedValid := device.HasValidExtendedChecksum()
	if extendedValid == false {
		printLog("Fixing invalid extended checksum")
		device.FixExtendedChecksum()
	} else {
		printLog("Extended checksum is valid")
	}
}
