package di

import (
	"context"
	"hinsun-backend/adapters/shared/databases"
	"hinsun-backend/configs"
	"hinsun-backend/pkg/jwt"
	"hinsun-backend/pkg/security"
	"time"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var CoreModule = fx.Module("database",
	fx.Provide(
		PrivideDatabase,
		ProvidePasswordHasher,
		ProvideKeyManager,
		ProvideJwtService,
	),
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

func ProvidePasswordHasher() security.PasswordHasher {
	return security.NewArgon2Hasher(security.DefaultArgon2Params())
}

func ProvideKeyManager() *jwt.KeyManager {
	jwtConfig := configs.GlobalConfig.Jwt
	return jwt.NewKeyManager(jwtConfig.Algorithm, jwtConfig.KeySize)
}

func ProvideJwtService(keyManager *jwt.KeyManager) jwt.JwtService {
	jwtConfig := configs.GlobalConfig.Jwt
	return jwt.NewJwtService(
		keyManager,
		time.Duration(jwtConfig.AccessTokenExpiry),
		time.Duration(jwtConfig.RefreshTokenExpiry),
	)
}
