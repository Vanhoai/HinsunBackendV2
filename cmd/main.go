package main

import (
	"context"
	"fmt"
	v1 "hinsun-backend/adapters/primary/v1"
	v2 "hinsun-backend/adapters/primary/v2"
	"hinsun-backend/adapters/shared/di"
	"hinsun-backend/configs"
	_ "hinsun-backend/docs"
	"hinsun-backend/internal/core/log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
)

type HTTPServer struct {
	server  *http.Server
	address string
}

// ServerParams contains all dependencies for HTTP server
type ServerParams struct {
	fx.In

	// V1 Handlers
	V1Routes *v1.V1Routes
	V2Routes *v2.V2Routes
}

// ProvideHTTPServer creates and configures the HTTP server
func ProvideHTTPServer(params ServerParams) *HTTPServer {
	return NewHTTPServer(
		configs.GlobalConfig.Server.Address,
		params,
	)
}

func NewHTTPServer(address string, params ServerParams) *HTTPServer {
	r := chi.NewRouter()

	corsConfig := configs.GlobalConfig.Cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   corsConfig.AllowedOrigins,
		AllowedMethods:   corsConfig.AllowedMethods,
		AllowedHeaders:   corsConfig.AllowedHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
		MaxAge:           corsConfig.MaxAge,
	}))

	// Global middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Timeout(60 * time.Second))

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hi there! I'm healthy 🐳"))
	})

	// Swagger
	r.Mount("/swagger/", httpSwagger.WrapHandler)

	// API Routes
	r.Mount("/api/v1", params.V1Routes.RegisterRoutes())
	r.Mount("/api/v2", params.V2Routes.RegisterRoutes())

	server := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  time.Duration(configs.GlobalConfig.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(configs.GlobalConfig.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(configs.GlobalConfig.Server.IdleTimeout) * time.Second,
	}

	return &HTTPServer{
		server:  server,
		address: address,
	}
}

func (s *HTTPServer) Start() error {
	log.Logger.Info("🚀 Server starting on address " + s.address)
	log.Logger.Info("📚 Swagger UI: http://" + s.address + "/swagger/index.html")
	log.Logger.Info("💚 Health check: http://" + s.address + "/health")

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	log.Logger.Info("🛑 Shutting down server")
	return s.server.Shutdown(ctx)
}

// RegisterServerHooks registers lifecycle hooks for HTTP server
func RegisterServerHooks(lc fx.Lifecycle, server *HTTPServer) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Start(); err != nil {
					log.Logger.Fatal("Server error: " + err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop(ctx)
		},
	})
}

var AppModule = fx.Module("http_server",
	fx.Provide(ProvideHTTPServer),
	fx.Invoke(RegisterServerHooks),
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
	// Ensure configurations are initialized
	configs.Init()
	log.Init()

	app := fx.New(
		// Core modules
		di.CoreModule,
		di.RepositioryModule,
		di.ServiceModule,
		di.ApplicationModule,

		// HTTP modules
		di.HandlerModule,
		di.RouterVersionModule,
		AppModule,
	)

	app.Run()
}
