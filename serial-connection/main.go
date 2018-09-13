package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	s, err := serial.OpenPort(&serial.Config{
		Name: "/dev/rfcomm0",
		Baud: 9600,
	})

	if err != nil {
		log.Fatalf("Failed to open serial port %v", err)
	}
	defer s.Close()

	if _, err := s.Write([]byte{0x03}); err != nil {
		fmt.Printf("failed writing: %v", err)
	}
}
