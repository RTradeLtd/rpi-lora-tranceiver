package main

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	// SX1272 - Raspberry connections
	ssPin   = 6
	dio0    = 7
	RST     = 0
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
	pinRST.Write(rpio.High)
	time.Sleep(time.Millisecond * 100)
	log.Println("writing rst pin - high")
	pinRST.Write(rpio.Low)
	time.Sleep(time.Millisecond * 100)

	log.Println("reading version")
	selectReceiver := func() {
		pinSS.Write(rpio.Low)
	}
	unselectReceiver := func() {
		pinSS.Write(rpio.High)
	}
	readReg := func(addr byte) byte {
		log.Println("running select receiver")
		selectReceiver()
		var spibuf = [2]byte{}
		spibuf[0] = addr & 0x7f
		spibuf[1] = 0x00
		log.Println("transmitting spi data")
		rpio.SpiTransmit(spibuf[0], spibuf[1])
		log.Println("unslecting receiver")
		unselectReceiver()
		return spibuf[1]
	}
	version := readReg(byte(REG_VERSION))
	log.Println("version: ", string(version))
}
