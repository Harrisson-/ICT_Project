package main

import (
	"os"
	"net"
	"fmt"
	"bufio"
)

var fp map[string]func([]string, net.Conn)error


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4243")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	initFp()
	for {
		reader := bufio.NewReader(os.Stdin)
		cmd, _ := reader.ReadString('\n')
		if len(cmd) > 1 {
			sendEvent(cmd, conn)
		}
	}
	

}
