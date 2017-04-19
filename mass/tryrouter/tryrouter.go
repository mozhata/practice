package tryrouter

import (
	"fmt"
	"net/http"

	goji "goji.io"
	"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func TryGoji() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)

	fmt.Println("localhost:8989")
	http.ListenAndServe("localhost:8989", mux)
}
