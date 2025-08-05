package types

type App interface {
	RequestHandler(request Request)
}
