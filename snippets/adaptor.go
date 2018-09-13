package gobot

type Adaptor interface {
	Name() string

	SetName(n string)

	Connect() error

	Finalize() error
}
