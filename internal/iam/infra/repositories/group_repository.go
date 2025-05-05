package repositories

import (
	"context"
	"deligo/internal/iam/infra/models"
	pkgPostgres "deligo/pkg/postgres"
)

type GroupRepository struct {
	db *pkgPostgres.PGHandler
}

func NewGroupRepository(db *pkgPostgres.PGHandler) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
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
