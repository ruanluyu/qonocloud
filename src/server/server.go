package server

import (
	"fmt"
	"net/http"
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

func handler(server *Server, writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, fmt.Sprintf("MiLai Cloud Server \nMethod: %s\nHost: %s\nURL: %s",
		req.Method,
		req.Host,
		req.URL))
	fmt.Printf("Request recieved:\n%s\n%s\n%s\n", req.Method, req.Host, req.URL)
}


func (s *Server) Serve() {
	addr := fmt.Sprintf("http://%s:%d/", s.IP, s.Port)
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request){
		handler(s, writer, req)
	})
	http.ListenAndServe(addr, nil)
}

func (s *Server) Stop() {
	
}
