package main

import (
	"fmt"
)

func initFp() {
	fp = make(map[string]func([]string, net.Conn)error)
	fp["list"] = List
	fp["get"] = Get
	fp["put"] = Put
	fp["delete"] = Delete
}

func parseEvent(items []string) {

}

func List(item []string) error {
	return nil
}

func Get(item []string) error {
	return nil
}

func Put(item []string) error {
	return nil
}

func Delete(item []string) error {
	return nil
}





