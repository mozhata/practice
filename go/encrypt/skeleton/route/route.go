package route

import (
	"net/http"

	"practice/go/encrypt/skeleton/context"

	"github.com/beego/mux"
	"github.com/urfave/negroni"
)

type Route struct {
	Pattern string
	Method  string
	Handle  context.ProcessRequest
}

func NewRoute(pattern, method string, handle context.ProcessRequest) *Route {
	return &Route{
		Pattern: pattern,
		Method:  method,
		Handle:  handle,
	}
}

func BuildHandler(routeLists ...[]*Route) http.Handler {
	router := mux.New()

	for _, routes := range routeLists {
		for _, rou := range routes {
			handler := func(route *Route) func(w http.ResponseWriter, r *http.Request) {
				return func(w http.ResponseWriter, r *http.Request) {
					ctx := &context.Context{
						Input: context.NewParam(r),
						Resp:  context.NewResponse(w),
					}
					replyer := route.Handle(ctx)
					ctx.Resp.ReplyFunc = replyer
					ctx.Reply()
				}
			}(rou)

			router.Handle(rou.Method, rou.Pattern, handler)
		}
	}

	// TODO: add serverfile route

	// use middleware
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}
