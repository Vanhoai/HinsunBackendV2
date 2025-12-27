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
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
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
