package main

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	// SX1272 - Raspberry connections
	ssPin   = 25 // (nss)
	dio0    = 4  // (dio)
	RST     = 7  // (reset)
	channel = 0  // (or rpio.Spio)
	/* original
	ssPin   = 6 // (nss)
	dio0    = 7 // (dio)
	RST     = 0 // (reset)
	*/
	/* from https://github.com/computenodes/dragino/blob/master/dragino/SX127x/board_config.py
	   DIO0 = 4   # RaspPi GPIO 4
	   DIO1 = 23   # RaspPi GPIO 23
	   DIO2 = 24   # RaspPi GPIO 24
	   DIO3 = None # Not connected on dragino header
	   LED  = 18   # RaspPi GPIO 18 connects to the LED on the proto shield
	   SWITCH = 4  # RaspPi GPIO 4 connects to a switch
	   SPI_CS = 2  # Chip Select pin to use
	*/
	reg byte = 0x42
)

func main() {
	readReg := func(addr byte) byte {
		rpio.WritePin(rpio.Pin(ssPin), rpio.Low)
		var spibuf = []byte{addr & 0x7F, 0x00}
		rpio.SpiExchange(spibuf)
		rpio.WritePin(rpio.Pin(RST), rpio.High)
		return spibuf[1]
	}
	/*writeReg := func(addr byte, value byte, pin rpio.Pin) {
		pin.High()
		var spibuf = []byte{addr | 0x80, value}
		rpio.SpiExchange(spibuf)
		pin.Low()
	}*/
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	// TODO(bonedaddy): do we need this
	// rpio.SpiChipSelect(0)
	// set pins in output mode
	log.Println("setting up pins")
	rpio.PinMode(rpio.Pin(ssPin), rpio.Output)
	rpio.PinMode(rpio.Pin(dio0), rpio.Output)
	rpio.PinMode(rpio.Pin(RST), rpio.Output)

	// setup spi
	log.Println("setting up spi")
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		log.Fatal(err)
	}
	defer rpio.SpiEnd(rpio.Spi0)
	rpio.SpiSpeed(500000)
	rpio.SpiChipSelect(0)

	// setup lora
	log.Println("setting up lora")
	rpio.WritePin(rpio.Pin(RST), rpio.High)
	time.Sleep(time.Millisecond * 100)
	rpio.WritePin(rpio.Pin(RST), rpio.Low)
	time.Sleep(time.Millisecond * 100)

	log.Println("reading version")
	var sx1272, sx1276 bool
	version := readReg(reg)
	fmt.Println("sx1272 ", 0x22)
	fmt.Println("sx1276 ", 0x12)
	if version == 0x22 {
		log.Println("SX1272 detected")
		sx1272 = true
	} else {
		rpio.WritePin(rpio.Pin(RST), rpio.Low)
		time.Sleep(time.Millisecond * 100)
		rpio.WritePin(rpio.Pin(RST), rpio.High)
		time.Sleep(time.Millisecond * 100)
		version = readReg(reg)
		if version == 0x12 {
			sx1276 = true
			log.Println("SX1276 detected")
		} else {
			log.Fatalf("unrecognized transceiver: %v", version)
		}
	}
	_, _ = sx1272, sx1276
	log.Println("version: ", string(version))
	//	writeReg(byte(RegPaRamp), (readReg(byte(RegPaRamp)&0xF0|0x08, pinSS)), pinSS)
}
