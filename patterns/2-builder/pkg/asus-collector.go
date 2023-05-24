package pkg

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     int
}

func (colector *AsusCollector) SetCore() {
	colector.Core = 4
}

func (colector *AsusCollector) SetBrand() {
	colector.Brand = "Asus"
}

func (colector *AsusCollector) SetMemory() {
	colector.Memory = 8
}

func (colector *AsusCollector) SetGraphicCard() {
	colector.GraphicCard = 1
}

func (colector *AsusCollector) SetMonitor() {
	colector.Monitor = 1
}

func (colector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        colector.Core,
		Memory:      colector.Memory,
		GraphicCard: colector.GraphicCard,
		Monitor:     colector.Monitor,
		Brand:       colector.Brand,
	}
}
