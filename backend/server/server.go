package main

import (
	"net"
	"fmt"
)

func handleConnection(net.Conn) {

}

fun main() {
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
