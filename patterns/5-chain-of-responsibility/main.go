package main

import "chain/pkg"

func main() {
	device1 := &pkg.Device{Name: "device-1"}
	updateSrv := &pkg.UpdateDataService{Name: "update-1"}
	saveSrv := &pkg.DataService{}

	device1.SetNext(updateSrv)
	updateSrv.SetNext(saveSrv)

	data := &pkg.Data{}
	device1.Execute(data)
}
