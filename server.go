package main

import (
	"net/http"
	"github.com/jensluch/stringutil"
	"fmt"
)

type Server struct {
	Pattern string
	Addr string
	Port string
}

func (s *Server) addr() string {
	return s.Addr + s.Port
}

func (s *Server) SetPattern(pattern string) {
	s.Pattern = pattern
}

func (s *Server) SetAddr(addr string) {
	s.Addr = addr
}

func (s *Server) SetPort(port string) {
	s.Port = port
}

func (s *Server) Start() {
	http.HandleFunc(s.Pattern, handler)
	http.ListenAndServe(s.addr(), nil)
}

// handlers

func handler(w http.ResponseWriter, r *http.Request) {
	reversed := stringutil.Reverse("!oG ,olleH")
	prepared := reversed + "   ---  %s!"
	fmt.Fprintf(w, prepared, r.URL.Path[1:])
}