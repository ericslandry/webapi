package main

import (
	"log"
	"net"
	"net/http"
	"server/api"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func NewServer(handler *api.APIHandler, port string) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec\n: %s", err)
	}
	swagger.Servers = nil // Skips validating that server names match
	r := gin.Default()
	r.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(r, handler)
	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}
	return s
}

func main() {
	s := NewServer(api.NewAPIHandler(), "8080")
	log.Fatal(s.ListenAndServe())
}
