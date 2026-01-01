package configs

type Config struct {
	Env      Env
	App      AppConfig
	Server   ServerConfig
	Cors     CorsConfig
	Metrics  MetricsConfig
	Log      LogConfig
	Database DatabaseConfig
	Caching  CachingConfig
	Jwt      JwtConfig
}

type AppConfig struct {
	Name    string
	Debug   bool
	Version string
}

type ServerConfig struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type CorsConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
	MaxAge           int
}

type MetricsConfig struct {
	Enabled bool
	Port    int
}

type LogConfig struct {
	SavePath          string
	FileName          string
	MaxSize           int
	MaxAge            int
	LocalTime         bool
	Compress          bool
	Level             string
	EnableWriteToFile bool
	EnableConsole     bool
	EnableColor       bool
	EnableCaller      bool
	EnableStacktrace  bool
}

type DatabaseConfig struct {
	User            string
	Password        string
	Host            string
	Port            int
	Database        string
	SSLMode         string
	MaxConnections  int
	MinConnections  int
	MaxConnLifetime int
	IdleTimeout     int
	ConnectTimeout  int
	TimeZone        string
}

type CachingConfig struct {
	Host         string
	Port         int
	Password     string
	DB           int
	PoolSize     int
	IdleTimeout  int
	MinIdleConns int
}

type JwtConfig struct {
	Algorithm          string
	KeySize            int
	AccessTokenExpiry  int
	RefreshTokenExpiry int
}
