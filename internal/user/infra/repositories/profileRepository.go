package repositories

import (
	"context"
	"delivery/internal/user/domain/entities"
	"delivery/internal/user/infra/models"
	pkgPostgres "delivery/pkg/postgres"
	"sync"
)

type ProfileRepository struct {
	sync.RWMutex
	db pkgPostgres.PGHandler
}

func NewProfileRepository(db pkgPostgres.PGHandler) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (ur *ProfileRepository) Create(ctx context.Context, entity entities.ProfileEntity) (entities.ProfileEntity, error) {
	sql := `INSERT INTO profiles(user_id, full_name, avatar, bio) VALUES($1, $2, $3, $44) RETURNING id`
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	var lastInsertID string
	err = tx.QueryRow(sql, entity.GetUserID(), entity.GetFullName(), entity.GetAvatar(), entity.GetBio()).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	entity.SetID(lastInsertID)
	return entity, nil
}

func (ur *ProfileRepository) Delete(ctx context.Context, id string) (bool, error) {
	sql := `UPDATE profiles SET deleted_at = now(), updated_at = now() WHERE id = $1 `
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return false, err
	}
	_, err = tx.Exec(sql, id)
	if err != nil {
		return false, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return false, err
	}
	return true, nil
}

func (ur *ProfileRepository) Update(ctx context.Context, entity entities.ProfileEntity) (entities.ProfileEntity, error) {

	sql := `
			UPDATE profiles 
			SET user_id=$1, full_name=$2, avatar=$3, bio=$4, updated_at=now()
			WHERE id=$5 
			AND deleted_at = NULL
		`
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return nil, err
	}

	_ = tx.QueryRow(sql, entity.GetUserID(), entity.GetFullName(), entity.GetAvatar(), entity.GetBio(), entity.GetID())

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return entity, nil
}

func (ur *ProfileRepository) GetOne(ctx context.Context, id string) (entities.ProfileEntity, error) {
	var entity models.Profile
	sql := `
			SELECT id, user_id, full_name, avatar, bio, created_at, updated_at
			FROM users 
			WHERE id = $1 AND deleted_at = NULL 
			LIMIT 1 
		`
	err := ur.db.GetDB().QueryRow(sql, id).Scan(&entity.ID, &entity.UserID, &entity.FullName, &entity.Avatar, &entity.Bio, &entity.CreatedAt, &entity.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (ur *ProfileRepository) GetMany(ctx context.Context, limit, offset int) ([]entities.ProfileEntity, error) {
	entities := make([]entities.ProfileEntity, offset)
	sql := `
			SELECT id, user_id, full_name, avatar, bio, created_at, updated_at
			FROM users 
			WHERE deleted_at = NULL 
			LIMIT $1 
			OFFSET $2
		`
	rows, err := ur.db.GetDB().Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}
	index := 0
	for ; rows.Next(); index++ {
		entity := models.Profile{}
		rows.Scan(&entity.ID, &entity.UserID, &entity.FullName, &entity.Avatar, &entity.Bio, &entity.CreatedAt, &entity.UpdatedAt)
		entities[index] = &entity
		index++
	}
	return entities[:index], nil
}

func (ur *ProfileRepository) Search(ctx context.Context, query string, limit, offset int) ([]entities.ProfileEntity, error) {
	entities := make([]entities.ProfileEntity, offset)
	sql := `
			SELECT id, user_id, full_name, avatar, bio, created_at, updated_at
			FROM users 
			WHERE deleted_at = NULL 
			AND (id like '%$1' OR user_id like '%$1' OR full_name like '%$1' OR avatar like '%$1' OR bio like '%$1')
			LIMIT $2
			OFFSET $3
		`
	rows, err := ur.db.GetDB().Query(sql, query, limit, offset)
	if err != nil {
		return nil, err
	}
	index := 0
	for ; rows.Next(); index++ {
		entity := models.Profile{}
		rows.Scan(&entity.ID, &entity.UserID, &entity.FullName, &entity.Avatar, &entity.Bio, &entity.CreatedAt, &entity.UpdatedAt)
		entities[index] = &entity
		index++
	}
	return entities[:index], nil
}
