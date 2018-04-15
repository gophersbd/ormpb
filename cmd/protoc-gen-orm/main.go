package main

import "fmt"

func main() {
	fmt.Println(Dummy())
}

// Dummy exists to generate test coverage
func Dummy() string {
	return "Hello I am going to be ORM. Please Code me up"
}
