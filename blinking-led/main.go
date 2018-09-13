package main

import (
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	adaptor := raspi.NewAdaptor()

	pin := gpio.NewDirectPinDriver(adaptor, "8")

	pin.Off()
	time.Sleep(time.Second * 5)
	pin.On()
}
