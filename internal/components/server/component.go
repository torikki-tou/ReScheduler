package server

import (
	"net/http"
	"strconv"
)

type Server struct {
	mux    mux
	config config
}

func New(mux mux, config config) *Server {
	return &Server{
		mux:    mux,
		config: config,
	}
}

func (s *Server) Run() error {
	addr := ":" + strconv.Itoa(s.config.AppPort())
	return http.ListenAndServe(addr, s.mux)
}
