package sfp

import (
	"math"
	"time"

	"golang.org/x/exp/io/i2c"
)

type Device struct {
	device      *i2c.Device
	ddm         *i2c.Device
	data        []byte
	transceiver Transceiver
}

func NewDevice(devisePath string) (*Device, error) {
	d := &Device{}
	con, err := i2c.Open(&i2c.Devfs{Dev: devisePath}, 0x50)
	if err != nil {
		return d, err
	}
	d.device = con
	err = d.readAll()
	if err != nil {
		d.Close()
		return &Device{}, err
	}
	d.newTransceiver()
	if d.HasDDM() {
		ddmCon, err := i2c.Open(&i2c.Devfs{Dev: devisePath}, 0x51)
		if err != nil {
			return d, err
		}
		d.ddm = ddmCon
	}
	return d, nil
}

func (d *Device) Close() {
	d.device.Close()
	if d.HasDDM() {
		d.ddm.Close()
	}
}

func (d *Device) Write(address byte, data []byte) error {
	err := d.device.WriteReg(address, data)
	if err != nil {
		return err
	}
	time.Sleep(200 * time.Millisecond)
	return nil
}

func (d *Device) Read(address byte, data *[]byte) error {
	err := d.device.ReadReg(address, *data)
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) readAll() error {
	data := make([]byte, 256)
	err := d.device.ReadReg(0, data)
	if err != nil {
		return err
	}
	d.data = data
	return nil
}

func (d *Device) generateChecksum(from, to int) int {
	crc := 0
	for _, b := range d.data[from:to] {
		crc += int(b)
		crc &= 255
	}
	return crc
}

func (d *Device) FixBaseChecksum() {
	baseCrc := d.generateChecksum(0, 63)
	d.Write(63, []byte{byte(baseCrc)})
}

func (d *Device) FixExtendedChecksum() {
	extendedCrc := d.generateChecksum(64, 95)
	d.Write(95, []byte{byte(extendedCrc)})
}

func (d *Device) Raw() []byte {
	return d.data
}

func (d *Device) HasValidBaseChecksum() bool {
	return int(d.data[63]) == d.generateChecksum(0, 63)
}

func (d *Device) HasValidExtendedChecksum() bool {
	return int(d.data[95]) == d.generateChecksum(64, 95)
}

func (d *Device) GetTransceiver() Transceiver {
	return d.transceiver
}

func (d *Device) HasDDM() bool {
	return d.transceiver.DDMOptions&0x60 == 0x60
}

func (d *Device) newTransceiver() {
	t := Transceiver{
		ModuleIdentifier: d.data[0],
		Connector:        d.data[2],
		Compliance:       d.data[3:11],
		Encoding:         d.data[11],
		BaudRatex100MBd:  d.data[12],
		RateIdentifier:   d.data[13],
		Length:           d.data[14:20],
		Wavelength:       d.data[60:62],
		DWDM:             d.data[62],
		Options:          d.data[64:66],
		BaudRateMax:      d.data[66],
		BaudRateMin:      d.data[67],
		VendorSN:         d.data[68:84],
		VendorDate:       d.data[84:90],
		DDMOptions:       d.data[92],
		ExtendedOptions:  d.data[93],
		SFF8472:          d.data[94],
	}
	t.Vendor = Vendor{
		Name: string(d.data[2:36]),
		OUI:  []byte{d.data[37], d.data[38], d.data[39]},
		PN:   string(d.data[40:56]),
		Rev:  string(d.data[56:60]),
	}

	t.ValidBaseCRC = d.HasValidBaseChecksum()
	t.ValidExtendedCRC = d.HasValidExtendedChecksum()
	d.transceiver = t
}

func (d *Device) GetDDM() (DDM, error) {
	data := make([]byte, 10)
	err := d.ddm.ReadReg(96, data)
	if err != nil {
		return DDM{}, err
	}
	return DDM{
		Temperature: float64(data[0]) + float64(data[1])/256.0,
		Vcc:         float64(int64(data[2])<<uint(8)|int64(data[3])) * 0.0001,
		TxBias:      float64(int64(data[4])<<uint(8)|int64(data[5])) * 0.002,
		OpticalTx:   10 * math.Log10(float64(int64(data[6])<<uint(8)|int64(data[7]))*0.0001),
		OpticalRx:   10 * math.Log10(float64(int64(data[8])<<uint(8)|int64(data[9]))*0.0001),
	}, nil
}

func (d *Device) SetPassword(password []byte) error {
	err := d.WriteDDM(0x7b, password)
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) ReadDDM(address byte, data []byte) error {
	err := d.ddm.ReadReg(address, data)
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) WriteDDM(address byte, data []byte) error {
	err := d.ddm.WriteReg(address, data)
	if err != nil {
		return err
	}
	time.Sleep(200 * time.Millisecond)
	return nil
}

func (d *Device) IsProtected() bool {
	orig := make([]byte, 1)
	err := d.Read(20, &orig)
	if err != nil {
		return true
	}
	var replace []byte
	if orig[0] == '#' {
		replace = []byte("+")
	} else {
		replace = []byte("#")
	}
	_ = d.Write(20, replace)
	newValue := make([]byte, 1)
	_ = d.Read(20, &newValue)
	if replace[0] == newValue[0] {
		_ = d.Write(20, orig)
		return false
	}
	return true
}
