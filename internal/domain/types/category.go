package types

type Category int

const (
	CategoryMainFace Category = iota
	CategoryMainBody
	CategoryHairCare
	CategoryMen
	CategoryIntensive
	CategoryLuxury
	CategoryExclusive
	CategoryWholesale
	CategoryLeech
)

func (c Category) String() string {
	switch c {
	case CategoryMainFace:
		return "Main care (face and décolleté)"
	case CategoryMainBody:
		return "Main care (body)"
	case CategoryHairCare:
		return "Hair care"
	case CategoryMen:
		return "Men's line"
	case CategoryIntensive:
		return "Intensive care"
	case CategoryLuxury:
		return "Luxury"
	case CategoryExclusive:
		return "Exclusive"
	case CategoryWholesale:
		return "Wholesale and retail cosmetics"
	case CategoryLeech:
		return "Leech-based cosmetics"
	default:
		return "Unknown category"
	}
}
