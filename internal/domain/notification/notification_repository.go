package notification

import (
	"context"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *NotificationEntity) error
}
