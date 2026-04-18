package main

import "fmt"

func GetMessage() string {
	return "Hello"
}

func PrintHello() {
	fmt.Println(GetMessage())
}
