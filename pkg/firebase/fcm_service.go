package firebase

import (
	"context"
	"fmt"

	googleFirebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"go.uber.org/zap"
)

type PushNotification struct {
	Title    string
	Body     string
	ImageURL string
	Data     map[string]string
	Priority string
}

type FCMService interface {
	SendNotification(ctx context.Context, token string, notification *PushNotification) error
	SendMulticast(ctx context.Context, tokens []string, notification *PushNotification) error
	SendToTopic(ctx context.Context, topic string, notification *PushNotification) error
}

type fcmService struct {
	client *messaging.Client
	logger *zap.Logger
}

func NewFCMService(app *googleFirebase.App, logger *zap.Logger) (FCMService, error) {
	messaging, err := app.Messaging(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Messaging client: %v", err)
	}

	logger.Info("FCM service initialized successfully")
	return &fcmService{
		client: messaging,
		logger: logger,
	}, nil
}

func (s *fcmService) SendNotification(ctx context.Context, token string, notification *PushNotification) error {
	message := s.buildMessage(token, notification)
	response, err := s.client.Send(ctx, message)
	if err != nil {
		s.logger.Error("Failed to send FCM notification",
			zap.String("token", token),
			zap.Error(err))

		return fmt.Errorf("failed to send notification: %w", err)
	}

	s.logger.Info("FCM notification sent successfully",
		zap.String("message_id", response),
		zap.String("token", token))

	return nil
}

func (s *fcmService) SendMulticast(ctx context.Context, tokens []string, notification *PushNotification) error {
	if len(tokens) == 0 {
		s.logger.Warn("No tokens provided for multicast")
		return nil
	}

	message := s.buildMulticastMessage(tokens, notification)
	response, err := s.client.SendMulticast(ctx, message)
	if err != nil {
		s.logger.Error("Failed to send multicast notification", zap.Error(err))
		return fmt.Errorf("failed to send multicast: %w", err)
	}

	s.logger.Info("Multicast notification sent",
		zap.Int("success_count", response.SuccessCount),
		zap.Int("failure_count", response.FailureCount))

	if response.FailureCount > 0 {
		for idx, resp := range response.Responses {
			if !resp.Success {
				s.logger.Warn("Failed to send to token",
					zap.String("token", tokens[idx]),
					zap.String("error", resp.Error.Error()))
			}
		}
	}

	return nil
}

func (s *fcmService) SendToTopic(ctx context.Context, topic string, notification *PushNotification) error {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title:    notification.Title,
			Body:     notification.Body,
			ImageURL: notification.ImageURL,
		},
		Data:  notification.Data,
		Topic: topic,
		Android: &messaging.AndroidConfig{
			Priority: notification.Priority,
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
		},
	}

	response, err := s.client.Send(ctx, message)
	if err != nil {
		s.logger.Error("Failed to send topic notification",
			zap.String("topic", topic),
			zap.Error(err))
		return fmt.Errorf("failed to send to topic: %w", err)
	}

	s.logger.Info("Topic notification sent successfully",
		zap.String("message_id", response),
		zap.String("topic", topic))

	return nil
}

func (s *fcmService) buildMulticastMessage(
	tokens []string,
	notification *PushNotification,
) *messaging.MulticastMessage {
	return &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title:    notification.Title,
			Body:     notification.Body,
			ImageURL: notification.ImageURL,
		},
		Data: notification.Data,
		Android: &messaging.AndroidConfig{
			Priority: notification.Priority,
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
		},
	}
}

func (s *fcmService) buildMessage(
	token string,
	notification *PushNotification,
) *messaging.Message {
	return &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title:    notification.Title,
			Body:     notification.Body,
			ImageURL: notification.ImageURL,
		},
		Data: notification.Data,
		Android: &messaging.AndroidConfig{
			Priority: notification.Priority,
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
		},
	}
}
