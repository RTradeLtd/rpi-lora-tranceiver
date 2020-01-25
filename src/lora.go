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
	/*selectReceiver := func(pin rpio.Pin) {
		pin.Low()
	}
	unselectReceiver := func(pin rpio.Pin) {
		pin.High()
	}*/
	readReg := func(addr byte, pin rpio.Pin) byte {
		log.Println("running select receiver")
		//selectReceiver(pin)
		var spibuf = []byte{addr & 0x7F, 0x00}
		log.Println("transmitting spi data")
		rpio.SpiTransmit(spibuf[0], spibuf[1])
		rpio.SpiExchange(spibuf)
		//unselectReceiver(pin)
		return spibuf[1]
	}
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		log.Fatal(err)
	}
	defer rpio.SpiEnd(rpio.Spi0)
	// TODO(bonedaddy): do we need this
	// rpio.SpiChipSelect(0)
	// set pins in output mode
	log.Println("setting up pins")
	pinSS := rpio.Pin(ssPin)
	pinDIO := rpio.Pin(dio0)
	pinRST := rpio.Pin(RST)
	/*pinSS.Output()
	pinDIO.Output()
	pinRST.Output()
	*/
	// setup spi
	log.Println("setting up spi")

	rpio.SpiSpeed(500000)

	/*// setup lora
	log.Println("setting up lora")
	pinRST.High()
	time.Sleep(time.Millisecond * 100)
	pinRST.Low()
	time.Sleep(time.Millisecond * 100)
	*/
	log.Println("reading version")
	var sx1272, sx1276 bool
	version := readReg(byte(REG_VERSION), pinSS)
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
