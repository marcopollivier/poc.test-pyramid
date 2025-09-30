package repository

import (
	"github.com/marcopollivier/poc.test-pyramid/model"
	"gorm.io/gorm"
)

type DiscountRepository struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) *DiscountRepository {
	return &DiscountRepository{db: db}
}

func (r *DiscountRepository) Save(discount *model.Discount) error {
	return r.db.Create(discount).Error
}

func (r *DiscountRepository) FindAll() ([]model.Discount, error) {
	var discounts []model.Discount
	err := r.db.Find(&discounts).Error
	return discounts, err
}
