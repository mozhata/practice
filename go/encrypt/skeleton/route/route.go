package route

import (
	"net/http"

	"github.com/beego/mux"
	"github.com/urfave/negroni"
)

type Route struct {
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

func NewRoute(pattern, method string, handler http.HandlerFunc) *Route {
	return &Route{
		Pattern: pattern,
		Method:  method,
		Handler: handler,
	}
}

func BuildHandler(routeLists ...[]*Route) http.Handler {
	router := mux.New()

	for _, routes := range routeLists {
		for _, rou := range routes {
			router.Handle(rou.Method, rou.Pattern, rou.Handler)
		}
	}

	// TODO: add serverfile route

	// use middleware
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}
