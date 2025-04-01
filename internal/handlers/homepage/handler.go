package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"personal-site/internal/handlers"
	"personal-site/internal/logging"
)

type Handler struct {
	logger *logging.ModuleLogger

	StaticFiles []string
}

func init() {
	handlers.RegisterFactory(func(registry *logging.Registry, serviceLocator handlers.ServiceLocator) handlers.Handler {
		return &Handler{
			logger:      registry.LoggerFor("HomepageLogger"),
			StaticFiles: []string{"./ui/html/home.page.gohtml", "./ui/html/base.layout.gohtml", "./ui/html/footer.partial.gohtml"},
		}
	})
}

func (h *Handler) TemplatesPath() []string {
	return h.StaticFiles
}

func (h *Handler) ResetPath(newPath []string) {
	h.StaticFiles = newPath
}

func (h *Handler) Name() string {
	return "Homepage"
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
	context.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	ts, err := template.ParseFiles(h.StaticFiles...)
	if err != nil {
		h.logger.Warn("error parse html file", zap.Error(err))
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := ts.Execute(context.Writer, nil); err != nil {
		h.logger.Warn("error execute template", zap.Error(err))
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
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
