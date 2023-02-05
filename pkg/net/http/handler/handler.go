package handler

import "github.com/gin-gonic/gin"

// Handler interface
type Handler interface {
	Ping(c *gin.Context)
}

type handler struct {
	Logger
}

// New returns a new instance that implements the Handler interface
func New(logger Logger) Handler {
	return &handler{logger}
}

// Logger interface represents an object that can log.
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
