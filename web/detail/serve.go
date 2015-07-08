package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8080", mux)
}

// var mux map[string]func(http.ResponseWriter, *http.Request)

// func main() {
// 	server := http.Server{
// 		Addr:        ":8080",
// 		Handler:     &myHandler{},
// 		ReadTimeout: 4 * time.Second,
// 	}

// 	mux = make(map[string]func(http.ResponseWriter, *http.Request))
// 	mux["/hello"] = sayHello
// 	mux["/bye"] = sayBye

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if h, ok := mux[r.URL.String()]; ok {
	// 	h(w, r)
	// 	return
	// }

	io.WriteString(w, "URL: "+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world, this is version3")
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye bye this is version6")
}
