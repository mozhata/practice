package mmux_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"practice/go/mmux"
	"testing"

	"strings"

	"github.com/mozhata/handy"
	"github.com/smartystreets/goconvey/convey"
)

var (
	static5  = "/abc/bcd/cde/efg/fgh"
	dynamic5 = "/dy1/:id1/dy2/:id2/dy3"

	static9  = "/st1/st2/st3/st4/st5/st6/st7/st8/st9/st10/"
	dynamic9 = "/dy1/:id1/st3/:id2/st5/:st6/st7/:st8/st9/"

	mhandler = func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		fmt.Printf("the URL.Path: %q, params: %s\n", r.URL.Path, handy.MarshalJSONOrDie(params))
		fmt.Fprintf(w, "the URL.Path: %q, params: %s\n", r.URL.Path, handy.MarshalJSONOrDie(params))
	}

	mMux = mmux.New()
)

func TestMux(t *testing.T) {
	convey.Convey("testMux", t, func() {
		mMux := mmux.New()
		convey.Convey("test register static, and ServeHTTP", func() {
			mMux.Register(static5, "GET", mhandler)
			resp := lunchHTTPGET(mMux, static5, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(strings.Contains(resp.Body.String(), static5), convey.ShouldBeTrue)
			fmt.Println("resp.body: ", resp.Body)
		})
		convey.Convey("test register an invalid pattern", func() {
			var msg string
			invalidStaticPattern := "/a/b//c"
			convey.Convey("test `//`", func() {
				func() {
					defer func() {
						r := recover()
						msg = fmt.Sprintf("%v", r)
					}()
					mMux.Register(invalidStaticPattern, "GET", mhandler)
				}()
				convey.So(msg, convey.ShouldEqual, "pattern "+invalidStaticPattern+"is not vald")
			})
			convey.Convey("test `:`", func() {
				invalidStaticPattern = "/a/b/:/c"
				func() {
					defer func() {
						r := recover()
						msg = fmt.Sprintf("%v", r)
					}()
					mMux.Register(invalidStaticPattern, "GET", mhandler)
				}()
				convey.So(msg, convey.ShouldEqual, "pattern "+invalidStaticPattern+"is not vald")
			})
			convey.Convey("test emtpy", func() {
				invalidStaticPattern = ""
				func() {
					defer func() {
						r := recover()
						msg = fmt.Sprintf("%v", r)
					}()
					mMux.Register(invalidStaticPattern, "GET", mhandler)
				}()
				convey.So(msg, convey.ShouldEqual, "emtpty pattern !")
			})
			convey.Convey("test invalid root", func() {
				invalidStaticPattern = "abc/abc/ad"
				func() {
					defer func() {
						r := recover()
						msg = fmt.Sprintf("%v", r)
					}()
					mMux.Register(invalidStaticPattern, "GET", mhandler)
				}()
				convey.So(msg, convey.ShouldEqual, "path must begin with '/'. pattern: 'abc/abc/ad'")
			})
		})
		convey.Convey("test register dynamic pattern and ServeHTTP", func() {
			mMux.Register(dynamic5, "GET", mhandler)
			dynamic51 := "/dy1/this-is-id1/dy2/id2/dy3"
			resp := lunchHTTPGET(mMux, dynamic51, nil)
			convey.So(strings.Contains(resp.Body.String(), dynamic51), convey.ShouldBeTrue)
			convey.So(strings.Contains(resp.Body.String(), `"id1":"this-is-id1"`), convey.ShouldBeTrue)
			convey.So(strings.Contains(resp.Body.String(), `"id2":"id2"`), convey.ShouldBeTrue)
			dynamic52 := "/dy1/id1/dy2/this-is-id2/dy3"
			resp = lunchHTTPGET(mMux, dynamic52, nil)
			convey.So(strings.Contains(resp.Body.String(), dynamic52), convey.ShouldBeTrue)
			convey.So(strings.Contains(resp.Body.String(), `"id1":"id1"`), convey.ShouldBeTrue)
			convey.So(strings.Contains(resp.Body.String(), `"id2":"this-is-id2"`), convey.ShouldBeTrue)
		})
		convey.Convey("test register conflict pattern", func() {
			convey.Convey("static conflict", func() {
				var msg string
				staticPattern := "/a/b/c/static"
				mMux.Register(staticPattern, "GET", mhandler)
				func() {
					defer func() {
						r := recover()
						msg = fmt.Sprintf("%v", r)
					}()
					mMux.Register(staticPattern, "GET", mhandler)
				}()
				convey.So(msg, convey.ShouldEqual, "pattern "+staticPattern+"/ already registered.")
			})
			convey.Convey("dynamic conflict", func() {
				var msg string
				pattern1 := "/a/:b/:c/dynamic"
				pattern2 := "/a/:bb/:cc/dynamic"
				mMux.Register(pattern1, "GET", mhandler)
				func() {
					defer func() {
						r := recover()
						msg = fmt.Sprintf("%v", r)
					}()
					mMux.Register(pattern2, "GET", mhandler)
				}()
				convey.So(msg, convey.ShouldEqual, fmt.Sprintf("pattern %s/ conflict with %s/", pattern2, pattern1))
			})
		})
		convey.Convey("static pattern priority to dynamic", func() {

		})
	})
}

func lunchHTTPGET(h http.Handler, path string, params map[string]string) *httptest.ResponseRecorder {
	urlStr := buildUrl(path, params)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	return w
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
