package apiserver

import (
	"ApiServer/internal/app/model"
	"ApiServer/internal/app/store"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router       *gin.Engine
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

const (
	sessionName        = "main"
	ctxKeyUser  ctxKey = iota
)

var (
	errIncorrectLoginOrPassword = errors.New("incorrect login or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type ctxKey int8

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       gin.Default(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	s.router.Handle("POST", "/student", s.handleStudentCreate())
	s.router.Handle("POST", "/session", s.handleSessionsCreate())
	s.router.Handle("GET", "/whoAmI", s.handleWhoAmI())
	//private := s.router.Group("/private")
	//private.Use(s.authenticateStudent)
	//
	//{
	//	auth := private.Group("/")
	//	{
	//		auth.GET("/", s.handleWhoAmI())
	//	}
	//
	//}
}

func (s *server) handleStudentCreate() gin.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(c *gin.Context) {
		req := &request{}
		if err := json.NewDecoder(c.Request.Body).Decode(req); err != nil {
			s.error(c.Writer, c.Request, http.StatusBadRequest, c.Error(err))
			return
		}

		student := &model.Student{
			Login:    req.Login,
			Password: req.Password,
		}
		if err := s.store.Student().Create(student); err != nil {
			s.error(c.Writer, c.Request, http.StatusUnprocessableEntity, err)
		}

		student.Sanitaize()
		s.respond(c.Writer, c.Request, http.StatusCreated, student)
	}
}

func (s *server) configureLogger() {
	s.logger.SetLevel(s.logger.Level)
}

func (s *server) authenticateStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		stud, err := s.store.Student().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, stud)))
	})
}

func (s *server) handleWhoAmI() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.respond(c.Writer, c.Request, http.StatusOK, c.Request.Context().Value(ctxKeyUser).(model.Student))
	}
}

func (s *server) handleSessionsCreate() gin.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(c *gin.Context) {
		req := &request{}
		if err := json.NewDecoder(c.Request.Body).Decode(req); err != nil {
			s.error(c.Writer, c.Request, http.StatusBadRequest, c.Error(err))
			return
		}

		student, err := s.store.Student().FindByLogin(req.Login)
		if err != nil || !student.ComparePassword(req.Password) {
			s.error(c.Writer, c.Request, http.StatusUnauthorized, errIncorrectLoginOrPassword)
			return
		}
		session, err := s.sessionStore.Get(c.Request, sessionName)
		if err != nil {
			s.error(c.Writer, c.Request, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = student.ID
		if err := s.sessionStore.Save(c.Request, c.Writer, session); err != nil {
			s.error(c.Writer, c.Request, http.StatusInternalServerError, err)
			return
		}

		s.respond(c.Writer, c.Request, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
	}
}
