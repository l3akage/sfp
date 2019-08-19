package cmd

import (
	"fmt"
	"log"
	"os"

	sfp "github.com/l3akage/sfp/libsfp"
	"github.com/spf13/cobra"
)

var (
	file string
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump firmware",
	Long:  "Read firmware from SFP and write to file",
	Run: func(cmd *cobra.Command, args []string) {
		dumpFirmware()
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.PersistentFlags().StringVar(&file, "file", "firmware.bin", "Write firmware to filename")
}

func dumpFirmware() {
	printLog(fmt.Sprintf("Trying to open device %s", device))
	device, err := sfp.NewDevice(device)
	if err != nil {
		log.Fatal(err)
	}
	defer device.Close()

	printLog(fmt.Sprintf("Creating file %s", file))
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	printLog("Reading eeprom")
	data := device.Raw()
	n, err := f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	printLog(fmt.Sprintf("Finished. Wrote %d btye to %s", n, file))
}
