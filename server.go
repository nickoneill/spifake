package spi

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Server does x
type Server struct {
	pattern string
	clients map[int]*Client
	sendCh  chan []byte
	doneCh  chan bool
	errCh   chan error
}

// NewServer creates new chat server.
func NewServer(pattern string) *Server {
	clients := make(map[int]*Client)
	sendCh := make(chan []byte)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		clients,
		sendCh,
		doneCh,
		errCh,
	}
}

func (s *Server) send(msg []byte) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}

// Listen and serve
func (s *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.clients[client.id] = client
		client.Listen()
	}

	http.Handle(s.pattern, websocket.Handler(onConnected))

	for {
		select {
		case msg := <-s.sendCh:
			s.send(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}

func (s *Server) Send(msg []byte) {
	s.sendCh <- msg
}

func (s *Server) Err(err error) {
	s.errCh <- err
}
