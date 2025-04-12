package repositories

import (
	"context"
	"delivery/internal/product/domian/entities"
	"delivery/internal/product/infrastructure/models"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db    *gorm.DB
	model interface{}
}

func NewProductRepository(db *gorm.DB, model interface{}) *ProductRepository {
	return &ProductRepository{
		db:    db,
		model: model,
	}
}

func (obj *ProductRepository) Insert(ctx context.Context, product entities.ProductEntity) (entities.ProductEntity, error) {
	res := obj.db.Model(&models.Product{}).Create(product)
	fmt.Println("product : ", product)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("record could not be saved")
	}
	return product, nil
}

func (obj *ProductRepository) GetById(ctx context.Context, id int) (entities.ProductEntity, error) {

	product := &models.Product{}
	obj.db.Model(&models.Product{}).Where("id = ?", id).Find(&product)
	if product == nil {
		return nil, fmt.Errorf("record could not be found")
	}
	if product.Id == 0 {
		return nil, nil
	}

	return product, nil
}

func (obj *ProductRepository) Search(ctx context.Context, seasrch string) ([]models.Product, error) {
	items := []models.Product{}
	res := obj.db.Model(obj.model).Where("( label like  ?  or price like  ?  or type like  ? )", fmt.Sprintf("%%%s", seasrch), fmt.Sprintf("%%%s%%", seasrch), fmt.Sprintf("%%%s%%", seasrch)).Limit(3).Offset(0).Find(&items)
	if res.Error != nil {
		return []models.Product{}, nil
	}
	return items, nil
}

func (obj *ProductRepository) Update(ctx context.Context, id int, data map[string]interface{}) (entities.ProductEntity, error) {

	res := obj.db.Model(&obj.model).Where("id", id).Updates(data)
	if res.Error != nil {
		return nil, fmt.Errorf("something happen, %v", res.Error)
	}

	product, err := obj.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (obj *ProductRepository) Delete(ctx context.Context, id int) (entities.ProductEntity, error) {

	product, err := obj.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("product not found")
	}
	res := obj.db.Model(obj.model).Where("id = ? ", id).Delete(&product)
	if res.Error != nil {
		return nil, fmt.Errorf("something happen %v", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("could not delete the record")
	}
	return product, nil
}

func (obj *ProductRepository) GetProductByMultipleId(ctx context.Context, ids []int) ([]models.Product, error) {
	products := []models.Product{}

	query := ""

	for _, id := range ids {
		query += fmt.Sprintf(" %d,", id)
	}
	query = strings.TrimSuffix(query, ",")

	fmt.Println(query)
	res := obj.db.Model(obj.model).Where(" id in ( " + query + " )").Find(&products)

	if res.Error != nil {
		fmt.Println(res.Error.Error())
		return nil, res.Error
	}
	return products, nil
}
