package server

import (
	"fmt"
	"net/http"

	"github.com/3P3-21/curriculum/internal/server/req"
	"github.com/3P3-21/curriculum/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	server *http.Server
	router chi.Router

	service *service.Service
}

type Config struct {
	Addr    string
	Port    int
	Service *service.Service
}

func NewServer(config *Config) *Server {
	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprint(config.Addr, ":", config.Port),
			Handler: http.NotFoundHandler(),
		},
		router: chi.NewRouter(),

		service: config.Service,
	}
}

func (s *Server) SetupRouter() {
	s.setupCors()

	s.router.Route("/user", func(r chi.Router) {
		r.Method("POST", "/signup", req.NewHandler(s.service.User.SignUp))
		r.Method("POST", "/signip", req.NewHandler(s.service.User.SignIn))
	})

	s.server.Handler = s.router
}

func (s *Server) RunServer() error {
	return s.server.ListenAndServe()
}

func (s *Server) setupCors() {
	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)
}
