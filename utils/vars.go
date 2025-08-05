package utils

type Delimiters struct {
	Body      []byte
	Space     []byte
	Equal     []byte
	Colon     []byte
	NewLine   []byte
	Solidus   []byte
	Question  []byte
	Ampersand []byte
}

var Del = Delimiters{
	Body:      []byte("\r\n\r\n"),
	Space:     []byte(" "),
	Equal:     []byte("="),
	Colon:     []byte(":"),
	NewLine:   []byte("\r\n"),
	Solidus:   []byte("/"),
	Question:  []byte("?"),
	Ampersand: []byte("&"),
}

var SupportedProtocols = map[string]struct{}{
	"http": {},
}
