package router

import (
	"github.com/gin-gonic/gin"
	"personal-site/internal/core/router/middlewares"
	"personal-site/internal/handlers"
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

	router.Use(middlewares.DynamicEnableMiddleware())

	// Регистрируем все маршруты в роутере
	for _, h := range handlers.All() {
		switch h.Method() {
		case "GET":
			router.GET(h.Path(), h.Handle)
		case "POST":
			router.POST(h.Path(), h.Handle)
		case "PUT":
			router.PUT(h.Path(), h.Handle)
		case "PATCH":
			router.PATCH(h.Path(), h.Handle)
		case "DELETE":
			router.DELETE(h.Path(), h.Handle)
		}
	}

	return router

}
