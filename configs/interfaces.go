package configs

type Config struct {
	Env      Env
	App      AppConfig
	Server   ServerConfig
	Metrics  MetricsConfig
	Log      LogConfig
	Database DatabaseConfig
	Caching  CachingConfig
}

type AppConfig struct {
	Name    string
	Debug   bool
	Version string
}

type ServerConfig struct {
	Address      string
	ReadTimeout  string
	WriteTimeout string
}

type MetricsConfig struct {
	Enabled bool
	Port    int
}

type LogConfig struct {
	Level  string
	Format string
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
