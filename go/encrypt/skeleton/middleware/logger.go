package middleware

// 记录request, response, err info

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httputil"
	"practice/go/encrypt/merr"
	"practice/go/encrypt/skeleton/reply"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/golang/glog"

	"git.meiqia.com/triones/compass/json"
)

const (
	AccessLogKey = "access_log"
)

type AccessLog struct {
	Method     string `json:"method"`
	Path       string `json:"path"`
	ClientIP   string `json:"client_ip"`
	Request    string `json:"request,omitempty"`
	Cost       string `json:"cost"`
	StatusCode int    `json:"status_code"`
	Response   string `json:"response,omitempty"`
}

func Logger(next http.Handler) http.Handler {
	return loggerMiddleware(next)
}

func loggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		buf := new(bytes.Buffer)
		ww.Tee(buf)

		dumpRequest, err := httputil.DumpRequest(r, true)
		if err != nil {
			err = merr.WrapErr(err, "dump request failed: %v", err)
			// log err
			glog.Errorln(merr.ErrDetail(err))
			// resp err
			resp := reply.Response{
				Code: merr.InternalError,
				Msg:  err.Error(),
			}
			replyer := reply.JSON(http.StatusInternalServerError, resp)
			replyer(w, r)
			return
		}

		start := time.Now()
		al := &AccessLog{
			Method:   r.Method,
			Path:     r.RequestURI,
			ClientIp: r.RemoteAddr,
			Request:  string(dumpRequest),
		}

		defer func() {
			al.Cost = time.Now().Sub(start).String()
			al.StatusCode = int(ww.Status())
			al.Response = string(buf.Bytes())

			// glog.Infof("")
			alb, err := json.Marshal(al)
			if err != nil {
				// log.Error("loggerMiddleware failed", zap.Error(err))
				return
			}
			// log.Info(string(alb))
		}()

		ctx := context.WithValue(r.Context(), AccessLogKey, al)
		next.ServeHTTP(ww, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func GetAccessLog(r *http.Request) *AccessLog {
	return r.Context().Value(AccessLogKey).(*AccessLog)
}
