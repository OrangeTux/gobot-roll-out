package gpio

// gobot/drivers/gpio/direct_pin_driver.go

// DigitalWrite writes to the pin. Acceptable values are 1 or 0
func (d *DirectPinDriver) DigitalWrite(level byte) (err error) {
	if writer, ok := d.Connection().(DigitalWriter); ok { // HL
		return writer.Connection(d.Pin(), level) // HL
	} // HL
	err = ErrDigitalWriteUnsupported
	return
}
