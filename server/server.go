package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Server interface {
	Start()
	Stop()
}

type server struct {
	port     int
	listener net.Listener
	handlers ServerHandlers
}

func NewServer(port int, addr string, handlers ServerHandlers) Server {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		panic(err)
	}

	return server{
		port:     port,
		listener: listener,
		handlers: handlers,
	}
}

func (s server) Start() {
	s.registerHandlers()
	http.Serve(s.listener, nil)
}

func (s server) registerHandlers() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/sign").HandlerFunc(s.handlers.SignUrl)

	http.Handle("/", r)
}

func (s server) Stop() {
	s.listener.Close()
}
