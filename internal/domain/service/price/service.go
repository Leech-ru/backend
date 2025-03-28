package price

// CalculateTotal calculates the total cost of the order
func (s *priceService) CalculateTotal(size1, size2, size3, packageType int) (float64, error) {
	total := size1*s.leechPrices[1] +
		size2*s.leechPrices[2] +
		size3*s.leechPrices[3]

	return float64(total) / 100, nil
}
