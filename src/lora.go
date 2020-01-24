package main

import (
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
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	// set pins in output mode
	log.Println("setting up ss pin")
	pinSS := rpio.Pin(ssPin)
	defer pinSS.PullDown()
	pinSS.Output()
	log.Println("declaring receiver functions")

	log.Println("setting up dio pin")
	pinDIO := rpio.Pin(dio0)
	defer pinDIO.PullDown()
	pinDIO.Output()

	log.Println("setting up rst pin")
	pinRST := rpio.Pin(RST)
	defer pinRST.PullDown()
	pinRST.Output()

	// setup spi
	log.Println("setting up spi")
	rpio.SpiBegin(rpio.Spi0)
	defer rpio.SpiEnd(rpio.Spi0)
	rpio.SpiSpeed(500000)
	rpio.SpiChipSelect(0)

	// setup lora
	log.Println("setting up lora")
	log.Println("writing rst pin - low")
	pinRST.High()
	time.Sleep(time.Millisecond * 100)
	log.Println("writing rst pin - high")
	pinRST.Low()
	time.Sleep(time.Millisecond * 100)

	log.Println("reading version")
	selectReceiver := func(pin rpio.Pin) {
		pin.Low()
	}
	unselectReceiver := func(pin rpio.Pin) {
		pin.High()
	}
	readReg := func(addr byte, pin rpio.Pin) byte {
		log.Println("running select receiver")
		selectReceiver(pin)
		var spibuf = [2]byte{}
		spibuf[0] = addr & 0x7f
		spibuf[1] = 0x00
		log.Println("transmitting spi data")
		rpio.SpiTransmit(spibuf[0], spibuf[1])
		log.Println("unslecting receiver")
		unselectReceiver(pin)
		return spibuf[1]
	}
	var sx1272, sx1276 bool
	version := readReg(byte(REG_VERSION))
	if version == 0x22 {
		log.Println("SX1272 detected")
		sx1272 = true
	} else {
		pinRST.Write(rpio.Low)
		time.Sleep(time.Millisecond * 100)
		pinRST.Write(rpio.High)
		time.Sleep(time.Millisecond * 100)
		version = readReg(byte(REG_VERSION), pinSS)
		if version == 0x12 {
			sx1276 = true
		} else {
			log.Fatalf("unrecognized transceiver: %v", version)
		}
	}
	_, _ = sx1272, sx1276
	log.Println("version: ", string(version))
}
