package types

type Request interface {
	Parse(request []byte)
}
