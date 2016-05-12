package main

import (
	"fmt"
)

func event(fp func(string)error, item string) {
	err := fp(item)
	if err != nil {
		fmt.Println(err)
	}
}

func parseEvent(cmds []string) {

}

func List(item string) error {
	return nil
}

func Get(item string) error {
	return nil
}

func Put(item string) error {
	return nil
}

func Delete(item string) error {
	return nil
}





