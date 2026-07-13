package main

import "fmt"

func main() {
	fmt.Println(GeneratePassword(30, true, true))
	fmt.Println(GeneratePassword(30, false, false))
}
