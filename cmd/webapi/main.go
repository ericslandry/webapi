package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"server/api"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"github.com/urfave/cli/v2"
)

var version string

func main() {
	app := &cli.App{
		Name:    "My API Server",
		Usage:   "This is a sample web API server",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "0.0.0.0",
				Usage: "Bind host address for the server",
			},
			&cli.StringFlag{
				Name:  "port",
				Value: "8080",
				Usage: "Bind port for the server",
			},
		},
		Action: func(cCtx *cli.Context) error {
			addr := net.JoinHostPort(cCtx.String("host"), cCtx.String("port"))
			s := NewServer(api.NewAPIHandler(), addr)
			log.Fatal(s.ListenAndServe())
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func NewServer(handler *api.APIHandler, addr string) *http.Server {
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
		Addr:    addr,
	}
	return s
}
