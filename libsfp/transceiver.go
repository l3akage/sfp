package sfp

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Transceiver struct {
	ModuleIdentifier byte
	Connector        byte
	Compliance       []byte
	Encoding         byte
	BaudRatex100MBd  byte
	RateIdentifier   byte
	Length           []byte
	Wavelength       []byte
	DWDM             byte
	Vendor           Vendor
	ValidBaseCRC     bool

	Options          []byte
	BaudRateMax      byte
	BaudRateMin      byte
	VendorSN         []byte
	VendorDate       []byte
	DDMOptions       byte
	ExtendedOptions  byte
	SFF8472          byte
	ValidExtendedCRC bool
}

type Vendor struct {
	Name string
	OUI  []byte
	PN   string
	Rev  string
}

func convert(b byte) string {
	return strconv.Itoa(int(b))
}

func (t *Transceiver) Print() {
	t.printTable()
}

func (t *Transceiver) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColMinWidth(0, 50)
	table.SetColMinWidth(1, 80)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Name", "Value"})

	table.Append([]string{"Module", ModuleIdentifier[uint64(rune(t.ModuleIdentifier))]})
	table.Append([]string{"Connector", ModuleConnector[uint64(rune(t.Connector))]})

	var comp uint64
	for i, x := range t.Compliance {
		comp |= uint64(rune(x)) << uint64(i*8)
	}
	for i, name := range Compliance {
		if (comp & i) == 0 {
			continue
		}
		table.Append([]string{"Compliance", name})
	}
	table.Append([]string{"Encoding", Encoding[uint64(rune(t.Encoding))]})
	table.Append([]string{"Nominal bit rate, units of 100 MBits/sec.", convert(t.BaudRatex100MBd)})
	table.Append([]string{"Rate identifier", RateIdentifier[uint64(rune(t.RateIdentifier))]})
	table.Append([]string{"Link length supported for 9/125 µm fiber, units of km", convert(t.Length[0])})
	table.Append([]string{"Link length supported for 9/125 µm fiber, units of 100 m", convert(t.Length[1])})
	table.Append([]string{"Link length supported for 50/125 µm OM2 fiber, units of 10 m", convert(t.Length[2])})
	table.Append([]string{"Link length supported for 62.5/125 µm OM1 fiber, units of 10 m", convert(t.Length[3])})
	table.Append([]string{"Link length supported for copper and Active Cable, units of meters", convert(t.Length[4])})
	table.Append([]string{"Link length supported for 50/125 µm fiber, units of 10 m", convert(t.Length[5])})
	table.Append([]string{"Wavelength", fmt.Sprintf("%d.%d", (int64(t.Wavelength[0])<<8)|int64(t.Wavelength[1]), t.DWDM)})
	table.Append([]string{"Vendor name", t.Vendor.Name})
	table.Append([]string{"Vendor OUI", fmt.Sprintf("0x%x 0x%x 0x%x", t.Vendor.OUI[0], t.Vendor.OUI[1], t.Vendor.OUI[2])})
	table.Append([]string{"Vendor PN", t.Vendor.PN})
	table.Append([]string{"Vendor Rev", t.Vendor.Rev})

	value := "False"
	if t.ValidBaseCRC {
		value = "True"
	}
	table.Append([]string{"Valid Base CRC", value})

	var opts uint64
	for i, x := range t.Options {
		opts |= uint64(rune(x)) << uint64(i*8)
	}
	for i, name := range Options {
		if (opts & i) == 0 {
			continue
		}
		table.Append([]string{"Option", name})
	}

	table.Append([]string{"Baud rate max", fmt.Sprintf("%d%%", t.BaudRateMax)})
	table.Append([]string{"Baud rate min", fmt.Sprintf("%d%%", t.BaudRateMin)})
	table.Append([]string{"Vendor s/n", string(t.VendorSN)})
	table.Append([]string{"Vendor manufacturing date", fmt.Sprintf("%s.%s.20%s", string(t.VendorDate[4:6]), string(t.VendorDate[2:4]), string(t.VendorDate[0:2]))})

	ddmopts := uint64(t.DDMOptions)
	for i, name := range DDMTypes {
		if (ddmopts & i) == 0 {
			continue
		}
		table.Append([]string{"DDM Option", name})
	}

	eopts := uint64(t.ExtendedOptions)
	for i, name := range EnhancedOptions {
		if (eopts & i) == 0 {
			continue
		}
		table.Append([]string{"Extended Option", name})
	}
	table.Append([]string{"SFF8472 compatibility:", SFF8472[uint64(rune(t.SFF8472))]})
	value = "False"
	if t.ValidExtendedCRC {
		value = "True"
	}
	table.Append([]string{"Valid Extended CRC", value})
	table.Render()
}
