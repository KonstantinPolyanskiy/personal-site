package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
)

var notFoundHandlerErr = errors.New("handler not found")

type Handler interface {
	Name() string
	Description() string
	Path() string
	Method() string
	Handle(*gin.Context)
	Required() bool
	Enabled() bool
	NeedAuth() bool
}

var (
	handlers []Handler
	mu       sync.RWMutex
)

func Register(handler Handler) {
	mu.Lock()
	defer mu.Unlock()
	handlers = append(handlers, handler)
}

// ByPath Возвращает Handler из глобального реестра по его path, если такой есть
func ByPath(path string) (Handler, error) {
	for _, handler := range handlers {
		if handler.Path() == path {
			return handler, nil
		}
	}

	return nil, notFoundHandlerErr
}

func All() []Handler {
	mu.RLock()
	defer mu.RUnlock()

	result := make([]Handler, len(handlers))
	copy(result, handlers)

	return result
}
