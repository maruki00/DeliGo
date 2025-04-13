package repositories

import (
	"context"
	"delivery/internal/user/domain/entities"
	"delivery/internal/user/infra/models"
	pkgPostgres "delivery/pkg/postgres"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository struct {
	db pkgPostgres.PGHandler
}

func NewUserRepository(db pkgPostgres.PGHandler) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(ctx context.Context, entity entities.UserEntity) (entities.UserEntity, error) {

	sql := `INSERT INTO users(id, email, password, role) VALUES($1, $2, $3, $4) RETURNING id`
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	var lastInsertID string
	fmt.Println("UUID: ", uuid.New().String())
	err = tx.QueryRow(sql, uuid.New().String(), entity.GetEmail(), entity.GetPassword(), entity.GetRole()).Scan(&lastInsertID)
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

func (ur *UserRepository) Delete(ctx context.Context, id string) (bool, error) {
	sql := `UPDATE users SET deleted_at = now(), updated_at = now() WHERE id = $1 `
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

func (ur *UserRepository) Update(ctx context.Context, entity entities.UserEntity) (entities.UserEntity, error) {
	sql := `
			UPDATE users 
			SET email=$1, role=$2, updated_at = now() 
			WHERE id=$3 
			AND deleted_at = NULL
		`
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	_ = tx.QueryRow(sql, entity.GetEmail(), entity.GetRole(), entity.GetID())
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return entity, nil
}

func (ur *UserRepository) GetOne(ctx context.Context, userId string) (entities.UserEntity, error) {
	fmt.Println("result : ", userId)
	sql := `
			SELECT id,email,role
			FROM users 
			WHERE id = $1 
			-- AND deleted_at = NULL 
			LIMIT 1
		`

	var id, email, role = "", "", ""
	err := ur.db.GetDB().QueryRow(sql, userId).Scan(&id, &email, &role)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:    id,
		Email: email,
		Role:  role,
	}, nil
}

func (ur *UserRepository) GetMany(ctx context.Context, offset, limit int32) ([]*models.User, error) {
	entities := make([]*models.User, limit)
	offset = (offset - 1) * offset
	sql := `
			SELECT id,email,role
			FROM users 
			-- WHERE deleted_at = NULL 
			OFFSET $1
			LIMIT $2
			
		`
	rows, err := ur.db.GetDB().Query(sql, offset, limit)
	if err != nil {
		return nil, err
	}

	index := 0
	for rows.Next() && index < int(limit) {
		var id, email, role = "", "", ""
		rows.Scan(&id, &email, &role)
		entities[index] = &models.User{
			ID:    id,
			Email: email,
			Role:  role,
		}
		index++
	}
	return entities[:index], nil
}

func (ur *UserRepository) Search(ctx context.Context, query string, offset, limit int32) ([]*models.User, error) {
	sql := `
			SELECT id,email,role,created_at,updated_at
			FROM users 
			WHERE deleted_at = NULL 
			AND (email like '%$1' OR role like '%$1' or id like '%$1')
			LIMIT $2 
			OFFSET $3
		`

	entities := make([]*models.User, limit)

	offset = (offset - 1) * offset

	rows, err := ur.db.GetDB().Query(sql, query, offset, limit)
	if err != nil {
		return nil, err
	}

	index := 0
	for rows.Next() && index < int(limit) {
		var id, email, role = "", "", ""
		rows.Scan(&id, &email, &role)
		entities[index] = &models.User{
			ID:    id,
			Email: email,
			Role:  role,
		}
		index++
	}
	return entities[:index], nil
}
