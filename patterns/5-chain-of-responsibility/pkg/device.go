package pkg

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Data from device [%s] already get.\n", d.Name)
		d.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s].\n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(service Service) {
	d.Next = service
}
