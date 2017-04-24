package kmux

import (
	"net/http"
)

type Mux struct {
	trees map[string]*route

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
	RedirectTrailingSlash bool

	// If enabled, the router tries to fix the current request path, if no
	// handle is registered for it.
	// First superfluous path elements like ../ or // are removed.
	// Afterwards the router does a case-insensitive lookup of the cleaned path.
	// If a handle can be found for this route, the router makes a redirection
	// to the corrected path with status code 301 for GET requests and 307 for
	// all other request methods.
	// For example /FOO and /..//Foo could be redirected to /foo.
	// RedirectTrailingSlash is independent of this option.
	RedirectFixedPath bool

	// If enabled, the router checks if another method is allowed for the
	// current route, if the current request can not be routed.
	// If this is the case, the request is answered with 'Method Not Allowed'
	// and HTTP status code 405.
	// If no other Method is allowed, the request is delegated to the NotFound
	// handler.
	HandleMethodNotAllowed bool

	// If enabled, the router automatically replies to OPTIONS requests.
	// Custom OPTIONS handlers take priority over automatic replies.
	HandleOPTIONS bool

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
	return &Mux{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}
}

// Make sure the Mux conforms with the http.Handler interface
var _ http.Handler = New()

func (m *Mux) Register(path, method string, handle Handle) {
	if m.trees == nil {
		m.trees = make(map[string]*route)
	}

	root := m.trees[method]
	if root == nil {
		root = newRoute()
		m.trees[method] = root
	}

	root.add(validatePattern(path), handle)
}

func (k *Mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if k.PanicHandler != nil {
		defer k.recv(w, req)
	}

	path := unifyPath(req.URL.Path)
	if root := k.trees[req.Method]; root != nil {
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

func validatePattern(pattern string) string {
	if pattern[0] != '/' {
		panic("path must begin with '/' in path '" + pattern + "'")
	}

	return unifyPath(pattern)
}
