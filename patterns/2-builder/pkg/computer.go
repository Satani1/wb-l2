package pkg

import "fmt"

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	GraphicCard int
	Monitor     int
}

func (pc Computer) Print() {
	fmt.Printf("%v Core:[%d] Mem:[%d] Graphic:[%d] Monitor:[%d]\n", pc.Brand, pc.Core, pc.Memory, pc.GraphicCard, pc.Monitor)
}
