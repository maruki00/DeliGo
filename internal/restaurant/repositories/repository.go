package repositories

import (
	"errors"

	"github.com/maruki00/deligo/internal/restaurant/models"
	"gorm.io/gorm"
)

type CatalogRepository interface {
	CreateRestaurant(res *models.Restaurant) error
	GetRestaurantByID(id string) (*models.Restaurant, error)
	UpdateRestaurant(res *models.Restaurant) error

	CreateProduct(prod *models.Product) error
	GetProductByID(id string) (*models.Product, error)
	UpdateProduct(prod *models.Product) error
	DeleteProduct(id string) error

	GetRestaurantMenu(restaurantID string) (*models.Restaurant, error)
}

type catalogRepo struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepo{db: db}
}

func (r *catalogRepo) CreateRestaurant(res *models.Restaurant) error {
	return r.db.Create(res).Error
}

func (r *catalogRepo) GetRestaurantByID(id string) (*models.Restaurant, error) {
	var res models.Restaurant
	err := r.db.First(&res, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &res, err
}

func (r *catalogRepo) UpdateRestaurant(res *models.Restaurant) error {
	return r.db.Save(res).Error
}

func (r *catalogRepo) CreateProduct(prod *models.Product) error {
	return r.db.Create(prod).Error
}

func (r *catalogRepo) GetProductByID(id string) (*models.Product, error) {
	var prod models.Product
	err := r.db.First(&prod, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &prod, err
}

func (r *catalogRepo) UpdateProduct(prod *models.Product) error {
	return r.db.Save(prod).Error
}

func (r *catalogRepo) DeleteProduct(id string) error {
	return r.db.Delete(&models.Product{}, "id = ?", id).Error
}

func (r *catalogRepo) GetRestaurantMenu(restaurantID string) (*models.Restaurant, error) {
	var res models.Restaurant
	err := r.db.Preload("Products", "is_available = ?", true).First(&res, "id = ?", restaurantID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &res, err
}
