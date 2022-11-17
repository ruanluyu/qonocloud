package server

import (
	"fmt"
)

type Server struct {
	Name   		string
	Port   		int
	IP     		string
	IPVer  		string
	Modules		RTTree
}

func (s *Server) Init() {
	fmt.Printf("Server '%s' init", s.Name)
}



func (s *Server) Run() {
	addr := fmt.Sprintf("http://%s:%d/", s.IP, s.Port)
	http.HandleFunc("/", handler)
	
}

func (s *Server) Stop() {

}
