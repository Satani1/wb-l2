package pkg

import "fmt"

type Server struct {
	Type   string
	Core   int
	Memory int
}

func (pc Server) GetType() string {
	return pc.Type
}

func (pc Server) PrintDetails() {
	fmt.Printf("%s Core:[%d] Memory:[%d]\n", pc.Type, pc.Core, pc.Memory)
}

func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   32,
		Memory: 128,
	}
}
