package horn

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

type Header string
type Headers map[string]Header

type server struct {
	app      App
	config   Config
	listener net.Listener
}

func (s server) startup() {
	s.setListener()

	fmt.Println("Server is startup")

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {

	// 	}
	// 	buffer := make([]byte, 1024)
	// 	conn.Read(buffer)
	// 	// request := parseRequest(buffer)
	// 	// response := app.requestHandler(request)
	// 	// conn.Write(getResponse(response))
	// 	n, err := conn.Write([]byte("тест"))
	// 	n += 1
	// 	conn.Close()
	// }
}

func (s *server) setListener() {
	var err error
	s.listener, err = net.Listen(s.config.Transport, s.addr())
	if err != nil {
		fmt.Println("Server is not")
		log.Fatal(err)
	}
}

func (s server) addr() string {
	return s.config.Host + ":" + strconv.Itoa(s.config.Port)
}
