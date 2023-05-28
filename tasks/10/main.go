package main

import (
	"os"
	"telnet/telnet"
)

func main() {
	os.Exit(telnet.CLI(os.Args[1:]))
}
