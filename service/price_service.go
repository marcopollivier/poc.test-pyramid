package service

import (
	"log"

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
	log.Println("PriceService: Initializing service")
	return &PriceService{
		calculator: NewPriceCalculator(),
		repo:       repo,
		publisher:  publisher,
	}
}

func (ps *PriceService) CalculateDiscount(price float64, discountPercent float64) (float64, error) {
	log.Printf("PriceService: Calculating discount - Price: %.2f, Discount: %.2f%%", price, discountPercent)
	
	finalPrice, err := ps.calculator.Calculate(price, discountPercent)
	if err != nil {
		log.Printf("PriceService: Calculation error - %v", err)
		return 0, err
	}
	
	log.Printf("PriceService: Calculation result - Final price: %.2f", finalPrice)
	
	// Salvar no banco
	discountModel := &model.Discount{
		Price:      price,
		Discount:   discountPercent,
		FinalPrice: finalPrice,
	}
	
	log.Println("PriceService: Saving to database...")
	if err := ps.repo.Save(discountModel); err != nil {
		log.Printf("PriceService: Database save error - %v", err)
		return 0, err
	}
	log.Printf("PriceService: Saved to database with ID: %d", discountModel.ID)
	
	// Publicar no Kafka
	log.Println("PriceService: Publishing to Kafka...")
	if err := ps.publisher.PublishDiscount(discountModel); err != nil {
		log.Printf("PriceService: Kafka publish error - %v", err)
		return 0, err
	}
	log.Println("PriceService: Successfully published to Kafka")
	
	log.Printf("PriceService: Process completed successfully - Final price: %.2f", finalPrice)
	return finalPrice, nil
}
