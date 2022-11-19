package server

import (
	"fmt"
	"time"
	"net/http"
	"os"
	"os/signal"
	"context"
)

type ServerSettings struct{
	Name   		string
	Port   		int
	IP     		string
	IPVer  		string
	ReadTO		time.Duration
	WriteTO		time.Duration
}

type Server struct {
	ServerSettings
	modules		*RTTree
	server		*http.Server
	shutdownSig, terminatedSig	chan int

}

func (s *Server) Init() error{
	s.modules = &(RTTree{})
	s.modules.Init()
	s.shutdownSig = make(chan int)
	s.terminatedSig = make(chan int)
	fmt.Printf("Server '%s' Inited. \n", s.Name)
	return nil
}

func (s *Server) Add(route string, module IModule) error{
	return s.modules.Add(route, module)
}


type _Handler struct{
	s *Server
}
// implementation of http.Handler
func (h _Handler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	h.s.modules.Run(r.URL.String(), &ModuleContext{r,w,nil})
}



func (s *Server) Serve() error {

	addr := fmt.Sprintf("%s:%d", s.IP, s.Port)

	// http.HandleFunc("/", )

	s.server = &http.Server{
		Addr: addr,
		ReadTimeout: s.ReadTO,
		ReadHeaderTimeout: s.ReadTO,
		WriteTimeout: s.WriteTO,
		Handler: _Handler{s},
	}

	ctx, cancel:= context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		for{
			select {
			case <-sigint:
				fmt.Println("Interrupt signal. ")
				s.shutdownSig <- 1
				fmt.Println("Waiting for shutdown. ")
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for{
			select {
			case <-s.shutdownSig:
				fmt.Println("Handling shutdown...")
				if err := s.server.Shutdown(context.Background()); err != nil {
					fmt.Printf("Server Shutdown Error: %v", err)
				}else{
					fmt.Println("Server shutdown successfully. \nGoodbye. :D ")
				}
				close(s.terminatedSig)
				return
			case <-ctx.Done():
				fmt.Println("Closed inproperly. ")
				return
			default:
				time.Sleep(1*time.Second)
			}
		}
	}()

	fmt.Printf("Listening at: %s\n", addr)
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Server ListenAndServe Error: %v", err)
	}
	
	<-s.terminatedSig
	return nil
}

func (s *Server) Stop() error {
	s.shutdownSig <- 1
	<-s.terminatedSig
	return nil
}
