package apiserver

import (
	"ApiServer/internal/app/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: gin.Default(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	s.configureRoute()

	if err := s.configureStore(); err != nil {
		return err
	}

	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRoute() {
	r := s.router
	r.LoadHTMLGlob("web/templates/*.html")
	r.GET("/", s.handlerIndex)
	r.GET("/ping", s.handleFunc)
	if err := r.Run(s.config.BindAddr); err != nil {
		log.Fatal(err)
	}
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st

	return nil
}

func (s *APIServer) handleFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (s *APIServer) handlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
