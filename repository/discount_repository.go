package repository

import (
	"log"

	"github.com/marcopollivier/poc.test-pyramid/model"
	"gorm.io/gorm"
)

type DiscountRepository struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) *DiscountRepository {
	log.Println("DiscountRepository: Initializing repository")
	return &DiscountRepository{db: db}
}

func (r *DiscountRepository) Save(discount *model.Discount) error {
	log.Printf("DiscountRepository: Saving discount - Price: %.2f, Discount: %.2f%%, Final: %.2f", 
		discount.Price, discount.Discount, discount.FinalPrice)
	
	err := r.db.Create(discount).Error
	if err != nil {
		log.Printf("DiscountRepository: Save error - %v", err)
		return err
	}
	
	log.Printf("DiscountRepository: Successfully saved with ID: %d", discount.ID)
	return nil
}

func (r *DiscountRepository) FindAll() ([]model.Discount, error) {
	log.Println("DiscountRepository: Finding all discounts")
	
	var discounts []model.Discount
	err := r.db.Find(&discounts).Error
	if err != nil {
		log.Printf("DiscountRepository: FindAll error - %v", err)
		return nil, err
	}
	
	log.Printf("DiscountRepository: Found %d discounts", len(discounts))
	return discounts, err
}
