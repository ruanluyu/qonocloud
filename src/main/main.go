package main

import (
	"fmt"
	"milai/qonocloud/server"
	"time"
)

//"fmt"
//"net/http"

var defaultPort int = 8000

type HomeModule struct{
	greetingInfo string
}

func (m *HomeModule) Run(context *server.ModuleContext) error{
	fmt.Println("Entered HomeModule")
	context.Response.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(context.Response, "Done. \n")
	return nil
}

type MatchModule struct{
	matchedInfo string
}

func (m *MatchModule) Run(context *server.ModuleContext) error{
	fmt.Fprintf(context.Response, "Matched.\n")
	context.Fallback()
	return nil
}


func main() {
	fmt.Println("Start...")
	settings := server.ServerSettings{
			Name: "Test server",
			Port: defaultPort,
			IP: "0.0.0.0",
			ReadTO: 20 * time.Second,
			WriteTO: 20 * time.Second,}
	s := &server.Server{
		ServerSettings: settings}
	s.Init()
	s.Add("/", &HomeModule{""})
	s.Add("/mazfer/", &MatchModule{""})
	s.Serve()
}
