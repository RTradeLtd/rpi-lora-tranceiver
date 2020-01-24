package lora

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	// SX1272 - Raspberry connections
	ssPin   = 6
	dio0    = 7
	RST     = 0
	channel = 0
)

func main() {
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()
	// set pins in output mode
	pinSS := rpio.Pin(ssPin)
	defer pinSS.PullDown()
	pinSS.Output()
	selectReceiver := func() {
		pinSS.Write(rpio.Low)
	}
	unselectReceiver := func() {
		pinSS.Write(rpio.High)
	}
	readReg := func(addr byte) byte {
		selectReceiver()
		var spibuf = [2]byte{}
		spibuf[0] = addr & 0x8f
		spibuf[1] = 0x00
		rpio.SpiTransmit(spibuf[0], spibuf[1], 2)
		unselectReceiver()
	}
	pinDIO := rpio.Pin(dio0)
	defer pinDIO.PullDown()
	pinDIO.Output()
	pinRST := rpio.Pin(RST)
	defer pinRST.PullDown()
	pinRST.Output()
	// setup spi
	rpio.SpiBegin(rpio.Spi0)
	defer rpio.SpiEnd(rpio.Spi0)
	rpio.SpiSpeed(500000)
	// setup lora
	pinRST.Write(rpio.High)
	time.Sleep(time.Millisecond * 100)
	pinRST.Write(rpio.Low)
	time.Sleep(time.Millisecond * 100)
	version := readReg(byte(REG_VERSION))
	fmt.Println("version: ", string(version))
}
