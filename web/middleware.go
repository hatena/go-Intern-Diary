package web

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hatena/go-Intern-Diary/loader"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{w, http.StatusOK}
		next.ServeHTTP(lrw, r)
		log.Printf("%s %s took %.2fmsec and returned %d %s",
			r.Method, r.URL.Path, float64(time.Now().Sub(start).Nanoseconds())/1e6,
			lrw.statusCode, http.StatusText(lrw.statusCode),
		)
	})
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Oprions", "DENY")
		next.ServeHTTP(w, r)
	})
}

func (s *server) resolveUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", user)))
	})
}

func (s *server) attachLoaderMiddleware(next http.Handler) http.Handler {
	loaders := loader.New(s.app)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(loaders.Attach(r.Context())))
	})
}
