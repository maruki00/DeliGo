package repositories

import (
	"context"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/domain/entities"
	pkgPostgres "delivery/pkg/postgres"
	"sync"
)

type UserRepository struct {
	sync.RWMutex
	db pkgPostgres.PGHandler
}

func NewUserRepository(db pkgPostgres.PGHandler) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(ctx context.Context,entity entities.UserEntity) (entities.UserEntity, error){
	sql := `INSERT INTO users(email, password, role) VALUES($1, $2, $3) RETURNING id`
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	var lastInsertID string
	stmnt, err := tx.QueryRow(sql, entity.GetEmail(), entity.GetPassword(), entity.GetRole()).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}
	defer stmnt.Close()
	if err := tx.Commit(); err!= nil {
		tx.Rollback()
		return nil, err
	}
	entity.SetID(lastInsertID)
	return entity, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id string) (bool, error){
	//sql := `DELETE FROM users WHERE id = $1`
	sql := `UPDATE users SET deleted_at = now(), updated_at = now() WHERE id = $1 `
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return false, err
	}
	var lastInsertID string
	stmnt, err := tx.QueryRow(sql, id)
	if err != nil {
		return false, err
	}
	defer stmnt.Close()
	if err := tx.Commit(); err!= nil {
		tx.Rollback()
		return false, err
	}
	return true, nil
}

func (ur *UserRepository) Update(ctx context.Context, entity entities.UserEntity) (entities.UserEntity, error){

	sql := `UPDATE users SET email=$1, role=$2, updated_at = now() WHERE id=$3 AND deleted_at = NULL`
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return nil, err
	}

	stmnt, err := tx.QueryRow(sql, entity.GetEmail(), entity.GetRole(), entity.GetID)
	if err != nil {
		return nil, err
	}
	defer stmnt.Close()
	if err := tx.Commit(); err!= nil {
		tx.Rollback()
		return nil, err
	}

	return entity, nil
}

func (ur *UserRepository) GetOne(ctx context.Context, id string) (entities.UserEntity, error){
	var entity models.User
	sql := `
			SELECT id, email, role 
			FROM users 
			WHERE id = $1 AND deleted_at = NULL 
			LIMIT 1 
			RETURNING id,email,role,created_at,updated_at
		`
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	var lastInsertID string
	stmnt, err := tx.QueryRow(sql, id).Scan(&entity.ID, &entity.Email, &entity.role, &entity.CreatedAt, &entity.UpdatedAt)
	if err != nil {
		return nil, err
	}
	defer stmnt.Close()
	if err := tx.Commit(); err!= nil {
		tx.Rollback()
		return nil, err
	}
	return &entity, nil
}

func (ur *UserRepository) GetMany(ctx context.Context, page, offset int)([]entities.UserEntity, error){
	var entity *[]*models.User
	sql := `
			SELECT id, email, role 
			FROM users 
			WHERE deleted_at = NULL 
			LIMIT $1, $2
		`
	tx, err := ur.db.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	var lastInsertID string
	stmnt, err := tx.QueryRow(sql, id).Scan(&entity.ID, &entity.Email, &entity.role, &entity.CreatedAt, &entity.UpdatedAt)
	if err != nil {
		return nil, err
	}
	defer stmnt.Close()
	if err := tx.Commit(); err!= nil {
		tx.Rollback()
		return nil, err
	}
	return &entity, nil
}

func (ur *UserRepository) Search(ctx context.Context, any) ([]entities.UserEntity, error){
	return nil, nil
}