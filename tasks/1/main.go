package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

type AppTime struct {
	host string
}

func (a *AppTime) Run() error {
	tNTP, err := ntp.Time(a.host)
	if err != nil {
		return err
	}

	fmt.Println(tNTP.UTC().Format(time.UnixDate))
	return nil
}

func main() {

}
