package service

import "testing"

func TestDiscountService_CalculateDiscount(t *testing.T) {
	ds := &DiscountService{}
	tests := []struct {
		name   string
		amount int
		expect int
	}{
		{"no discount for 1000", 1000, 0},
		{"10% discount for 2000", 2000, 200},
		{"no discount for 0", 0, 0},
		{"10% discount for 1500", 1500, 150},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			discount := ds.CalculateDiscount(tt.amount)
			if discount != tt.expect {
				t.Errorf("amount=%d: expected discount %d, got %d", tt.amount, tt.expect, discount)
			}
		})
	}
}
