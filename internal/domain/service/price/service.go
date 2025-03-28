package price

import "LutiLeech/internal/domain/common/errorz"

// CalculateTotal calculates the total cost of the order
func (s *priceService) CalculateTotal(size1, size2, size3, packageType int) (float64, error) {
	total := size1*s.leechPrices[1] +
		size2*s.leechPrices[2] +
		size3*s.leechPrices[3]
	if total < 0 || size1 < 0 || size2 < 0 || size3 < 0 {
		return 0, errorz.TooMuchLeeches
	}
	return float64(total) / 100, nil
}
