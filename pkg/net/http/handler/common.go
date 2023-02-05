package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}
