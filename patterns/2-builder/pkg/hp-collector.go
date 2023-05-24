package pkg

type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     int
}

func (colector *HpCollector) SetCore() {
	colector.Core = 8
}

func (colector *HpCollector) SetBrand() {
	colector.Brand = "HP"
}

func (colector *HpCollector) SetMemory() {
	colector.Memory = 16
}

func (colector *HpCollector) SetGraphicCard() {
	colector.GraphicCard = 1
}

func (colector *HpCollector) SetMonitor() {
	colector.Monitor = 2
}

func (colector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        colector.Core,
		Memory:      colector.Memory,
		GraphicCard: colector.GraphicCard,
		Monitor:     colector.Monitor,
		Brand:       colector.Brand,
	}
}
