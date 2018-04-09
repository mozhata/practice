package reply

import (
	"encoding/json"
	"net/http"

	"practice/go/encrypt/merr"
)

/*
TODO:
错误码-msg对照表
中英文错误
*/
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Body interface{} `json:"body"`
}

func Wrap(f func(w http.ResponseWriter, r *http.Request) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responser := f(w, r)
		responser(w, r)
	}
}

func Success(content interface{}) http.HandlerFunc {
	return JSON(http.StatusOK, Response{
		Code: merr.OK,
		Body: content,
	})
}

//TODO: 创建errorcode=> msg 中英文映射,
func Err(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e, ok := err.(*merr.MErr)
		if !ok {
			e = merr.WrapErr(err)
		}
		rp := JSON(http.StatusBadRequest, Response{
			Code: e.Code,
		})
		rp(w, r)
	}
}

// func JSON(code, statusCode int, msg string, content interface{}) http.HandlerFunc {
func JSON(statusCode int, resp Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(statusCode)
		if resp.Msg == "" {
			// TODO: language
			resp.Msg = merr.GetMsg(resp.Code, []string{})
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}
}
