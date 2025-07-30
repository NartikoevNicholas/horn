package horn

type Middleware interface{}
type Handler interface{}
type App interface{
	Handlers() []Handler
	Middleware() []Middleware
}
