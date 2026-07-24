package service

import (
	"log"

	"github.com/maruki00/deligo/internal/notifier/domain"
)

type NotifierService struct {
	repo domain.NotifierRepository
}

func NewNotifierService(repo domain.NotifierRepository) *NotifierService {
	return &NotifierService{repo: repo}
}

func (s *NotifierService) ProcessInboundNotification(userID, title, body string) error {
	n := &domain.Notification{
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

func (s *NotifierService) GetUnreadNotifications(userID string) ([]domain.Notification, error) {
	return s.repo.GetUnreadNotifications(userID)
}

func (s *NotifierService) MarkNotificationAsRead(id uint, userID string) error {
	return s.repo.MarkNotificationAsRead(id, userID)
}

func (s *NotifierService) SaveChatMessage(m *domain.ChatMessage) error {
	return s.repo.CreateChatMessage(m)
}

func (s *NotifierService) GetChatHistory(orderID string) ([]domain.ChatMessage, error) {
	return s.repo.GetChatHistory(orderID)
}

func (s *NotifierService) UploadFileMetadata(f *domain.File) error {
	return s.repo.CreateFile(f)
}
