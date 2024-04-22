package server

import "github.com/gin-gonic/gin"

type Api interface {
	GetRoutes() []ApiRoute
}

type ApiRoute struct {
	MethodType  MethodType
	IsAnonymous bool
	IsAdmin     bool
	IsStore     bool
	Path        string
	Handler     gin.HandlerFunc
}

type MethodType int

const (
	GET MethodType = iota
	POST
	PUT
	DELETE
)
