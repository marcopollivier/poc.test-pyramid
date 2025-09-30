package service

import "errors"

type PriceService struct{}

func NewPriceService() *PriceService {
	return &PriceService{}
}

func (ps *PriceService) CalculateDiscount(price float64, discountPercent float64) (float64, error) {
	if price < 0 {
		return 0, errors.New("price cannot be negative")
	}
	if discountPercent < 0 || discountPercent > 100 {
		return 0, errors.New("discount must be between 0 and 100")
	}
	
	discount := price * (discountPercent / 100)
	return price - discount, nil
}
