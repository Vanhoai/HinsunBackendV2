package databases

import (
	"fmt"
	"hinsun-backend/configs"
	"hinsun-backend/internal/core/log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresClient struct {
	DB *gorm.DB
}

func NewPostgresClient() (*PostgresClient, error) {
	cfg := configs.GlobalConfig.Database
	dsn := cfg.GetDSN()

	// Configure GORM logger
	gormLogger := logger.Default.LogMode(logger.Info)
	if !configs.GlobalConfig.App.Debug {
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	// Open database connection
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation(cfg.TimeZone)
			return time.Now().In(loc)
		},
		PrepareStmt: true,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres database: %w", err)
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(cfg.MaxConnections)
	db.SetMaxIdleConns(cfg.MinConnections)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.IdleTimeout) * time.Second)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// gormDB.AutoMigrate(
	// 	&models.ExperienceModel{},
	// 	&models.AccountModel{},
	// 	&models.ProjectModel{},
	// 	&models.BlogModel{},
	// )
	log.Logger.Info("✅ Database connection established successfully")
	return &PostgresClient{DB: gormDB}, nil
}
