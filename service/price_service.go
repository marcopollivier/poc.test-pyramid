package service

import (
	"github.com/marcopollivier/poc.test-pyramid/messaging"
	"github.com/marcopollivier/poc.test-pyramid/model"
	"github.com/marcopollivier/poc.test-pyramid/repository"
)

type PriceService struct {
	calculator *PriceCalculator
	repo       *repository.DiscountRepository
	publisher  *messaging.KafkaPublisher
}

func NewPriceService(repo *repository.DiscountRepository, publisher *messaging.KafkaPublisher) *PriceService {
	return &PriceService{
		calculator: NewPriceCalculator(),
		repo:       repo,
		publisher:  publisher,
	}
}

func (ps *PriceService) CalculateDiscount(price float64, discountPercent float64) (float64, error) {
	finalPrice, err := ps.calculator.Calculate(price, discountPercent)
	if err != nil {
		return 0, err
	}
	
	// Salvar no banco
	discountModel := &model.Discount{
		Price:      price,
		Discount:   discountPercent,
		FinalPrice: finalPrice,
	}
	
	if err := ps.repo.Save(discountModel); err != nil {
		return 0, err
	}
	
	// Publicar no Kafka
	if err := ps.publisher.PublishDiscount(discountModel); err != nil {
		return 0, err
	}
	
	return finalPrice, nil
}
