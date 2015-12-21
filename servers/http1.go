package servers

import (
	"net/http"
	"time"
)

func NewHttp1Server() *Http1Server {
	return &Http1Server{}
}

type Http1Server struct {
	http.Server
}

func (s *Http1Server) BindAddress() string { return s.Addr }

func (s *Http1Server) SetBindAddress(addr string) {
	s.Addr = addr
}

func (s *Http1Server) SetTimeouts(read, write time.Duration) {
	s.ReadTimeout = read
	s.WriteTimeout = write
}

func (s *Http1Server) SetDispatcher(h http.Handler) {
	s.Handler = h
}

func (s *Http1Server) Run() error {
	return s.ListenAndServe()
}
