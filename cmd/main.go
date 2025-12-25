package main

import (
	v1 "hinsun-backend/adapters/primary/v1"
	v2 "hinsun-backend/adapters/primary/v2"
	"hinsun-backend/configs"
	_ "hinsun-backend/docs"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Hinsun Backend API
// @version         1.0
// @description     Backend API for Hinsun portfolio and blog system
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.hinsun.dev/support
// @contact.email  support@hinsun.dev

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Initialize configurations
	configs.Init()
	mux := http.NewServeMux()

	// Register API routes
	mux.Handle("/api/v1/", v1.RegisterV1Routes())
	mux.Handle("/api/v2/", v2.RegisterV2Routes())

	// Swagger UI endpoint
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Printf("Server starting on %s", configs.GlobalConfig.Server.Address)
	log.Printf("Swagger UI: http://%s/swagger/index.html", configs.GlobalConfig.Server.Address)
	log.Fatal(http.ListenAndServe(configs.GlobalConfig.Server.Address, mux))
}
