package price

// CalculateTotal calculates the total cost of the order
func (s *priceService) CalculateTotal(size1, size2, size3, packageType int) (float64, error) {
	total := float64(size1)*s.leechPrices[1] +
		float64(size2)*s.leechPrices[2] +
		float64(size3)*s.leechPrices[3]

	return total, nil
}
