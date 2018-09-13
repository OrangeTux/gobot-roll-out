package main

import (
	"fmt" // OMIT
	// OMIT
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
)

func main() {
	adaptor := joystick.NewAdaptor()
	controller := joystick.NewDriver(adaptor, "dualshock3")

	work := func() { // HL
		controller.On(joystick.RightY, func(data interface{}) { // HL
			fmt.Printf("Y: %v\n", data) // HL
		}) // HL
	} // HL

	robot := gobot.NewRobot("ps3-controller",
		[]gobot.Connection{adaptor},
		[]gobot.Device{controller},
		work,
	)

	robot.Start()
}
