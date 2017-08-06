package util

import (
	"fmt"
	"net/http"
	"runtime"
)

// BaseErr basic error class
type BaseErr struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg"`
	stackTrace []byte
}

func (e *BaseErr) Error() string {
	return fmt.Sprintf("%v: %v", e.StatusCode, e.Message)
}

func (e *BaseErr) Stack() string {
	return string(e.stackTrace)
}

func NotFoundError(fmtAndArgs ...interface{}) error {
	return wrapErr(http.StatusNotFound, fmtAndArgs...)
}

func InvalidArgumentErr(fmtAndArgs ...interface{}) error {
	return wrapErr(http.StatusBadRequest, fmtAndArgs...)
}

func ForbiddenError(fmtAndArgs ...interface{}) error {
	// return &BaseErr{http.StatusForbidden, BuildErrMsg(fmtAndArgs...)}
	return wrapErr(http.StatusForbidden, fmtAndArgs...)
}

func InternalError(fmtAndArgs ...interface{}) error {
	// return &BaseErr{http.StatusInternalServerError, BuildErrMsg(fmtAndArgs...)}
	return wrapErr(http.StatusInternalServerError, fmtAndArgs...)
}

func WrapErr(err error) *BaseErr {
	return wrapErr(http.StatusInternalServerError, err)
}

func wrapErr(code int, fmtAndArgs ...interface{}) *BaseErr {
	const size = 1 << 12
	buf := make([]byte, size)
	n := runtime.Stack(buf, false)
	err := &BaseErr{
		StatusCode: code,
		stackTrace: buf[:n],
	}
	if len(fmtAndArgs) == 0 || fmtAndArgs == nil {
		return err
	}
	for _, arg := range fmtAndArgs {
		if e, ok := arg.(*BaseErr); ok {
			return e
		}
	}
	err.Message = BuildErrMsg(fmtAndArgs...)
	return err
}
