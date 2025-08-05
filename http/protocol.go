package http

import (
	"github.com/NartikoevNicholas/horn/types"
)

func GetProtocolHandler() types.ProtocolHandler {
	return types.ProtocolHandler{Request: &Request{}, Response: Response{}}
}
