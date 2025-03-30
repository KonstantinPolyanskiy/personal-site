package handlers

import (
	"github.com/gin-gonic/gin"
	"sync"
)

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

func All() []Handler {
	mu.RLock()
	defer mu.RUnlock()

	result := make([]Handler, len(handlers))
	copy(result, handlers)

	return result
}
