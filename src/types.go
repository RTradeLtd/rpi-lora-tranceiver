package main

var (
	REG_FIFO                 = 0x00
	REG_OPMODE               = 0x01
	REG_FIFO_ADDR_PTR        = 0x0D
	REG_FIFO_TX_BASE_AD      = 0x0E
	REG_FIFO_RX_BASE_AD      = 0x0F
	REG_RX_NB_BYTES          = 0x13
	REG_FIFO_RX_CURRENT_ADDR = 0x10
	REG_IRQ_FLAGS            = 0x12
	REG_DIO_MAPPING_1        = 0x40
	REG_DIO_MAPPING_2        = 0x41
	REG_MODEM_CONFIG         = 0x1D
	REG_MODEM_CONFIG2        = 0x1E
	REG_MODEM_CONFIG3        = 0x26
	REG_SYMB_TIMEOUT_LSB     = 0x1F
	REG_PKT_SNR_VALUE        = 0x19
	REG_PAYLOAD_LENGTH       = 0x22
	REG_IRQ_FLAGS_MASK       = 0x11
	REG_MAX_PAYLOAD_LENGTH   = 0x23
	REG_HOP_PERIOD           = 0x24
	REG_SYNC_WORD            = 0x39
	REG_VERSION              = 0x42
	PAYLOAD_LENGTH           = 0x40

	// low noise amplifier
	REG_LNA        = 0x0C
	LNA_MAX_GAIN   = 0x23
	LNA_OFF_GAIN   = 0x00
	LNA_LOW_GAIN   = 0x20
	RegDioMapping1 = 0x40 // common
	RegDioMapping2 = 0x41 // common
	RegPaConfig    = 0x09 // common
	RegPaRamp      = 0x0A // common
	RegPaDac       = 0x5A // common
	SX72_MC2_FSK   = 0x00
	SX72_MC2_SF7   = 0x70
	SX72_MC2_SF8   = 0x80
	SX72_MC2_SF9   = 0x90
	SX72_MC2_SF10  = 0xA0
	SX72_MC2_SF11  = 0xB0
	SX72_MC2_SF12  = 0xC0

	SX72_MC1_LOW_DATA_RATE_OPTIMIZE = 0x01 // mandated for SF11 and SF12

	// sx1276 RegModemConfig1
	SX1276_MC1_BW_125 = 0x70
	SX1276_MC1_BW_250 = 0x80
	SX1276_MC1_BW_500 = 0x90
	SX1276_MC1_CR_4_5 = 0x02
	SX1276_MC1_CR_4_6 = 0x04
	SX1276_MC1_CR_4_7 = 0x06
	SX1276_MC1_CR_4_8 = 0x08

	SX1276_MC1_IMPLICIT_HEADER_MODE_ON = 0x01

	// sx1276 RegModemConfig2
	SX1276_MC2_RX_PAYLOAD_CRCON = 0x04

	// sx1276 RegModemConfig3
	SX1276_MC3_LOW_DATA_RATE_OPTIMIZE = 0x08
	SX1276_MC3_AGCAUTO                = 0x04

	// preamble for lora networks (nibbles swapped)
	LORA_MAC_PREAMBLE = 0x34

	RXLORA_RXMODE_RSSI_REG_MODEM_CONFIG1 = 0x0A

	/* how to handle:
	#define RXLORA_RXMODE_RSSI_REG_MODEM_CONFIG1 0x0A
	#ifdef LMIC_SX1276
	#define RXLORA_RXMODE_RSSI_REG_MODEM_CONFIG2 0x70
	#elif LMIC_SX1272
	#define RXLORA_RXMODE_RSSI_REG_MODEM_CONFIG2 0x74
	#endif
	*/

	REG_FRF_MSB = 0x06
	REG_FRF_MID = 0x07
	REG_FRF_LSB = 0x08
	FRF_MSB     = 0xD9 // 868.1 Mhz
	FRF_MID     = 0x06
	FRF_LSB     = 0x66
)

const (
	// Constants for radio registers
	OPMODE_LORA      = 0x80
	OPMODE_MASK      = 0x07
	OPMODE_SLEEP     = 0x00
	OPMODE_STANDBY   = 0x01
	OPMODE_FSTX      = 0x02
	OPMODE_TX        = 0x03
	OPMODE_FSRX      = 0x04
	OPMODE_RX        = 0x05
	OPMODE_RX_SINGLE = 0x06
	OPMODE_CAD       = 0x07

	// Bits masking the corresponding IRQs from the radio
	IRQ_LORA_RXTOUT_MASK = 0x80
	IRQ_LORA_RXDONE_MASK = 0x40
	IRQ_LORA_CRCERR_MASK = 0x20
	IRQ_LORA_HEADER_MASK = 0x10
	IRQ_LORA_TXDONE_MASK = 0x08
	IRQ_LORA_CDDONE_MASK = 0x04
	IRQ_LORA_FHSSCH_MASK = 0x02
	IRQ_LORA_CDDETD_MASK = 0x01

	// DIO function mappings                D0D1D2D3
	MAP_DIO0_LORA_RXDONE = 0x00 // 00------
	MAP_DIO0_LORA_TXDONE = 0x40 // 01------
	MAP_DIO1_LORA_RXTOUT = 0x00 // --00----
	MAP_DIO1_LORA_NOP    = 0x30 // --11----
	MAP_DIO2_LORA_NOP    = 0xC0 // ----11--
)
