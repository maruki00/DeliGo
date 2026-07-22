package service

import (
	"log"

	"github.com/maruki00/deligo/notifier/models"
	"github.com/maruki00/deligo/notifier/repository"
)

type Service interface {
	ProcessInboundNotification(userID, title, body string) error
	GetUnreadNotifications(userID string) ([]models.Notification, error)
	MarkNotificationAsRead(id uint, userID string) error
	SaveChatMessage(m *models.ChatMessage) error
	GetChatHistory(orderID string) ([]models.ChatMessage, error)
	UploadFileMetadata(f *models.File) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) ProcessInboundNotification(userID, title, body string) error {
	n := &models.Notification{
		UserID: userID,
		Title:  title,
		Body:   body,
	}
	if err := s.repo.CreateNotification(n); err != nil {
		return err
	}

	// Mocking third-party integration communications loops (SMS/Email)
	log.Printf("[MOCK DISPATCH] Sending Email alert to User ID %s: %s - %s", userID, title, body)
	log.Printf("[MOCK DISPATCH] Sending SMS push wire to User ID %s: %s", userID, title)

	return nil
}

func (s *service) GetUnreadNotifications(userID string) ([]models.Notification, error) {
	return s.repo.GetUnreadNotifications(userID)
}

func (s *service) MarkNotificationAsRead(id uint, userID string) error {
	return s.repo.MarkNotificationAsRead(id, userID)
}

func (s *service) SaveChatMessage(m *models.ChatMessage) error {
	return s.repo.CreateChatMessage(m)
}

func (s *service) GetChatHistory(orderID string) ([]models.ChatMessage, error) {
	return s.repo.GetChatHistory(orderID)
}

func (s *service) UploadFileMetadata(f *models.File) error {
	return s.repo.CreateFile(f)
}
