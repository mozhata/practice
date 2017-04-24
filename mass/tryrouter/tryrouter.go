package tryrouter

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"practice/go/forkrouter"
	"practice/go/kmux"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	mux9 "github.com/issue9/mux"
	goji "goji.io"
	"goji.io/pat"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func getIP(w http.ResponseWriter, r *http.Request) {
	ip, err := ParseIP(r)
	check(err)
	fmt.Fprintf(w, "the ip: %s\n", ip)
}

// ParseIP extracts the user IP address from req, if present.
func ParseIP(req *http.Request) (net.IP, error) {
	glog.Infof("req.RemoteAddr: %s", req.RemoteAddr)
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil
}

func TryGoji() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	mux.HandleFunc(pat.Get("/topic/add"), hello)
	mux.HandleFunc(pat.Get("/topic/:id/reply/add"), hello)

	fmt.Println("localhost:8989")
	http.ListenAndServe(":8989", mux)
}

func TryIssue9Mux() {
	mux := mux9.NewServeMux(false)
	mux.GetFunc("/", hello)
	mux.GetFunc("/topic/add", hello)
	mux.GetFunc("/topic/:id/reply/add", hello)

	fmt.Println("localhost:8989")
	http.ListenAndServe(":8989", mux)
}

func TrygorillaMux() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/ip", getIP)
	r.HandleFunc("/topic/add", hello)
	r.HandleFunc("/topic/:id/reply/add", hello)

	fmt.Println("localhost:8989")
	http.ListenAndServe(":8989", r)
}

func TryForkRouter() {
	router := forkrouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
func Index(w http.ResponseWriter, r *http.Request, _ forkrouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func TryKmux() {
	router := kmux.New()
	router.Register("/", "GET", testWork)

	port := ":8080"
	glog.Infof("serving at %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func testWork(w http.ResponseWriter, r *http.Request, ps kmux.Params) {
	fmt.Fprintf(w, "ps: %#v\nheader: %v\n", ps, r.Header)
}
