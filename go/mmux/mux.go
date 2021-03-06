package mmux

import "net/http"

type Mux struct {
	routes map[string]*route

	// Configurable http.Handler which is called when no matching route is
	// found. If it is not set, http.NotFound is used.
	NotFound http.Handler

	// Configurable http.Handler which is called when a request
	// cannot be routed and HandleMethodNotAllowed is true.
	// If it is not set, http.Error with http.StatusMethodNotAllowed is used.
	// The "Allow" header with allowed request methods is set before the handler
	// is called.
	MethodNotAllowed http.Handler

	// Function to handle panics recovered from http handlers.
	// It should be used to generate a error page and return the http error code
	// 500 (Internal Server Error).
	// The handler can be used to keep your server from crashing because of
	// unrecovered panics.
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

func New() *Mux {
	routes := make(map[string]*route)
	return &Mux{routes: routes}
}

// Make sure the Mux conforms with the http.Handler interface
var _ http.Handler = New()

func (m *Mux) Register(pattern, method string, handle Handle) {
	root := m.routes[method]
	if root == nil {
		root = newRoute()
		m.routes[method] = root
	}

	root.add(pattern, handle)
}

func (k *Mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if k.PanicHandler != nil {
		defer k.recv(w, req)
	}

	path := unifyPath(req.URL.Path)
	if root := k.routes[req.Method]; root != nil {
		if handle, ps := root.getHandle(path); handle != nil {
			handle(w, req, ps)
			return
		}
	}
	if k.NotFound != nil {
		k.NotFound.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func (k *Mux) recv(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		k.PanicHandler(w, req, rcv)
	}
}

func unifyPath(path string) string {
	if path[len(path)-1] != '/' {
		path += "/"
	}
	return path
}
