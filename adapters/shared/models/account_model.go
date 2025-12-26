package models

import "github.com/google/uuid"

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
