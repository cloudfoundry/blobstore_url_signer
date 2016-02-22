package server

import (
	"net"
	"net/http"
	"os"
)

type Server interface {
	Start()
	Stop()
}

type server struct {
	protocal   string
	unixSocket string
	listener   net.Listener
	handlers   ServerHandlers
}

func NewServer(protocal string, socket string, handlers ServerHandlers) Server {
	listener, err := net.Listen(protocal, socket)
	if err != nil {
		panic(err)
	}

	return server{
		unixSocket: socket,
		listener:   listener,
		handlers:   handlers,
	}
}

func (s server) Start() {
	http.Serve(s.listener, http.HandlerFunc(s.handlers.SignUrl))
}

func (s server) Stop() {
	s.listener.Close()
	if s.protocal == "unix" {
		os.Remove(s.unixSocket)
	}
}
