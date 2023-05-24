package pkg

import "fmt"

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core:[%d] Memory:[%d] Monitor:[%v] \n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    15,
		Memory:  64,
		Monitor: true,
	}
}
