package di

import (
	"context"
	"hinsun-backend/adapters/shared/databases"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var DatabaseModule = fx.Module("database",
	fx.Provide(PrivideDatabase),
	fx.Invoke(RegisterDatabaseHook),
)

func PrivideDatabase() (*gorm.DB, error) {
	client, err := databases.NewPostgresClient()
	if err != nil {
		return nil, err
	}

	return client.DB, nil
}

func RegisterDatabaseHook(lc fx.Lifecycle, db *gorm.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Database is already connected in ProvideDatabase
			return nil
		},
		OnStop: func(ctx context.Context) error {
			sqlDB, err := db.DB()
			if err != nil {
				return err
			}
			return sqlDB.Close()
		},
	})
}
