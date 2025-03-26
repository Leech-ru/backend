package price

type priceService struct {
	leechPrices map[int]float64
}

func NewPriceService() *priceService {
	return &priceService{
		leechPrices: map[int]float64{
			1: 105.0,
			2: 135.0,
			3: 168.0,
		},
	}
}
