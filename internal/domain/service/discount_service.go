package service

// DiscountService provides discount calculation logic.
type DiscountService struct{}

// CalculateDiscount returns 10% discount for amount > 1000, otherwise 0.
func (s *DiscountService) CalculateDiscount(amount int) int {
	if amount > 1000 {
		return amount / 10 // 10% discount
	}
	return 0
}
