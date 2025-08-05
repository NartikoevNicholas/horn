package horn

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

type Server struct {
	config   Config
	listener net.Listener
	workers  []Worker
}

func (s *Server) Init() {
	s.config.ApplyDefault()
	s.setListener()
}

func (s *Server) Startup() {
	for id := 1; id <= s.config.WorkerCount; id++ {
		worker := Worker{id: id, listener: s.listener}
		s.workers = append(s.workers, worker)
		go worker.requestHandler(s.config.ProtocolHandler)
	}
	fmt.Println("Server is started")
	defer s.Shutdown()
}

func (s *Server) Shutdown() {
	s.listener.Close()
	fmt.Println("Server is stopped")
}

func (s *Server) setListener() {
	var err error
	if s.listener != nil {
		return
	}

	s.listener, err = net.Listen(s.config.Transport, s.getAddr())
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) getAddr() string {
	return s.config.Host + ":" + strconv.Itoa(s.config.Port)
}
