package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
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
	r.LoadHTMLGlob("html/*.html")
	r.GET("/", s.handlerIndex)
	r.GET("/ping", s.handleFunc)
	if err := r.Run(s.config.BindAddr); err != nil {
		log.Fatal(err)
	}
}

func (s *APIServer) handleFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "pong")
		if err != nil {
			return
		}
	}
}

func (s *APIServer) handlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
