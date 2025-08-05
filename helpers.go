package horn

import (
	"fmt"
	"log"
	"net"

	"github.com/NartikoevNicholas/horn/http"
	"github.com/NartikoevNicholas/horn/types"
	"github.com/NartikoevNicholas/horn/utils"
)

func getProtocolHandler(protocol string) types.ProtocolHandler {
	var result types.ProtocolHandler
	if _, ok := utils.SupportedProtocols[protocol]; ok == false {
		log.Fatal(fmt.Printf("Protocol '%s' is not supported!", protocol))
	}

	if protocol == "http" {
		result = http.GetProtocolHandler()
	}
	return result
}

func getBuffer(conn net.Conn) []byte {
	return nil
}
