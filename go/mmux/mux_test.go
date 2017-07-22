package mmux_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"practice/go/mmux"
	"testing"

	"github.com/mozhata/handy"
	"github.com/smartystreets/goconvey/convey"
)

var (
	p  = fmt.Println
	pf = fmt.Printf
)

func TestMux(t *testing.T) {

	mhandler := func(w http.ResponseWriter, r *http.Request, params mmux.PathVars) {
		fmt.Printf("!!the URL.Path: %q, params: %s\n", r.URL.Path, handy.MarshalJSONOrDie(params))
		fmt.Fprintf(w, "!!!!the URL.Path: %q, params: %s\n", r.URL.Path, handy.MarshalJSONOrDie(params))
	}
	convey.Convey("testMux", t, func() {
		mMux := mmux.New()
		/*convey.Convey("test register static, and ServeHTTP", func() {
			static := "/abc/bcd/cde/efg/fgh"
			mMux.Register(static, "GET", mhandler)
			resp := lunchHTTPGET(mMux, static, nil)
			convey.So(resp, convey.ShouldNotBeNil)
			convey.So(strings.Contains(resp.Body.String(), static), convey.ShouldBeTrue)
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
				convey.So(msg, convey.ShouldEqual, fmt.Sprintf("pattern %q is not valid", unifyPath(invalidStaticPattern)))
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
				convey.So(msg, convey.ShouldEqual, fmt.Sprintf("pattern %q is not valid", unifyPath(invalidStaticPattern)))
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
				convey.So(msg, convey.ShouldEqual, fmt.Sprintf("path must begin with '/'. pattern: '%s'", invalidStaticPattern))
			})
		})
		convey.Convey("test register dynamic pattern and ServeHTTP", func() {
			dynamic := "/dy1/:id1/dy2/:id2/dy3"
			mMux.Register(dynamic, "GET", mhandler)
			dynamicInstance := "/dy1/this-is-id1/dy2/id2/dy3"
			resp := lunchHTTPGET(mMux, dynamicInstance, nil)
			convey.So(strings.Contains(resp.Body.String(), dynamicInstance), convey.ShouldBeTrue)
			pf("\n\n~~dynamic pattern: %q\n~~url: %q\n~~resp body:\n%s\n\n", dynamic, dynamicInstance, resp.Body.String())
			convey.So(strings.Contains(resp.Body.String(), dynamicInstance), convey.ShouldBeTrue)
			convey.So(strings.Contains(resp.Body.String(), ` params: [{"Key":"id1","Value":"this-is-id1"},{"Key":"id2","Value":"id2"}]`), convey.ShouldBeTrue)
			dynamic52 := "/dy1/id1/dy2/this-is-id2/dy3"
			resp = lunchHTTPGET(mMux, dynamic52, nil)
			convey.So(strings.Contains(resp.Body.String(), dynamic52), convey.ShouldBeTrue)
			convey.So(strings.Contains(resp.Body.String(), dynamic52), convey.ShouldBeTrue)
			convey.So(strings.Contains(resp.Body.String(), `params: [{"Key":"id1","Value":"id1"},{"Key":"id2","Value":"this-is-id2"}]`), convey.ShouldBeTrue)
		})*/
		convey.Convey("test register conflict pattern", func() {
			convey.Convey("static same", func() {
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
				convey.So(msg, convey.ShouldEqual, "pattern "+staticPattern+"/ conflict with "+unifyPath(staticPattern))
			})
			convey.Convey("dynamic conflict", func() {
				var msg string
				pattern1 := "/a/:b/:c/dynamic"
				pattern2 := "/a/:bb/:cc/dynamic"
				pattern3 := "/a/:bb/ss/dynamic"
				mMux.Register(pattern1, "GET", mhandler)
				func() {
					defer func() {
						r := recover()
						msg = fmt.Sprintf("%v", r)
					}()
					mMux.Register(pattern2, "GET", mhandler)
				}()
				convey.So(msg, convey.ShouldEqual, fmt.Sprintf("pattern %s/ conflict with %s/", pattern2, pattern1))
				func() {
					defer func() {
						r := recover()
						msg = fmt.Sprintf("%v", r)
					}()
					mMux.Register(pattern3, "GET", mhandler)
				}()
				convey.So(msg, convey.ShouldEqual, fmt.Sprintf("pattern %s/ conflict with %s/", pattern3, pattern1))
			})
		})
		/*convey.Convey("static pattern prior to dynamic", func() {
			static := "/this/is/static/pattern"
			dynamic := "/this/is/:dynamic/pattern"
			mMux.Register(static, "GET", mhandler)
			mMux.Register(dynamic, "GET", mhandler)
			resp := lunchHTTPGET(mMux, static, nil)
			convey.So(strings.Contains(resp.Body.String(), static), convey.ShouldBeTrue)
			dynamicInstance := "/this/is/test/pattern"
			resp = lunchHTTPGET(mMux, dynamicInstance, nil)
			convey.So(strings.Contains(resp.Body.String(), static), convey.ShouldBeFalse)
			// convey.So(strings.Contains(resp.Body.String(), `{"dynamic":"test"}`), convey.ShouldBeTrue)
			// dynamic2 := "/topic/:topicID/reply/add"
			// dynamic3 := "/topic/:topicID/reply/:topicID"
		})*/
		/*convey.Convey("test match root", func() {
			urlPath := "/"
			resp := lunchHTTPGET(mMux, urlPath, nil)
			convey.So(resp.Body.String(), convey.ShouldEqual, "404 page not found\n")
		})*/

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

func unifyPath(path string) string {
	if path[len(path)-1] != '/' {
		path += "/"
	}
	return path
}
