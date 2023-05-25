package pkg

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (d *UpdateDataService) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Data in service [%s] is already update.\n", d.Name)
		d.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from service [%s].\n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *UpdateDataService) SetNext(service Service) {
	d.Next = service
}
