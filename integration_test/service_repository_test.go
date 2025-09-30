package integration_test

import (
	"testing"

	"github.com/marcopollivier/poc.test-pyramid/model"
	"github.com/marcopollivier/poc.test-pyramid/repository"
	"github.com/marcopollivier/poc.test-pyramid/service"
)

func TestRepositoryIntegration_SaveAndRetrieve(t *testing.T) {
	db := setupTestDB(t)
	
	// Auto migrate the table
	err := db.AutoMigrate(&model.Discount{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
	
	repo := repository.NewDiscountRepository(db)
	calculator := service.NewPriceCalculator()

	tests := []struct {
		name     string
		price    float64
		discount float64
		want     float64
	}{
		{"10% discount", 100.0, 10.0, 90.0},
		{"50% discount", 200.0, 50.0, 100.0},
		{"0% discount", 150.0, 0.0, 150.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Calculate discount using service
			finalPrice, err := calculator.Calculate(tt.price, tt.discount)
			if err != nil {
				t.Fatalf("Calculate() error = %v", err)
			}

			if finalPrice != tt.want {
				t.Errorf("Calculate() = %v, want %v", finalPrice, tt.want)
			}

			// Save to database using repository
			discount := &model.Discount{
				Price:      tt.price,
				Discount:   tt.discount,
				FinalPrice: finalPrice,
			}

			err = repo.Save(discount)
			if err != nil {
				t.Fatalf("Save() error = %v", err)
			}

			// Verify it was saved in database
			var saved model.Discount
			err = db.Where("price = ? AND discount = ?", tt.price, tt.discount).First(&saved).Error
			if err != nil {
				t.Fatalf("Failed to find saved discount: %v", err)
			}

			if saved.FinalPrice != tt.want {
				t.Errorf("Saved discount final price = %v, want %v", saved.FinalPrice, tt.want)
			}
		})
	}
}
