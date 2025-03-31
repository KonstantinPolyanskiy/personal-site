package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"personal-site/internal/logging"
	"sync"
)

var NotFoundHandlerErr = errors.New("handler not found")

type ServiceLocator interface {
	Register(name string, service interface{})
	Get(name string) (interface{}, bool)
}
type HandlerFactory func(registry *logging.Registry, serviceLocator ServiceLocator) Handler

type Metadata interface {
	Name() string
	Description() string
}

type Routing interface {
	Path() string
	Method() string
}

type Config interface {
	Required() bool
	Enabled() bool
	NeedAuth() bool
}

type Html interface {
	TemplatesPath() []string
	ResetPath(newPath []string)
}

type Handler interface {
	Metadata
	Routing
	Config
	Html

	Handle(*gin.Context)
}

var (
	factories []HandlerFactory
	mu        sync.RWMutex

	// AllHandlers Глобальное хранилище Handler, ключ - method:path
	AllHandlers = make(map[string]Handler)
)

func RegisterFactory(factory HandlerFactory) {
	mu.Lock()
	defer mu.Unlock()
	factories = append(factories, factory)
}

// CreateAllHandlers Создает с помощью фабрик Handler'ы и передает зависимости, сохраняет созданные Handler'ы в глобальную мапу
func CreateAllHandlers(lr *logging.Registry, sl ServiceLocator) {
	mu.Lock()
	defer mu.Unlock()

	// Очищаем/пересоздаём мапу
	AllHandlers = make(map[string]Handler)

	// Перебираем все фабрики
	for _, f := range factories {
		h := f(lr, sl)
		AllHandlers[h.Method()+":"+h.Path()] = h
	}
}

func GetAllHandlers() []Handler {
	mu.RLock()
	defer mu.RUnlock()

	result := make([]Handler, 0, len(AllHandlers))
	for _, h := range AllHandlers {
		result = append(result, h)
	}

	return result
}

// ByMethodAndPath находит уже созданный Handler по методу:пути
func ByMethodAndPath(key string) (Handler, error) {
	mu.RLock()
	defer mu.RUnlock()

	h, ok := AllHandlers[key]
	if !ok {
		return nil, NotFoundHandlerErr
	}
	return h, nil
}
