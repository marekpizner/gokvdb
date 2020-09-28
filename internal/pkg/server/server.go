package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run(hostPort string) error {
	log.Printf("[INFO] server: runing on %s\n", hostPort)
	l, err := net.Listen("tcp", hostPort)
	if err != nil {
		return fmt.Errorf("server: could not listen %s: %v", hostPort, err)
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("[WARN] server: could not accept connection: %v\n", err)
			continue
		}
		cl := newClient(conn)
		go s.handleClient(cl)
	}
}

func (s *Server) handleClient(cl *client) {
	defer cl.Close()

	cl.respondWithCommandWaiting()

	scanner := bufio.NewScanner(cl.conn)
	for scanner.Scan() {
		input := scanner.Text()
		log.Printf("[INFO] server: recieve command string: %s", input)
		cl.respondWithCommandWaiting()
	}

	if err := scanner.Err(); err != nil {
		log.Printf("[WARN] server: scanner error: %v", err)
	}
}
