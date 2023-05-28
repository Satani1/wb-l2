package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go TelnetConnection(conn)
	}
}

func TelnetConnection(conn net.Conn) {
	defer conn.Close()
	errChan := make(chan error)

	go func() {
		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				errChan <- err
				return
			}
			log.Printf("Got: %s\n", data)
			fmt.Fprintf(conn, "returnin message %s\n", data)
		}
	}()

	if err := <-errChan; err == io.EOF {
		log.Println("connection dropped")
	} else {
		log.Printf("error: %v\n", err)
	}
}
