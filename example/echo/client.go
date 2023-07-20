package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var (
		host   = "127.0.0.1"
		port   = "9011"
		remote = host + ":" + port
	)

	fmt.Println(remote)
	conn, err := net.Dial("tcp", remote)
	defer conn.Close()

	if err != nil {
		fmt.Println("connect server failed!.")
		os.Exit(-1)
		return
	}

	msg := "hello easy net "

	fmt.Println(0, "connect ok! sending file...")
	n, err := conn.Write([]byte(msg))
	fmt.Println("Write msg", string(msg), n, err)

	time.Sleep(time.Second * 2)

	var readstr = make([]byte, len(msg))
	n, err = conn.Read(readstr)
	fmt.Println("read msg:\n ", string(readstr), n, err)

	time.Sleep(time.Second * 60)

}
