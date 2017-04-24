package kmux

import "net/http"

func (k *Mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer k.recv(w, req)
	}

	path := req.URL.Path
	if root := k.trees[req.Method]; root != nil {

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
