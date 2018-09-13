package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/platforms/joystick"
)

const (
	forward  = byte(0x01)
	backward = byte(0x03)

	leftToCenter = byte(0x04)
	left         = byte(0x05)

	rightToCenter = byte(0x06)
	right         = byte(0x07)
)

var (
	xDirection, yDirection byte
)

// SetX translates X movement of stick to a xDirection.
func setX(n int16) {
	switch true {
	case n < 0:
		xDirection = left
	case n > 0:
		xDirection = right
	default:
		if xDirection == left {
			xDirection = leftToCenter
		} else {
			xDirection = rightToCenter
		}
	}
}

// SetY translates Y movement of stick to a yDirection.
func setY(n int16) {
	yDirection = byte(0x02)
	if n < 0 {
		yDirection = forward
	} else if n > 0 {
		yDirection = backward
	}
}

func main() {
	baudrate := flag.Int("baudrate", 9600, "baudrate")
	device := flag.String("device", "/dev/rfcomm0", "serial device")
	flag.Parse()

	s, err := serial.OpenPort(&serial.Config{
		Name: *device,
		Baud: *baudrate,
	})

	if err != nil {
		log.Fatalf("Failed to open serial port %v: %v", *baudrate, err)
	}
	defer s.Close()

	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor, "dualshock3")

	work := func() {
		stick.On(joystick.LeftX, func(data interface{}) {
			n, ok := data.(int16)
			if !ok {
				fmt.Println("Failed to convert event data.")
				return
			}

			setX(n)
			s.Write([]byte{yDirection})
		})

		stick.On(joystick.RightY, func(data interface{}) {
			n, ok := data.(int16)
			if !ok {
				fmt.Println("Failed to convert event data.")
				return
			}

			setY(n)
			s.Write([]byte{yDirection})
		})

		// PS3 won't repeat Y direction when you hold Y stick in same
		// position. therefore we've to repeat the Y direction.
		gobot.Every(500*time.Millisecond, func() {
			if _, err := s.Write([]byte{yDirection}); err != nil {
				fmt.Printf("failed writing: %v", err)
			}

		})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{stick},
		work,
	)

	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	master.AddRobot(robot)
	master.Start()
}
