// hybridgroup/gobot/eventer.go

// Publish new events to anyone that is subscribed
func (e *eventer) Publish(name string, data interface{}) {
	evt := NewEvent(name, data)
	e.in <- evt // HL
}

// On executes the event handler f when e is Published to.
func (e *eventer) On(n string, f func(s interface{})) (err error) {
	out := e.Subscribe()
	go func() {
		for {
			select {
			case evt := <-out: // HL
				if evt.Name == n { // HL
					f(evt.Data) // HL
				}
			}
		}
	}()

	return
}
