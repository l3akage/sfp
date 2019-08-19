package sfp

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type DDM struct {
	Temperature float64
	Vcc         float64
	TxBias      float64
	OpticalTx   float64
	OpticalRx   float64
}

func (d *DDM) Print() {
	d.printTable()
}

func (d *DDM) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColMinWidth(0, 50)
	table.SetColMinWidth(1, 80)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Name", "Value"})

	table.Append([]string{"Internal SFP Temperature", fmt.Sprintf("%4.2fC", d.Temperature)})
	table.Append([]string{"Internal supply voltage", fmt.Sprintf("%4.2fV", d.Vcc)})
	table.Append([]string{"TX bias current", fmt.Sprintf("%4.2fmA", d.TxBias)})
	table.Append([]string{"Optical power Tx", fmt.Sprintf("%4.2f dBm", d.OpticalTx)})
	table.Append([]string{"Optical power Rx", fmt.Sprintf("%4.2f dBm", d.OpticalRx)})

	table.Render()
}
