package price

func NewPriceService() *priceService {
	return &priceService{
		leechPrices: map[int]int{
			1: 10500,
			2: 13500,
			3: 16800,
		},
	}
}
