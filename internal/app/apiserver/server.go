package apiserver

import (
	"ApiServer/internal/app/model"
	"ApiServer/internal/app/store"
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
	sessionName        = "student_session"
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
	s.router.Static("/assets/", "web/")
	s.router.LoadHTMLGlob("./web/templates/*.html")
	s.router.GET("/", s.handlerIndex)
	s.router.GET("/authorization", s.handlerAuthorization)
	s.router.GET("/registration", s.handleRegistration)
	s.router.Handle("POST", "/student", s.handleStudentCreate)
	s.router.Handle("POST", "/session", s.handleSessionsCreate)
	private := s.router.Group("/private")
	private.Use(s.authenticateStudent)
	{
		auth := private.Group("/")
		{
			auth.GET("/whoAmI", s.handleWhoAmI)
		}

	}
}

func (s *server) handleStudentCreate(c *gin.Context) {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	req := &request{}
	if err := json.NewDecoder(c.Request.Body).Decode(req); err != nil {
		s.error(c, http.StatusBadRequest, c.Error(err))
		return
	}

	student := &model.Student{
		Login:    req.Login,
		Password: req.Password,
	}
	if err := s.store.Student().Create(student); err != nil {
		s.error(c, http.StatusUnprocessableEntity, err)
	}

	student.Sanitaize()
	s.respond(c, http.StatusCreated, student)

}

func (s *server) configureLogger() {
	s.logger.SetLevel(s.logger.Level)
}

func (s *server) authenticateStudent(c *gin.Context) {
	session, err := s.sessionStore.Get(c.Request, sessionName)
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
	}

	id, ok := session.Values["student_id"]
	if !ok {
		s.error(c, http.StatusUnauthorized, errNotAuthenticated)
		return
	}
	student, err := s.store.Student().Find(id.(int))
	if err != nil {
		s.error(c, http.StatusUnauthorized, errNotAuthenticated)
	}
	c.JSON(http.StatusOK, gin.H{
		"ctxKey":  ctxKeyUser,
		"Student": student,
	})
}

func (s *server) handleWhoAmI(c *gin.Context) {
	s.respond(c, http.StatusOK, nil)
}

func (s *server) handleSessionsCreate(c *gin.Context) {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	req := &request{}
	if err := json.NewDecoder(c.Request.Body).Decode(req); err != nil {
		s.error(c, http.StatusBadRequest, c.Error(err))
		return
	}

	student, err := s.store.Student().FindByLogin(req.Login)
	if err != nil || !student.ComparePassword(req.Password) {
		s.error(c, http.StatusUnauthorized, errIncorrectLoginOrPassword)
		return
	}
	session, err := s.sessionStore.Get(c.Request, sessionName)
	if err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}
	session.Values["student_id"] = student.ID
	if err := s.sessionStore.Save(c.Request, c.Writer, session); err != nil {
		s.error(c, http.StatusInternalServerError, err)
		return
	}

	s.respond(c, http.StatusOK, nil)

}

func (s *server) error(c *gin.Context, code int, err error) {
	s.respond(c, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(c *gin.Context, code int, data interface{}) {
	c.Writer.WriteHeader(code)
	if data != nil {
		c.JSON(code, data)
	}
}

func (s *server) handlerIndex(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func (s *server) handlerAuthorization(context *gin.Context) {
	context.HTML(http.StatusOK, "authorization.html", gin.H{})
}

func (s *server) handleRegistration(context *gin.Context) {
	context.HTML(http.StatusOK, "registration.html", gin.H{})
}
