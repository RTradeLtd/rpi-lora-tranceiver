package lora

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
}
