package sfp

// description of the physical device
var ModuleIdentifier = map[uint64]string{
	0x00: "Unknown or unspecified",
	0x01: "GBIC",
	0x02: "Module/connector soldered to motherboard ",
	0x03: "SFP or SFP+",
}

// ModuleConnector description
var ModuleConnector = map[uint64]string{
	0x00: "Unknown or unspecified",
	0x01: "SC",
	0x02: "Fibre Channel Style 1 copper connector",
	0x03: "Fibre Channel Style 2 copper connector",
	0x04: "BNC/TNC",
	0x05: "Fibre Channel coaxial headers",
	0x06: "FiberJack",
	0x07: "LC",
	0x08: "MT-RJ",
	0x09: "MU",
	0x0A: "SG",
	0x0B: "Optical pigtail",
	0x0C: "MPO Parallel Optic",
	0x20: "HSSDC II",
	0x21: "Copper Pigtail",
	0x22: "RJ45",
}

// Compliance: indicators define which electronic
// or optical interfaces are supported by the transceiver
var Compliance = map[uint64]string{
	1 << 0:  "IB: 1X copper passive",
	1 << 1:  "IB: 1X copper active",
	1 << 2:  "IB: 1X LX",
	1 << 3:  "IB: 1X SX",
	1 << 4:  "10G Base-SR",
	1 << 5:  "10G Base-LR",
	1 << 6:  "10G Base-LRM",
	1 << 7:  "10G Base-ER",
	1 << 8:  "OC-3 short reach",
	1 << 9:  "OC-3 SM, intermediate reach",
	1 << 10: "OC-3 SM, long reach",
	1 << 12: "OC-12 short reach",
	1 << 13: "OC-12 SM, intermediate reach",
	1 << 14: "OC-12 SM, long reach",
	1 << 16: "OC-48 short reach",
	1 << 17: "OC-48 intermediate reach",
	1 << 18: "OC-48 long reach",
	1 << 19: "SONET reach bit 1",
	1 << 20: "SONET reach bit 2",
	1 << 21: "OC-192 short reach",
	1 << 22: "ESCON SMF, 1310nm laser",
	1 << 23: "ESCON MMF, 1310nm LED",
	1 << 24: "1000Base-SX",
	1 << 25: "1000Base-LX",
	1 << 26: "1000Base-CX",
	1 << 27: "1000Base-T",
	1 << 28: "100Base-LX/LX10",
	1 << 29: "100Base-FX",
	1 << 30: "Base-BX/10",
	1 << 31: "Base-PX",
	1 << 32: "FC: electrical inter-enclosure (EL)",
	1 << 33: "FC: longwave laser (LC)",
	1 << 34: "FC: shortwave laser, linear Rx (SA)",
	1 << 35: "FC: medium distance (M)",
	1 << 36: "FC: long distance (L)",
	1 << 37: "FC: intermediate distance (I)",
	1 << 38: "FC: short distance (S)",
	1 << 39: "FC: very long distance (V)",
	1 << 42: "SFP+ Passive cable",
	1 << 43: "SFP+ Active cable",
	1 << 44: "FC: Longwave laser (LL)",
	1 << 45: "FC: Shortwave laser with OFC (SL)",
	1 << 46: "FC: Shortwave laser w/o OFC (SN)",
	1 << 47: "FC: Electrical intra-inclosure (EL)",
	1 << 48: "FC media: Single mode (SM)",
	1 << 50: "FC media: Multimode, 50um (M5, M5E)",
	1 << 51: "FC media: Multimode, 62.5um (M6)",
	1 << 52: "FC media: Video coax (TV)",
	1 << 53: "FC media: Miniature coax (MI)",
	1 << 54: "FC media: Twisted pair (TP)",
	1 << 55: "FC media: Twin axial pair (TW))",
	1 << 56: "FC speed: 100MB/sec",
	1 << 58: "FC speed: 200MB/sec",
	1 << 60: "FC speed: 400MB/sec",
	1 << 61: "FC speed: 1600MB/sec",
	1 << 62: "FC speed: 800MB/sec",
	1 << 63: "FC speed: 1200MB/sec",
}

// Encoding serial encoding mechanism
var Encoding = map[uint64]string{
	0x00: "Unspecified",
	0x01: "8B10B",
	0x02: "4B5B",
	0x03: "NRZ",
	0x04: "Manchester",
	0x05: "SONET Scrambled",
	0x06: "64B/66B ",
}

// RateIdentifier
var RateIdentifier = map[uint64]string{
	0x00: "Unspecified",
	0x01: "SFF-8079 (4/2/1G Rate Select and AS0/AS1)",
	0x02: "SFF-8431 (8/4/2G RX Rate Select Only)",
	0x03: "Unspecified",
	0x04: "SFF-8431 (8/4/2G TX Rate Select Only)",
	0x05: "Unspecified",
	0x06: "SFF-8431 (8/4/2G Independent TX and RX Rate Select)",
	0x07: "Unspecified",
	0x08: "FC-PI-5 (16/8/4G RX Rate Select Only) High=16G, Low=8/4G",
	0x09: "Unspecified",
	0x0A: "FC-PI-5 (16/8/4G Independent TX and RX Rate Select) High=16G, Low=8/4G",
}

// Options implemented in the transceiver
var Options = map[uint64]string{
	1 << 0:  "Linear Receiver Output Implemented",
	1 << 1:  "Power level 2 required",
	1 << 2:  "Cooled laser transmitter",
	1 << 9:  "Rx_LOS",
	1 << 10: "Signal detect (inverted Rx_LOS)",
	1 << 11: "TX_FAULT",
	1 << 12: "TX_DISABLE",
	1 << 13: "RATE_SELECT",
}

// DDM Diagnostic Monitoring Types
var DDMTypes = map[uint64]string{
	1 << 2: "Address change required",
	1 << 3: "Average input power",
	1 << 4: "Externally calibrated",
	1 << 5: "Internally calibrated",
	1 << 6: "DDM present",
}

// EnhancedOptions describe optional digital diagnostic features implemented
var EnhancedOptions = map[uint64]string{
	1 << 7: "Optional Alarm/warning flags implemented",
	1 << 6: "Soft TX_DISABLE control and monitoring implemented",
	1 << 5: "Soft TX_FAULT monitoring implemented",
	1 << 4: "Soft RX_LOS monitoring implemented",
	1 << 3: "Soft RATE_SELECT control and monitoring implemented",
	1 << 2: "Application Select control implemented per SFF-8079",
	1 << 1: "Soft Rate Select control implemented per SFF-8431",
}

// SFF8472 implemented feature set
var SFF8472 = map[uint64]string{
	0x00: "DDM not included or defined",
	0x01: "Includes functionality from Rev 9.3 SFF-8472",
	0x02: "Includes functionality from Rev 9.5 SFF-8472",
	0x03: "Includes functionality from Rev 10.2 SFF-8472",
	0x04: "Includes functionality from Rev 10.4 SFF-8472",
	0x05: "Includes functionality from Rev 11.0 SFF-8472",
}
