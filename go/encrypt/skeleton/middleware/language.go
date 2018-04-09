package middleware

import (
	"context"
	"net/http"
	"practice/go/encrypt/skeleton/input"
)

// SetCtxLanguage 从HTTP Header中取出languages，将其设置到Context中
func SetCtxLanguage(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		languages := input.FromHttpHeader(r.Header)
		ctx := context.WithValue(r.Context(), input.LanguageKey, languages)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
