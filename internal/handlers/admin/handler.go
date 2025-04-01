package admin

import (
	"github.com/gin-gonic/gin"
	"personal-site/internal/logging"
)

type Handler struct {
	logger *logging.ModuleLogger
}

func (h Handler) Name() string {
	return "AdminPage"
}

func (h Handler) Description() string {
	return "Панель администратора"
}

func (h Handler) Path() string {
	return "/admin"
}

func (h Handler) Method() string {
	return "GET"
}

func (h Handler) Handle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Required() bool {
	return true
}

func (h Handler) Enabled() bool {
	return true
}

func (h Handler) NeedAuth() bool {
	return true
}
