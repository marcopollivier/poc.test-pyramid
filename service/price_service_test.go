package service

import (
	"testing"
)

func TestPriceService_CalculateDiscount(t *testing.T) {
	ps := NewPriceService()

	tests := []struct {
		name            string
		price           float64
		discountPercent float64
		expectedPrice   float64
		expectError     bool
	}{
		{
			name:            "valid discount 10%",
			price:           100.0,
			discountPercent: 10.0,
			expectedPrice:   90.0,
			expectError:     false,
		},
		{
			name:            "valid discount 50%",
			price:           200.0,
			discountPercent: 50.0,
			expectedPrice:   100.0,
			expectError:     false,
		},
		{
			name:            "no discount",
			price:           100.0,
			discountPercent: 0.0,
			expectedPrice:   100.0,
			expectError:     false,
		},
		{
			name:            "negative price",
			price:           -100.0,
			discountPercent: 10.0,
			expectedPrice:   0.0,
			expectError:     true,
		},
		{
			name:            "invalid discount over 100%",
			price:           100.0,
			discountPercent: 150.0,
			expectedPrice:   0.0,
			expectError:     true,
		},
		{
			name:            "negative discount",
			price:           100.0,
			discountPercent: -10.0,
			expectedPrice:   0.0,
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ps.CalculateDiscount(tt.price, tt.discountPercent)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expectedPrice {
					t.Errorf("expected %f, got %f", tt.expectedPrice, result)
				}
			}
		})
	}
}
