package spi

import (
	"log"
	"net/http"
	"time"
)

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
	// v := string(buf[:])

	d.server.Send(buf[:])

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

	http.Handle("/", http.FileServer(http.Dir("/Users/nickoneill/gocode/src/spi/web")))
	go func() { log.Fatal(http.ListenAndServe(":5480", nil)) }()

	return &Device{server: s}, nil
}
