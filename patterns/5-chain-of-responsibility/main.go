package main

import "chain/pkg"

func main() {
	device1 := &pkg.Device{Name: "device-dev01"}
	updateSrv := &pkg.UpdateDataService{Name: "update-dev01"}
	saveSrv := &pkg.DataService{}

	device1.SetNext(updateSrv)
	updateSrv.SetNext(saveSrv)

	data := &pkg.Data{}
	device1.Execute(data)
}
