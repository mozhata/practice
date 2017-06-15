package mmux_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"practice/go/mmux"
	"testing"

	"github.com/mozhata/handy"
)

var (
	static5  = "/abc/bcd/cde/efg/fgh"
	dynamic5 = "/dy1/:id1/dy2/:id2/dy3"

	static9  = "/st1/st2/st3/st4/st5/st6/st7/st8/st9/st10/"
	dynamic9 = "/dy1/:id1/st3/:id2/st5/:st6/st7/:st8/st9/"

	mhandler = func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		fmt.Printf("resp: the URL.Path: %q =====> params: %s\n", r.URL.Path, handy.MarshalJSONOrDie(params))
	}

	mMux = mmux.New()
)

func TestRegister(t *testing.T) {
	mMux.Register(static5, "GET", mhandler)
	lunchHTTPGET(mMux, static5, nil)

	dynamic51 := "/dy1/this-is-id1/dy2/id2/dy3"
	dynamic52 := "/dy1/id1/dy2/this-is-id2/dy3"
	dynamic53 := "/dy1/id1//:dy2/this-is-id2/dy3"
	dynamic54 := "/dy1/id1//:/this-is-id2/dy3"
	dynamic55 := "/dy1/id1//id/this-is-id2/dy3"
	mMux.Register(dynamic5, "GET", mhandler)
	mMux.Register(dynamic53, "GET", mhandler)
	mMux.Register(dynamic54, "GET", mhandler)
	mMux.Register(dynamic55, "GET", mhandler)
	lunchHTTPGET(mMux, dynamic51, nil)
	lunchHTTPGET(mMux, dynamic52, nil)
}

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
