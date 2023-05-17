package main

import (
	"flag"
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

type AppTime struct {
	host string
}

func (a *AppTime) Run() error {
	t, err := ntp.Time(a.host)
	if err != nil {
		return err
	}

	fmt.Println(t.UTC().Format(time.UnixDate))
	return nil
}

func main() {
	host := flag.String("host", "0.beevik-ntp.pool.ntp.org", "host to get time")
	flag.Parse()

	aTime := AppTime{host: *host}
	if err := aTime.Run(); err != nil {
		log.Fatalln(err)
	}
}
