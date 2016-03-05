package spi

import "time"

// mocked spi interface matching
// https://github.com/rakyll/experimental/blob/master/spi/spi.go
// which really just forwards via websockets

const (
	Mode0 = 0x00
	Mode1 = 0x04
	Mode2 = 0x08
	Mode3 = 0x0C
)

type Device struct {
	server *Server
}

func (d *Device) Do(buf []byte, delay time.Duration) error {
	d.server.Send(buf[:])

	// delay arbitrary amount because we're not spi and otherwise
	// would have no delay
	tick := time.Duration(50) * time.Millisecond
	time.Sleep(tick)

	return nil
}

func (d *Device) SetMode(mode int) error {
	return nil
}

func (d *Device) SetSpeed(speedHz int) error {
	return nil
}

func (d *Device) SetBitsPerWord(bits int) error {
	return nil
}

func (d *Device) Close() error {
	// no file, nothing to close
	return nil
}

func Open(name string) (*Device, error) {
	s := NewServer("/spi")
	go s.Listen()

	// xmas.go already does this
	// go func() { log.Fatal(http.ListenAndServe(":5480", nil)) }()

	return &Device{server: s}, nil
}
