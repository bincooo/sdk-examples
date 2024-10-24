package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Router(path="/echo")
type EchoHandler struct {
}

// @Ioc()
func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

// @GetMapping(path="hi")
func (e *EchoHandler) Hi(gtx *gin.Context) {
	gtx.String(http.StatusOK, "hi ~")
}
