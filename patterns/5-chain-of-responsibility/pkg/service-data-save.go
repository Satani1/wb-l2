package pkg

import "fmt"

type DataService struct {
	Next Service
}

func (d *DataService) Execute(data *Data) {
	if !data.GetSource {
		fmt.Printf("Data not update.\n")
		d.Next.Execute(data)
		return
	}
	fmt.Printf("Data save.\n")
}

func (d *DataService) SetNext(service Service) {
	d.Next = service
}
