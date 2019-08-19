package cmd

import (
	"encoding/binary"
	"fmt"
	"log"

	sfp "github.com/l3akage/sfp/libsfp"
	"github.com/spf13/cobra"
)

var (
	dictionary string
	bruteForce bool
	startWith  uint32
)

// crackCmd represents the crack command
var crackCmd = &cobra.Command{
	Use:   "crack",
	Short: "Crack SFP password (NEEDS TESTING)",
	Long:  "Crack password of a write protected SFP (NEEDS TESTING)",
	Run: func(cmd *cobra.Command, args []string) {
		crackPassword()
	},
}

func init() {
	rootCmd.AddCommand(crackCmd)
	crackCmd.PersistentFlags().Uint32Var(&startWith, "start-with", 0, "Star value for brute force (default 0)")
	crackCmd.PersistentFlags().StringVar(&file, "dictionary", "dictionary.txt", "Dictionary containing passwords to try first")
	crackCmd.PersistentFlags().BoolVar(&bruteForce, "no-brute-force", false, "Don't brute force")
}

func crackPassword() {
	printLog(fmt.Sprintf("Trying to open device %s", device))
	device, err := sfp.NewDevice(device)
	if err != nil {
		log.Fatal(err)
	}
	defer device.Close()

	if device.IsProtected() {
		fmt.Println("Device has password")
	} else {
		fmt.Println("Device has no password")
		return
	}

	current := make([]byte, 4)
	loop := startWith
	for {
		if loop > 4294967295 {
			break
		}
		binary.LittleEndian.PutUint32(current, uint32(loop))
		for i := len(current); i < 5; i++ {
			current = append(current, 0x00)
		}
		fmt.Printf("Testing 0x%x 0x%x 0x%x 0x%x Percent %f\n", current[0], current[1], current[2], current[3], float32(loop)/4294967295.0*100.0)
		device.SetPassword(current)
		if !device.IsProtected() {
			break
		}
		loop++
	}
	fmt.Printf("Password is 0x%x 0x%x 0x%x 0x%x\n", current[0], current[1], current[2], current[3])
}
