package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/viniokamoto/go-store/source/authentication/middlewares"
)

var AuthMiddleware = middlewares.AuthMiddleware

type Server struct {
	Engine  *gin.Engine
	options Options
}

type Options struct {
	NoAuthMiddleware bool
}

func CreateServer(options Options) Server {
	server := Server{}
	server.Init(options)

	return server
}

func (s *Server) Init(options Options) {
	if s.Engine == nil {
		s.Engine = gin.Default()
		s.bindCommonMiddlewares()
		s.options = options

	}
}

func (s *Server) AddMiddleware(middleware gin.HandlerFunc) {
	s.Engine.Use(middleware)
}

func (s *Server) AddRoutes(routes []ApiRoute) {
	for _, route := range routes {
		s.AddRoute(route)
	}
}

func (s *Server) AddRoute(route ApiRoute) {
	handlers := []gin.HandlerFunc{route.Handler}

	if !route.IsAnonymous {
		handlers = append([]gin.HandlerFunc{AuthMiddleware}, handlers...)
	}

	switch route.MethodType {
	case GET:
		s.Engine.GET(route.Path, handlers...)

	case POST:
		s.Engine.POST(route.Path, handlers...)

	case PUT:
		s.Engine.PUT(route.Path, handlers...)

	case DELETE:
		s.Engine.DELETE(route.Path, handlers...)

	}
}

func (s *Server) Start(port int) {

	s.Engine.Run(fmt.Sprintf(":%d", port))
}

func (s *Server) bindCommonMiddlewares() {
	s.Engine.Use(gin.Logger())
}
