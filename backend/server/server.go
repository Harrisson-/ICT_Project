package main

import (
	"io/ioutil"
	"net"
	"fmt"
)

func handleConnection(conn net.Conn) {
	// buf := make([]byte, 256)
	// for {
	// 	_, err := conn.Read(buf)
	// 	if err != nil {
	// 		conn.Close()
	// 		return
	// 	}
	// 	fmt.Println(buf)
	// }

	bufReader := bufio.NewReader(conn)

	for {
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", bytes)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":4243")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			go handleConnection(conn)
		}
	}
}
