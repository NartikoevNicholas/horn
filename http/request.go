package http

import (
	"bytes"
	"net"

	"github.com/NartikoevNicholas/horn/utils"
)

type Request struct {
	Conn  net.Conn
	Scope Scope
}

type Scope struct {
	Method     []byte
	Endpoint   []byte
	QueryParam [][2][]byte
	Version    []byte
	Headers    [][2][]byte
	Body       []byte
}

func (r *Request) Parse(request []byte) {
	parts := bytes.SplitN(request, utils.Del.Body, 2)
	if len(parts) < 2 {
		return
	}

	lines := bytes.Split(parts[0], utils.Del.NewLine)
	if len(lines) == 0 {
		return
	}

	requestLine := bytes.Split(lines[0], utils.Del.Space)
	if len(requestLine) != 3 {
		return
	}

	var params [][2][]byte
	queryParam := bytes.Split(requestLine[1], utils.Del.Question)
	for _, param := range bytes.Split(queryParam[1], utils.Del.Ampersand) {
		kv := bytes.SplitN(param, utils.Del.Equal, 2)
		params = append(params, [2][]byte{bytes.TrimSpace(kv[0]), bytes.TrimSpace(kv[1])})
	}

	var headers [][2][]byte
	for _, header := range lines[1:] {
		kv := bytes.SplitN(header, utils.Del.Colon, 2)
		headers = append(headers, [2][]byte{bytes.TrimSpace(kv[0]), bytes.TrimSpace(kv[1])})
	}

	r.Scope = Scope{
		Method:     bytes.TrimSpace(requestLine[0]),
		Endpoint:   bytes.TrimSpace(queryParam[0]),
		QueryParam: params,
		Version:    bytes.TrimSpace(bytes.SplitN(requestLine[2], utils.Del.Solidus, 2)[1]),
		Headers:    headers,
		Body:       parts[1],
	}
}
