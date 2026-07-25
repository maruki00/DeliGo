package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maruki00/deligo/internal/restaurant/models"
	"github.com/maruki00/deligo/internal/restaurant/repositories"
	"github.com/maruki00/deligo/internal/restaurant/requests"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized ownership context verified failure")
)

type CatalogService interface {
	CreateRestaurant(ownerID string, req requests.CreateRestaurantRequest) (*models.Restaurant, error)
	ToggleRestaurantStatus(ownerID string, resID string, isOpen bool) error
	AddProduct(ownerID string, resID string, req requests.CreateProductRequest) (*models.Product, error)
	UpdateProduct(ownerID string, prodID string, req requests.UpdateProductRequest) (*models.Product, error)
	DeleteProduct(ownerID string, prodID string) error
	GetMenu(resID string) (*models.Restaurant, error)
}

type catalogService struct {
	repo repositories.CatalogRepository
}

func NewCatalogService(repo repositories.CatalogRepository) CatalogService {
	return &catalogService{repo: repo}
}

func (s *catalogService) CreateRestaurant(ownerID string, req requests.CreateRestaurantRequest) (*models.Restaurant, error) {
	res := &models.Restaurant{
		ID:      uuid.New().String(),
		OwnerID: ownerID,
		Name:    req.Name,
		Address: req.Address,
		IsOpen:  false,
	}
	if err := s.repo.CreateRestaurant(res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *catalogService) ToggleRestaurantStatus(ownerID string, resID string, isOpen bool) error {
	res, err := s.repo.GetRestaurantByID(resID)
	if err != nil {
		return err
	}
	if res == nil {
		return ErrNotFound
	}
	if res.OwnerID != ownerID {
		return ErrUnauthorized
	}

	res.IsOpen = isOpen
	return s.repo.UpdateRestaurant(res)
}

func (s *catalogService) AddProduct(ownerID string, resID string, req requests.CreateProductRequest) (*models.Product, error) {
	res, err := s.repo.GetRestaurantByID(resID)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, ErrNotFound
	}
	if res.OwnerID != ownerID {
		return nil, ErrUnauthorized
	}

	prod := &models.Product{
		ID:           uuid.New().String(),
		RestaurantID: resID,
		Name:         req.Name,
		Description:  req.Description,
		Price:        req.Price,
		IsAvailable:  true,
	}
	if err := s.repo.CreateProduct(prod); err != nil {
		return nil, err
	}
	return prod, nil
}

func (s *catalogService) UpdateProduct(ownerID string, prodID string, req requests.UpdateProductRequest) (*models.Product, error) {
	prod, err := s.repo.GetProductByID(prodID)
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, ErrNotFound
	}

	res, err := s.repo.GetRestaurantByID(prod.RestaurantID)
	if err != nil {
		return nil, err
	}
	if res == nil || res.OwnerID != ownerID {
		return nil, ErrUnauthorized
	}

	prod.Name = req.Name
	prod.Description = req.Description
	prod.Price = req.Price
	prod.IsAvailable = *req.IsAvailable

	if err := s.repo.UpdateProduct(prod); err != nil {
		return nil, err
	}
	return prod, nil
}

func (s *catalogService) DeleteProduct(ownerID string, prodID string) error {
	prod, err := s.repo.GetProductByID(prodID)
	if err != nil {
		return err
	}
	if prod == nil {
		return ErrNotFound
	}

	res, err := s.repo.GetRestaurantByID(prod.RestaurantID)
	if err != nil {
		return err
	}
	if res == nil || res.OwnerID != ownerID {
		return ErrUnauthorized
	}

	return s.repo.DeleteProduct(prodID)
}

func (s *catalogService) GetMenu(resID string) (*models.Restaurant, error) {
	res, err := s.repo.GetRestaurantMenu(resID)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}
