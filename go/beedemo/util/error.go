package util

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

// BaseErr basic error class
type BaseErr struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg"`
	stackPC    []uintptr
	originErr  error
}

func (e BaseErr) OriginErr() string {
	if e.originErr == nil {
		return "null"
	}
	return e.originErr.Error()
}

func (e BaseErr) CallStack() string {
	frames := runtime.CallersFrames(e.stackPC)
	var (
		f      runtime.Frame
		more   bool
		result string
		index  int
	)
	for {
		f, more = frames.Next()
		if index = strings.Index(f.File, "src"); index != -1 {
			// trim GOPATH or GOROOT prifix
			f.File = string(f.File[index+4:])
		}
		result = fmt.Sprintf("%s%s\n\t%s:%d\n", result, f.Function, f.File, f.Line)
		if !more {
			break
		}
	}
	return result
}

func (e *BaseErr) Error() string {
	return fmt.Sprintf("%v: %v", e.StatusCode, e.Message)
}

func NotFoundError(err error, fmtAndArgs ...interface{}) error {
	return wrapErr(err, http.StatusNotFound, fmtAndArgs...)
}

func InvalidArgumentErr(err error, fmtAndArgs ...interface{}) error {
	return wrapErr(err, http.StatusBadRequest, fmtAndArgs...)
}

func ForbiddenError(err error, fmtAndArgs ...interface{}) error {
	return wrapErr(err, http.StatusForbidden, fmtAndArgs...)
}

func InternalError(err error, fmtAndArgs ...interface{}) error {
	return wrapErr(err, http.StatusInternalServerError, fmtAndArgs...)
}

func WrapErr(err error) *BaseErr {
	return wrapErr(err, http.StatusInternalServerError)
}

func wrapErr(err error, code int, fmtAndArgs ...interface{}) *BaseErr {
	if e, ok := err.(*BaseErr); ok {
		return e
	}

	pcs := make([]uintptr, 32)
	count := runtime.Callers(3, pcs)
	e := &BaseErr{
		StatusCode: code,
		stackPC:    pcs[:count],
		originErr:  err,
	}
	e.Message = BuildErrMsg(fmtAndArgs...)
	if e.Message == "" {
		e.Message = e.OriginErr()
	}
	return e
}

func formatErr(err error) string {
	if err == nil {
		return "null"
	}
	return err.Error()
}
