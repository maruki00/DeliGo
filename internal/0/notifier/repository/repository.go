package repository

import (
	"github.com/maruki00/deligo/internal/notifier/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateNotification(n *models.Notification) error
	GetUnreadNotifications(userID string) ([]models.Notification, error)
	MarkNotificationAsRead(id uint, userID string) error
	CreateChatMessage(m *models.ChatMessage) error
	GetChatHistory(orderID string) ([]models.ChatMessage, error)
	CreateFile(f *models.File) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateNotification(n *models.Notification) error {
	return r.db.Create(n).Error
}

func (r *repository) GetUnreadNotifications(userID string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("user_id = ? AND is_read = ?", userID, false).Find(&notifications).Error
	return notifications, err
}

func (r *repository) MarkNotificationAsRead(id uint, userID string) error {
	return r.db.Model(&models.Notification{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("is_read", true).Error
}

func (r *repository) CreateChatMessage(m *models.ChatMessage) error {
	return r.db.Create(m).Error
}

func (r *repository) GetChatHistory(orderID string) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	err := r.db.Where("order_id = ?", orderID).Order("sent_at asc").Find(&messages).Error
	return messages, err
}

func (r *repository) CreateFile(f *models.File) error {
	return r.db.Create(f).Error
}
