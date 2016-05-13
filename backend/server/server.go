package main

import (
	"bufio"
	"net"
	"fmt"
)

var fp map[string]func([]string, net.Conn)error

func handleConnection(conn net.Conn) {
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
