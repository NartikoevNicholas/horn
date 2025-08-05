package horn

import (
	"fmt"
	"net"

	"github.com/NartikoevNicholas/horn/types"
)

type Worker struct {
	id       int
	listener net.Listener
}

func (w *Worker) requestHandler(protocolHandler types.ProtocolHandler) {
	for {
		conn, err := w.listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		buffer := getBuffer(conn)
		protocolHandler.Request.Parse(buffer)
		fmt.Println(buffer)
		conn.Close()
	}
}
