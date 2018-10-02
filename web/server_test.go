package web

import (
	"net/http"
)

func init() {
	csrfMiddleware = func(next http.Handler) http.Handler {
		return next
	}
	csrfToken = func(r *http.Request) string {
		return ""
	}
}
