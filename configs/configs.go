package configs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Env string // PROD or DEV

const (
	EnvProduction  Env = "PROD"
	EnvDevelopment Env = "DEV"
)

var GlobalConfig *Config

// Init initializes the configuration by loading environment variables
func Init() {
	// get root path of project
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory:", err)
	}

	// Load appropriate .env file
	envFile := filepath.Join(rootPath, ".env")
	env := getEnv("ENV", "DEV")

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: Error loading %s file: %v", envFile, err)
		log.Println("Continuing with system environment variables...")
	} else {
		log.Printf("Successfully loaded environment file: %s", envFile)
	}

	// Parse and set global config
	GlobalConfig = loadConfig(env)
}

// getEnv gets a string environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// loadConfig loads all configuration from environment variables
func loadConfig(env string) *Config {
	return &Config{
		Env:      Env(env),
		App:      loadAppConfig(),
		Server:   loadServerConfig(),
		Cors:     loadCorsConfig(),
		Metrics:  loadMetricsConfig(),
		Log:      loadLogConfig(),
		Database: loadDatabaseConfig(),
		Caching:  loadCachingConfig(),
		Jwt:      loadJWTConfig(),
	}
}

// loadAppConfig loads application configuration
func loadAppConfig() AppConfig {
	return AppConfig{
		Name:    getEnv("APP_NAME", "Hinsun Backend"),
		Debug:   getEnvAsBool("APP_DEBUG", true),
		Version: getEnv("APP_VERSION", "1.0.0"),
	}
}

// loadServerConfig loads server configuration
func loadServerConfig() ServerConfig {
	return ServerConfig{
		Address:      getEnv("SERVER_ADDRESS", ":8080"),
		ReadTimeout:  getEnvAsInt("SERVER_READ_TIMEOUT", 15),
		WriteTimeout: getEnvAsInt("SERVER_WRITE_TIMEOUT", 15),
		IdleTimeout:  getEnvAsInt("SERVER_IDLE_TIMEOUT", 60),
	}
}

// loadCorsConfig loads CORS configuration
func loadCorsConfig() CorsConfig {
	return CorsConfig{
		AllowedOrigins:   getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{}, ","),
		AllowedMethods:   getEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, ","),
		AllowedHeaders:   getEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, ","),
		AllowCredentials: getEnvAsBool("CORS_ALLOW_CREDENTIALS", false),
		MaxAge:           getEnvAsInt("CORS_MAX_AGE", 300),
	}
}

// getEnvAsSlice gets a string slice environment variable or returns a default value
func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	return splitAndTrim(valueStr, sep)
}

func splitAndTrim(s, sep string) []string {
	parts := []string{}
	for _, part := range strings.Split(s, sep) {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			parts = append(parts, trimmed)
		}
	}

	return parts
}

// loadMetricsConfig loads metrics configuration
func loadMetricsConfig() MetricsConfig {
	return MetricsConfig{
		Enabled: getEnvAsBool("METRICS_ENABLED", false),
		Port:    getEnvAsInt("METRICS_PORT", 9090),
	}
}

// loadLogConfig loads logging configuration
func loadLogConfig() LogConfig {
	return LogConfig{
		SavePath:          getEnv("LOG_SAVE_PATH", "./tmp"),
		FileName:          getEnv("LOG_FILE_NAME", "app"),
		MaxSize:           getEnvAsInt("LOG_MAX_SIZE", 100),
		MaxAge:            getEnvAsInt("LOG_MAX_AGE", 30),
		LocalTime:         getEnvAsBool("LOG_LOCAL_TIME", true),
		Compress:          getEnvAsBool("LOG_COMPRESS", true),
		Level:             getEnv("LOG_LEVEL", "debug"),
		EnableWriteToFile: getEnvAsBool("LOG_ENABLE_WRITE_TO_FILE", false),
		EnableConsole:     getEnvAsBool("LOG_ENABLE_CONSOLE", true),
		EnableColor:       getEnvAsBool("LOG_ENABLE_COLOR", true),
		EnableCaller:      getEnvAsBool("LOG_ENABLE_CALLER", true),
		EnableStacktrace:  getEnvAsBool("LOG_ENABLE_STACKTRACE", false),
	}
}

// loadDatabaseConfig loads database configuration
func loadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		User:            getEnv("DATABASE_USER", "postgres"),
		Password:        getEnv("DATABASE_PASSWORD", ""),
		Host:            getEnv("DATABASE_HOST", "localhost"),
		Port:            getEnvAsInt("DATABASE_PORT", 5432),
		Database:        getEnv("DATABASE_NAME", "hinsun"),
		SSLMode:         getEnv("DATABASE_SSL_MODE", "disable"),
		MaxConnections:  getEnvAsInt("DATABASE_MAX_CONNECTIONS", 25),
		MinConnections:  getEnvAsInt("DATABASE_MIN_CONNECTIONS", 5),
		MaxConnLifetime: getEnvAsInt("DATABASE_MAX_CONN_LIFETIME", 300),
		IdleTimeout:     getEnvAsInt("DATABASE_IDLE_TIMEOUT", 60),
		ConnectTimeout:  getEnvAsInt("DATABASE_CONNECT_TIMEOUT", 10),
		TimeZone:        getEnv("DATABASE_TIMEZONE", "Asia/Ho_Chi_Minh"),
	}
}

// loadCachingConfig loads caching (Redis) configuration
func loadCachingConfig() CachingConfig {
	return CachingConfig{
		Host:         getEnv("REDIS_HOST", "localhost"),
		Port:         getEnvAsInt("REDIS_PORT", 6379),
		Password:     getEnv("REDIS_PASSWORD", ""),
		DB:           getEnvAsInt("REDIS_DB", 0),
		PoolSize:     getEnvAsInt("REDIS_POOL_SIZE", 10),
		IdleTimeout:  getEnvAsInt("REDIS_IDLE_TIMEOUT", 300),
		MinIdleConns: getEnvAsInt("REDIS_MIN_IDLE_CONNS", 5),
	}
}

// loadJWTConfig loads JWT configuration
func loadJWTConfig() JwtConfig {
	return JwtConfig{
		Algorithm:          getEnv("JWT_ALGORITHM", "RS256"),
		KeySize:            getEnvAsInt("JWT_KEYSIZE", 2048),
		AccessTokenExpiry:  getEnvAsInt("JWT_ACCESS_TOKEN_EXPIRY", 900),
		RefreshTokenExpiry: getEnvAsInt("JWT_REFRESH_TOKEN_EXPIRY", 604800),
	}
}

// getEnvAsInt gets an integer environment variable or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid integer value for %s: %s, using default: %d", key, valueStr, defaultValue)
		return defaultValue
	}

	return value
}

// getEnvAsBool gets a boolean environment variable or returns a default value
func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid boolean value for %s: %s, using default: %t", key, valueStr, defaultValue)
		return defaultValue
	}

	return value
}

// String method for Env type
func (e Env) String() string {
	return string(e)
}

// IsDevelopment checks if the current environment is development
func (e Env) IsDevelopment() bool {
	return e == EnvDevelopment
}

// IsProduction checks if the current environment is production
func (e Env) IsProduction() bool {
	return e == EnvProduction
}

// GetDSN returns the database connection string
func (d DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		d.Host, d.Port, d.User, d.Password, d.Database, d.SSLMode, d.TimeZone,
	)
}

// GetRedisAddr returns the Redis address
func (c CachingConfig) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
