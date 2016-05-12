package main

import (
	"net"
	"fmt"
	"errors"
	"strings"
)

func initFp() {
	fp = make(map[string]func([]string, net.Conn)error)
	fp["list"] = List
	fp["get"] = Get
	fp["put"] = Put
	fp["delete"] = Delete
}

func sendEvent(cmd string, c net.Conn) {
	cmds := strings.Fields(cmd)
	act := strings.ToLower(cmds[0])
	if f, ok := fp[act]; ok {
		err := f(cmds[1:], c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func List(item []string, c net.Conn) error {
	_, err := c.Write([]byte("list\n"))
	if err != nil {
		return err
	}
	fmt.Println("List request sent")
	return nil
}
 
func Get(item []string, c net.Conn) error {
	if len(item) < 1 {
		return errors.New("Usage: get FILE...")
	}
	c.Write([]byte("get "+strings.Join(item, " ")+"\n"))
	fmt.Println("Get request sent")
	return nil
}

func Put(item []string, c net.Conn) error {
	if len(item) < 1 {
		return errors.New("Usage: put FILE...")
	}
	c.Write([]byte("put "+strings.Join(item, " ")+"\n"))
	fmt.Println("Put request sent")
	return nil
}

func Delete(item []string, c net.Conn) error {
	if len(item) < 1 {
		return errors.New("Usage: delete FILE...")
	}
	c.Write([]byte("delete "+strings.Join(item, " ")+"\n"))
	fmt.Println("Delete request sent")
	return nil
}
