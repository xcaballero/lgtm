package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/xcaballero/lgtm/pkg/net/http/handler"
)

type router struct {
	*gin.Engine
	handler.Handler
}

// New returns a gin implementation of the Router interface.
func New(engine *gin.Engine, h handler.Handler) http.Handler {
	r := &router{engine, h}
	r.withEndpoints()
	return r
}

func (r *router) withEndpoints() {
	r.withPublicEndpoints()
}

func (r *router) withPublicEndpoints() {
	r.GET("/ping", r.Handler.Ping)
}

// LogMiddleware logs all the requests to the router.
func LogMiddleware(log handler.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// Process Request
		c.Next()
		duration := time.Since(start)

		message := "%s %s returned %d in %v"
		args := []any{c.Request.Method, c.Request.RequestURI, c.Writer.Status(), duration}
		if c.Writer.Status() >= 400 {
			log.Errorf(message, args...)
		} else {
			log.Infof(message, args...)
		}
	}
}
