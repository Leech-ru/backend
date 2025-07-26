package types

type Package int

const (
	PackageWater Package = iota
	PackageGel
	PackagePeat // Торф
)

func (p Package) String() string {
	switch p {
	case PackageWater:
		return "Water"
	case PackageGel:
		return "Gel"
	case PackagePeat:
		return "Peat"
	default:
		return "unknown"
	}
}
