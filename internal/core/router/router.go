package router

import (
	"github.com/gin-gonic/gin"
	"personal-site/internal/logging"
)

type Router struct {
	logger *logging.ModuleLogger
}

func New(lr *logging.Registry) *Router {
	return &Router{
		lr.LoggerFor("RouterLogger"),
	}
}

func (r Router) Init() *gin.Engine {
	router := gin.Default()

	router.GET("/")
}
