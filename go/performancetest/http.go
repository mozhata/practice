package performancetest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/glog"
	"github.com/smartystreets/goconvey/convey"
)

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

func lunchHTTPGET(h http.Handler, path string, params map[string]string) {
	urlStr := buildUrl(path, params)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
}

func HTTPTest(t *testing.T, h http.Handler, method, path string, params map[string]string, body *bytes.Buffer, contentType string, result interface{}) {
	var w *httptest.ResponseRecorder
	url := buildUrl(path, params)
	convey.Convey("test http", t, func() {
		for i := 0; i < 10; i++ {
			req, err := http.NewRequest(method, url, body)
			req.Header.Set("Content-Type", contentType)
			convey.So(err, convey.ShouldBeNil)
			w = httptest.NewRecorder()
			h.ServeHTTP(w, req)
			if w.Code == 301 {
				path = w.Header().Get("Location")
				glog.Infof("http test path redirct to %s", path)
				continue
			}
			if result != nil {
				err := json.Unmarshal(w.Body.Bytes(), result)
				convey.So(err, convey.ShouldBeNil)
			}
		}
	})
}
