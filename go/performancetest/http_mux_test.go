package performancetest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"practice/go/forkrouter"
	"practice/go/kmux"
	"testing"

	"goji.io/pat"

	goji "goji.io"

	"github.com/gorilla/mux"
	mux9 "github.com/issue9/mux"
)

var (
	static5  = "/abc/bcd/cde/efg/fgh"
	dynamic5 = "/dy1/:id1/dy2/:id2/dy3"

	static9  = "/st1/st2/st3/st4/st5/st6/st7/st8/st9/st10/"
	dynamic9 = "/dy1/:id1/st3/:id2/st5/:st6/st7/:st8/st9/"

	printHeader = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world\nthe header is %v\n", r.Header)
	}
	stdHandler = func(w http.ResponseWriter, r *http.Request) {
		printHeader(w, r)
	}
	httprouterHandler = func(w http.ResponseWriter, r *http.Request, ps forkrouter.Params) {
		printHeader(w, r)
	}
	khandler = func(w http.ResponseWriter, r *http.Request, ps kmux.Params) {
		printHeader(w, r)
	}

	gorillaMux = func() *mux.Router {
		gomux := mux.NewRouter()
		gomux.Methods("GET").Path(static5).HandlerFunc(stdHandler)
		gomux.Methods("GET").Path(dynamic5).HandlerFunc(stdHandler)
		gomux.Methods("GET").Path(static9).HandlerFunc(stdHandler)
		gomux.Methods("GET").Path(dynamic9).HandlerFunc(stdHandler)
		return gomux
	}()
	gojiMux = func() *goji.Mux {
		gjmux := goji.NewMux()
		gjmux.HandleFunc(pat.Get(static5), stdHandler)
		gjmux.HandleFunc(pat.Get(dynamic5), stdHandler)
		gjmux.HandleFunc(pat.Get(static9), stdHandler)
		gjmux.HandleFunc(pat.Get(dynamic9), stdHandler)
		return gjmux
	}()
	issue9Mux = func() *mux9.ServeMux {
		i9mux := mux9.NewServeMux(false)
		i9mux.GetFunc(static5, stdHandler)
		i9mux.GetFunc(dynamic5, stdHandler)
		i9mux.GetFunc(static9, stdHandler)
		i9mux.GetFunc(dynamic9, stdHandler)
		return i9mux
	}()
	forkMux = func() *forkrouter.Router {
		fomux := forkrouter.New()
		fomux.GET(static5, httprouterHandler)
		fomux.GET(dynamic5, httprouterHandler)
		fomux.GET(static9, httprouterHandler)
		fomux.GET(dynamic9, httprouterHandler)
		return fomux
	}()
	kMux = func() *kmux.Mux {
		kkk := kmux.New()
		kkk.Register(static5, "GET", khandler)
		kkk.Register(dynamic5, "GET", khandler)
		kkk.Register(static9, "GET", khandler)
		kkk.Register(dynamic9, "GET", khandler)
		return kkk
	}()
)

// func init() {
// 	flag.Lookup("logtostderr").Value.Set("true")
// 	flag.Parse()
// }

func lunchHTTPGET(h http.Handler, path string, params map[string]string) {
	urlStr := buildUrl(path, params)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
}
func buildUrl(path string, params map[string]string) string {
	ret := path
	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}
	if len(values) > 0 {
		ret += "?" + values.Encode()
	}
	return ret
}

func BenchmarkGorillarMuxSt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(gorillaMux, static5, nil)
	}
}
func BenchmarkGojiMuxSt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(gojiMux, static5, nil)
	}
}
func BenchmarkIssue9MuxMuxSt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(issue9Mux, static5, nil)
	}
}
func BenchmarkForkMuxMuxSt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(forkMux, static5, nil)
	}
}
func BenchmarkKangMuxSt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(kMux, static5, nil)
	}
}

func BenchmarkGorillarMuxDy5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(gorillaMux, dynamic5, nil)
	}
}
func BenchmarkGojiMuxDy5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(gojiMux, dynamic5, nil)
	}
}
func BenchmarkIssue9MuxMuxDy5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(issue9Mux, dynamic5, nil)
	}
}
func BenchmarkForkMuxMuxDy5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(forkMux, dynamic5, nil)
	}
}
func BenchmarkKangMuxDy5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(kMux, dynamic5, nil)
	}
}

func BenchmarkGorillarMuxSt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(gorillaMux, static9, nil)
	}
}
func BenchmarkGojiMuxSt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(gojiMux, static9, nil)
	}
}
func BenchmarkIssue9MuxMuxSt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(issue9Mux, static9, nil)
	}
}
func BenchmarkForkMuxMuxSt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(forkMux, static9, nil)
	}
}
func BenchmarkKangMuxSt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(kMux, static9, nil)
	}
}

func BenchmarkGorillarMuxDy9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(gorillaMux, dynamic9, nil)
	}
}
func BenchmarkGojiMuxDy9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(gojiMux, dynamic9, nil)
	}
}
func BenchmarkIssue9MuxMuxDy9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(issue9Mux, dynamic9, nil)
	}
}
func BenchmarkForkMuxMuxDy9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(forkMux, dynamic9, nil)
	}
}
func BenchmarkKangMuxDy9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lunchHTTPGET(kMux, dynamic9, nil)
	}
}
