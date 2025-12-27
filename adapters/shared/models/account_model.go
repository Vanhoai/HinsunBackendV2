package models

import (
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/values"

	"github.com/google/uuid"
)

type AccountModel struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Email         string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	EmailVerified bool      `gorm:"type:boolean;default:false;not null"`
	IsActive      bool      `gorm:"type:boolean;default:true;not null"`
	Password      string    `gorm:"type:varchar(255);not null"`
	Avatar        string    `gorm:"type:varchar(255)"`
	Bio           string    `gorm:"type:text"`
	CreatedAt     int64     `gorm:"autoCreateTime:nano"`
	UpdatedAt     int64     `gorm:"autoUpdateTime:nano"`
	DeletedAt     *int64    `gorm:"index"`
}

func (AccountModel) TableName() string { return "accounts" }

func ToAccountEntity(model *AccountModel) *account.AccountEntity {
	email, err := values.NewEmail(model.Email)
	if err != nil {
		// Handle error appropriately, possibly returning nil or logging
		return nil
	}

	return &account.AccountEntity{
		ID:            model.ID,
		Name:          model.Name,
		Email:         email,
		EmailVerified: model.EmailVerified,
		IsActive:      model.IsActive,
		Password:      model.Password,
		Avatar:        model.Avatar,
		Bio:           model.Bio,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt,
	}
}

func FromAccountEntity(entity *account.AccountEntity) AccountModel {
	return AccountModel{
		ID:            entity.ID,
		Name:          entity.Name,
		Email:         entity.Email.Value(),
		EmailVerified: entity.EmailVerified,
		IsActive:      entity.IsActive,
		Password:      entity.Password,
		Avatar:        entity.Avatar,
		Bio:           entity.Bio,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
		DeletedAt:     entity.DeletedAt,
	}
}
