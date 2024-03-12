//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=types.cfg.yaml ./api.yaml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ./api.yaml

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIHandler struct{}

func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

func (s *APIHandler) GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, world!")
}
