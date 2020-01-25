package main

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	// SX1272 - Raspberry connections
	ssPin   = 6 // (nss)
	dio0    = 7 // (dio)
	RST     = 0 // (reset)
	channel = 0 // (or rpio.Spio)
)

func main() {
	readReg := func(addr byte, pin rpio.Pin) byte {
		pin.High()
		var spibuf = []byte{addr & 0x7F, 0x00}
		rpio.SpiExchange(spibuf)
		pin.Low()
		return spibuf[1]
	}
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	// TODO(bonedaddy): do we need this
	// rpio.SpiChipSelect(0)
	// set pins in output mode
	log.Println("setting up pins")
	pinSS := rpio.Pin(ssPin)
	pinDIO := rpio.Pin(dio0)
	pinRST := rpio.Pin(RST)
	pinSS.Output()
	pinDIO.Output()
	pinRST.Output()

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
	pinRST.High()
	time.Sleep(time.Millisecond * 100)
	pinRST.Low()
	time.Sleep(time.Millisecond * 100)

	log.Println("reading version")
	var sx1272, sx1276 bool
	version := readReg(byte(REG_VERSION), pinSS)
	fmt.Println("sx1272 ", 0x22)
	fmt.Println("sx1276 ", 0x12)
	if version == 0x22 {
		log.Println("SX1272 detected")
		sx1272 = true
	} else {
		pinRST.Low()
		time.Sleep(time.Millisecond * 100)
		pinRST.High()
		time.Sleep(time.Millisecond * 100)
		version = readReg(byte(REG_VERSION), pinSS)
		if version == 0x12 {
			sx1276 = true
			log.Println("SX1276 detected")
		} else {
			log.Fatalf("unrecognized transceiver: %v", version)
		}
	}
	_, _ = sx1272, sx1276
	log.Println("version: ", string(version))
}
