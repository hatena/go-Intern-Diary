package web

//go:generate go-assets-builder --package=web --output=./templates-gen.go --strip-prefix="/templates/" --variable=Templates ../templates

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/justinas/nosurf"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/resolver"
	"github.com/hatena/go-Intern-Diary/service"
)

type Server interface {
	Handler() http.Handler
}

const sessionKey = "DIARY_SESSION"

var templates map[string]*template.Template

func init() {
	var err error
	templates, err = loadTemplates()
	if err != nil {
		panic(err)
	}
}

func loadTemplates() (map[string]*template.Template, error) {
	templates := make(map[string]*template.Template)
	bs, err := ioutil.ReadAll(Templates.Files["main.tmpl"])
	if err != nil {
		return nil, err
	}
	mainTmpl := template.Must(template.New("main.tmpl").Parse(string(bs)))
	for fileName, file := range Templates.Files {
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		mainTmpl := template.Must(mainTmpl.Clone())
		templates[fileName] = template.Must(mainTmpl.New(fileName).Parse(string(bs)))
	}
	return templates, nil
}

func NewServer(app service.DiaryApp) Server {
	return &server{app: app}
}

type server struct {
	app service.DiaryApp
}

func (s *server) Handler() http.Handler {
	router := httptreemux.New()

	handle := func(method, path string, handler http.Handler) {
		router.UsingContext().Handler(method, path,
			csrfMiddleware(loggingMiddleware(headerMiddleware(handler))),
		)
	}

	handle("GET", "/", s.indexHandler())
	handle("GET", "/signup", s.willSignupHandler())
	handle("POST", "/signup", s.signupHandler())

	// handle("POST", "/signout", s.signoutHandler())
	handle("GET", "/signin", s.willSigninHandler())
	handle("POST", "/signin", s.signinHandler())

	handle("GET", "/diaries", s.diariesHandler())
	handle("GET", "/diaries/new", s.willAddDiaryHandler())
	handle("POST", "/diaries/new", s.addDiaryHandler())

	handle("GET", "/diaries/:id/articles", s.articlesHandler())
	handle("GET", "/diaries/:id/articles/new", s.willAddArticleHandler())
	handle("POST", "/diaries/:id/articles/new", s.addArticleHandler())
	handle("POST", "/diaries/:id/delete", s.deleteDiaryHandler())
	handle("GET", "/diaries/:diary_id/articles/:article_id", s.articleHandler())
	handle("POST", "/diaries/:diary_id/articles/:article_id/delete", s.deleteArticleHandler())

	handle("GET", "/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	handle("GET", "/graphiql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates["graphiql.tmpl"].ExecuteTemplate(w, "graphiql.tmpl", nil)
	}))

	handle("GET", "/spa/", s.spaHandler())
	handle("GET", "/spa/*", s.spaHandler())

	router.UsingContext().Handler("POST", "/query",
		s.attachLoaderMiddleware(
			s.resolveUserMiddleware(
				loggingMiddleware(headerMiddleware(resolver.NewHandler(s.app))))))

	router.UsingContext().Handler("POST", "/signout",
		loggingMiddleware(headerMiddleware(s.signoutHandler())))

	return router
}

var csrfMiddleware = func(next http.Handler) http.Handler {
	return nosurf.New(next)
}

var csrfToken = func(r *http.Request) string {
	return nosurf.Token(r)
}

func (s *server) findUser(r *http.Request) (user *model.User) {
	cookie, err := r.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ = s.app.FindUserByToken(cookie.Value)
	}
	return
}

func (s *server) renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["CSRFToken"] = csrfToken(r)
	err := templates[tmpl].ExecuteTemplate(w, "main.tmpl", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) getParams(r *http.Request, name string) string {
	return httptreemux.ContextParams(r.Context())[name]
}

func (s *server) indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		s.renderTemplate(w, r, "index.tmpl", map[string]interface{}{
			"User": user,
		})
	})
}

func (s *server) willSignupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.renderTemplate(w, r, "signup.tmpl", nil)
	})
}

func (s *server) signupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, password := r.FormValue("name"), r.FormValue("password")
		if err := s.app.CreateNewUser(name, password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user, err := s.app.FindUserByName(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		expiresAt := time.Now().Add(24 * time.Hour)
		token, err := s.app.CreateNewToken(user.ID, expiresAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    sessionKey,
			Value:   token,
			Expires: expiresAt,
		})
		http.Redirect(w, r, "/spa", http.StatusSeeOther)
	})
}

func (s *server) signoutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    sessionKey,
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}

func (s *server) willSigninHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.renderTemplate(w, r, "signin.tmpl", nil)
	})
}

func (s *server) signinHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, password := r.FormValue("name"), r.FormValue("password")
		if ok, err := s.app.LoginUser(name, password); err != nil || !ok {
			http.Error(w, "user not found or invalid password", http.StatusBadRequest)
			return
		}
		user, err := s.app.FindUserByName(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		expiresAt := time.Now().Add(24 * time.Hour)
		token, err := s.app.CreateNewToken(user.ID, expiresAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    sessionKey,
			Value:   token,
			Expires: expiresAt,
		})
		http.Redirect(w, r, "/spa", http.StatusSeeOther)
	})
}

func (s *server) diariesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		var page uint64 = 1
		var limit uint64 = 100 // todo
		diaries, err := s.app.ListDiariesByUserID(user.ID, page, limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.renderTemplate(w, r, "diaries.tmpl", map[string]interface{}{
			"User":    user,
			"Diaries": diaries,
		})
	})
}

func (s *server) willAddDiaryHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		s.renderTemplate(w, r, "add_diary.tmpl", map[string]interface{}{
			"User": user,
		})
	})
}

func (s *server) addDiaryHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Error(w, "please login", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		if _, err := s.app.CreateNewDiary(user.ID, name, []*model.TagWithCategory{}); err != nil {
			http.Error(w, "failed to create diary", http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/diaries", http.StatusSeeOther)
	})
}

func (s *server) deleteDiaryHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Error(w, "please login", http.StatusBadRequest)
			return
		}
		diaryID, err := strconv.ParseUint(s.getParams(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		if err := s.app.DeleteDiary(user.ID, diaryID); err != nil {
			http.Error(w, fmt.Sprintf("failed to delete diary: %+v", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/diaries", http.StatusSeeOther)
	})
}

func (s *server) articlesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Error(w, "please login", http.StatusBadRequest)
			return
		}
		diaryID, err := strconv.ParseUint(s.getParams(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		diary, err := s.app.FindDiaryByID(diaryID)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		var page int = 1
		var limit int = 100 // todo
		articles, pageInfo, err := s.app.ListArticlesByDiaryID(diaryID, page, limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.renderTemplate(w, r, "articles.tmpl", map[string]interface{}{
			"User":     user,
			"Diary":    diary,
			"Articles": articles,
			"pageInfo": pageInfo,
		})
	})
}

func (s *server) willAddArticleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Error(w, "please login", http.StatusBadRequest)
			return
		}
		diaryID, err := strconv.ParseUint(s.getParams(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		diary, err := s.app.FindDiaryByID(diaryID)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		s.renderTemplate(w, r, "add_article.tmpl", map[string]interface{}{
			"User":  user,
			"Diary": diary,
		})
	})
}

func (s *server) addArticleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Error(w, "please login", http.StatusBadRequest)
			return
		}
		diaryID, err := strconv.ParseUint(s.getParams(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		title, content := r.FormValue("title"), r.FormValue("content")
		if _, err := s.app.CreateNewArticle(diaryID, user.ID, title, content); err != nil {
			http.Error(w, "failed to create article", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/diaries/%d/articles", diaryID), http.StatusSeeOther)
	})
}

func (s *server) articleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Error(w, "please login", http.StatusBadRequest)
			return
		}
		diaryID, err := strconv.ParseUint(s.getParams(r, "diary_id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		articleID, err := strconv.ParseUint(s.getParams(r, "article_id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		diary, err := s.app.FindDiaryByID(diaryID)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		article, err := s.app.FindArticleByID(articleID, diaryID)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		s.renderTemplate(w, r, "article.tmpl", map[string]interface{}{
			"User":    user,
			"Diary":   diary,
			"Article": article,
		})
	})
}

func (s *server) deleteArticleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := s.findUser(r)
		if user == nil {
			http.Error(w, "please login", http.StatusBadRequest)
			return
		}
		diaryID, err := strconv.ParseUint(s.getParams(r, "diary_id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		articleID, err := strconv.ParseUint(s.getParams(r, "article_id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid diary id", http.StatusBadRequest)
			return
		}
		if err := s.app.DeleteArticle(articleID, user.ID); err != nil {
			http.Error(w, fmt.Sprintf("failed to delete diary: %+v", err), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/diaries/%d/articles", diaryID), http.StatusSeeOther)
	})
}

func (s *server) spaHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates["spa.tmpl"].ExecuteTemplate(w, "spa.tmpl", nil)
	})
}
