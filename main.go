package main

import "fmt"

var version = "dev"

func main() {
	fmt.Println("version: ", version)
	fmt.Println(sayHello())
}

func sayHello() string {
	return "Hello Golang!"
}
