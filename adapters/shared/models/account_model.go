package models

import (
	"hinsun-backend/internal/core/failure"
	"hinsun-backend/internal/domain/account"
	"hinsun-backend/internal/domain/values"

	"github.com/google/uuid"
)

type AccountModel struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Email         string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	EmailVerified bool      `gorm:"type:boolean;default:false;not null"`
	Role          int       `gorm:"type:int;default:0;not null"`
	IsActive      bool      `gorm:"type:boolean;default:true;not null"`
	Password      string    `gorm:"type:varchar(255);not null"`
	Avatar        string    `gorm:"type:varchar(255)"`
	Bio           string    `gorm:"type:text"`
	CreatedAt     int64     `gorm:"autoCreateTime"`
	UpdatedAt     int64     `gorm:"autoUpdateTime"`
	DeletedAt     *int64    `gorm:"index"`

	// Relationship: One Account has Many Blogs
	Blogs []BlogModel `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (AccountModel) TableName() string { return "accounts" }

func (e *AccountModel) ToEntity() (*account.AccountEntity, error) {
	email, err := values.NewEmail(e.Email)
	if err != nil {
		return nil, failure.NewInternalFailure("invalid email in account model: %v", err)
	}

	role, err := values.RoleFromInt(e.Role)
	if err != nil {
		return nil, failure.NewInternalFailure("invalid role in account model: %v", err)
	}

	return &account.AccountEntity{
		ID:            e.ID,
		Name:          e.Name,
		Email:         email,
		EmailVerified: e.EmailVerified,
		IsActive:      e.IsActive,
		Password:      e.Password,
		Role:          role,
		Avatar:        e.Avatar,
		Bio:           e.Bio,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		DeletedAt:     e.DeletedAt,
	}, nil
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
