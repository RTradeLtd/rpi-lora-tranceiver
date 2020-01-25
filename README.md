# rpi-lora-tranceiver

To install the "app" run `install.sh`



# Golang Implementation


Active (no CGO):

* https://github.com/stianeikeland/go-rpio
  * https://godoc.org/github.com/stianeikeland/go-rpio
  * https://blog.eikeland.se/2013/07/30/go-gpio-library-for-raspberry-pi/
  
* https://github.com/kidoman/embd
* https://github.com/mrmorphic/hwio

Old:

* https://github.com/hugozhu/rpi
* https://github.com/flyingyizi/go-wiringPi


Resources:

* https://github.com/rakyll/go-hardware
* https://medium.com/@farissyariati/go-raspberry-pi-hello-world-tutorial-7e830d08b3ae
* https://docs.heltec.cn/#/en/user_manual/how_to_install_esp32_Arduino?id=_1-execute-a-example-likes-factorytestino


# GPIO Pin Outs

* http://wiki.dragino.com/images/9/9b/RPi_GPIO.jpg
* http://wiki.dragino.com/images/e/e1/Lora_hat_wiring.png


| LoRa GPS HAT | RaspberryPi Wiring PI IO |
|--------------|--------------------------|
| 3.3v | 3.3v |
| 5v | 5v |
| GND | GND | 
| DIO0 | GPIO7 |
| GPS_RX | GPIO15/TX  |
| GPS_TX | GPIO16/RX |
| RESET | GPIO0 |
| LoRa_NSS | GPIO6 |
| LoRa_MISO | GPIO13/MISO |
| LoRa_MOSI | GPIO12/MOSI |
| SCK | GPIO14/SCLK |
| DIO1 | GPIO4 |
| DIO2 | GPIO5 |
| 1PPS | GPIO1 |
