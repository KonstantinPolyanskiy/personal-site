package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"personal-site/internal/handlers"
)

func DynamicEnableMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		route := c.FullPath()

		handler, err := handlers.ByPath(route)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if !handler.Enabled() {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.Next()
	}
}
