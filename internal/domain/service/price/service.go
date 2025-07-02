package price

type priceService struct {
	leechPrices map[int]int
}

func NewPriceService() *priceService {
	return &priceService{
		leechPrices: map[int]int{
			1: 10500,
			2: 13500,
			3: 16800,
		},
	}
}
