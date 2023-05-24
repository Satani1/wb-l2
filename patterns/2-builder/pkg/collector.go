package pkg

const (
	AsusColletorType = "asus"
	HpCollectorType  = "hp"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetGraphicCard()
	SetMonitor()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusColletorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}
