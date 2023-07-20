package main

import (
	"bytes"
	"encoding/binary"
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
	fmt.Println(0, "connect ok! sending file...")

	for i := 0; i < 1000; i++ {
		msg := fmt.Sprintf("hello easy net NO: %d ", i)
		n, err := conn.Write(pack(msg))
		fmt.Println("Write msg", string(msg), n, err)
	}

	time.Sleep(time.Second * 2)

	var readstr = make([]byte, 100000)
	n, err := conn.Read(readstr)
	fmt.Println("read msg:\n ", string(readstr), n, err)

	time.Sleep(time.Second * 60)

	return
}

func IntToBytes(n int) []byte {
	data := int16(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func pack(msg string) []byte {
	var msgBytes = []byte{}
	mlen := len(msg)
	mlenBytes := IntToBytes(mlen)
	msgBytes = append(msgBytes, mlenBytes...)
	msgBytes = append(msgBytes, []byte(msg)...)
	return msgBytes
}
