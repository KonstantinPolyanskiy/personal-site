package homepage

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-site/internal/handlers"
	"personal-site/internal/logging"
)

type Handler struct {
	logger *logging.ModuleLogger
}

func init() {
	handlers.Register(&Handler{})
}

func (h *Handler) RegisterLogger(l *logging.ModuleLogger) {
	h.logger = l
}

func (h *Handler) Name() string {
	return "Домашняя страница"
}

func (h *Handler) Description() string {
	return "Домашняя страница, отображается при открытии site_name.com/"
}

func (h *Handler) Path() string {
	return "/"
}

func (h *Handler) Method() string {
	return "GET"
}

func (h *Handler) Handle(context *gin.Context) {
	context.String(http.StatusOK, h.Description())
}

func (h *Handler) Required() bool {
	return true
}

func (h *Handler) Enabled() bool {
	return true
}

func (h *Handler) NeedAuth() bool {
	return false
}
