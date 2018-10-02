package web

//go:generate go-assets-builder --package=web --output=./templates-gen.go --strip-prefix="/templates/" --variable=Templates ../templates

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/justinas/nosurf"

	"github.com/hatena/go-Intern-Diary/service"
)

type Server interface {
	Handler() http.Handler
}

const sessionKey = "DIARY_SESSION"

func NewServer(app service.DiaryApp) Server {
	return &server{app: app}
}

type server struct {
	app service.DiaryApp
}

func (s *server) Handler() http.Handler {
	router := httptreemux.New()

	// ...

	return router
}

var csrfMiddleware = func(next http.Handler) http.Handler {
	return nosurf.New(next)
}

var csrfToken = func(r *http.Request) string {
	return nosurf.Token(r)
}
