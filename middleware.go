package components

type ScopedMiddleware struct {
	Middleware MiddlewareProvider
	Only       []string
	Except     []string
	Global     bool
}

type Middlewares []ScopedMiddleware

type MiddlewareProvider interface {
	InitMiddleware(u *Utils)
	New(c State, next func(State) error) error
}

type Middleware struct {
	*Utils
	onInit func()
}

func (m *Middleware) InitMiddleware(u *Utils) {
	m.Utils = u
	if m.onInit != nil {
		m.onInit()
	}
}

func (m *Middleware) OnInit(callback func()) {
	m.onInit = callback
}
