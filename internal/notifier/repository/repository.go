package repository

import (
	models "github.com/maruki00/deligo/internal/notifier/domain"
	"gorm.io/gorm"
)

type NotifierRepository struct {
	db *gorm.DB
}

func NewNotifierRepository(db *gorm.DB) *NotifierRepository {
	return &NotifierRepository{db: db}
}

func (r *NotifierRepository) CreateNotification(n *models.Notification) error {
	return r.db.Create(n).Error
}

func (r *NotifierRepository) GetUnreadNotifications(userID string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("user_id = ? AND is_read = ?", userID, false).Find(&notifications).Error
	return notifications, err
}

func (r *NotifierRepository) MarkNotificationAsRead(id uint, userID string) error {
	return r.db.Model(&models.Notification{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("is_read", true).Error
}

func (r *NotifierRepository) CreateChatMessage(m *models.ChatMessage) error {
	return r.db.Create(m).Error
}

func (r *NotifierRepository) GetChatHistory(orderID string) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	err := r.db.Where("order_id = ?", orderID).Order("sent_at asc").Find(&messages).Error
	return messages, err
}

func (r *NotifierRepository) CreateFile(f *models.File) error {
	return r.db.Create(f).Error
}
