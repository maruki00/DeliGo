package repositories

import (
	"context"
	"deligo/internal/iam/infra/models"
)

type GroupRepository struct{}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{}
}

func (_this *GroupRepository) AssignUserToGroup(ctx context.Context, userID, groupID string) error {
	return nil
}
func (_this *GroupRepository) RemoveUserFromGroup(ctx context.Context, userID, groupID string) error {
	return nil
}
func (_this *GroupRepository) GetUserGroups(ctx context.Context, userID string) ([]*models.Group, error) {
	return nil, nil
}
