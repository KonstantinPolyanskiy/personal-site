package router

import (
	"github.com/gin-gonic/gin"
	"personal-site/internal/core/router/middlewares"
	"personal-site/internal/core/service_locator"
	"personal-site/internal/handlers"
	"personal-site/internal/logging"
)

type Router struct {
	logger *logging.ModuleLogger

	loggerRegistry *logging.Registry
	serviceLocator *service_locator.ServiceLocator
}

func New(lr *logging.Registry, sl *service_locator.ServiceLocator) *Router {
	return &Router{
		logger:         lr.LoggerFor("RouterLogger"),
		loggerRegistry: lr,
		serviceLocator: sl,
	}
}

func (r *Router) Init() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.DynamicEnableMiddleware())

	// Регистрируем все зависимости для Handler'ов
	handlers.CreateAllHandlers(r.loggerRegistry, r.serviceLocator)

	// Регистрируем все маршруты в роутере c их обработчиками
	for _, h := range handlers.GetAllHandlers() {
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
