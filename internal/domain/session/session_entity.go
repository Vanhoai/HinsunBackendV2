package session

import (
	"time"

	"github.com/google/uuid"
)

type SessionEntity struct {
	ID         uuid.UUID
	AccountID  string
	JTI        string
	ExpiresAt  int64
	IPAddress  string
	UserAgent  string
	DeviceType string
	CreatedAt  int64
	UpdatedAt  int64
	DeletedAt  *int64
}

func NewSessionEntity(id uuid.UUID, accountID, jti string, expiresAt int64, ipAddress, userAgent, deviceType string) *SessionEntity {
	now := time.Now()
	return &SessionEntity{
		ID:         id,
		AccountID:  accountID,
		JTI:        jti,
		ExpiresAt:  expiresAt,
		IPAddress:  ipAddress,
		UserAgent:  userAgent,
		DeviceType: deviceType,
		CreatedAt:  now.Unix(),
		UpdatedAt:  now.Unix(),
		DeletedAt:  nil,
	}
}
