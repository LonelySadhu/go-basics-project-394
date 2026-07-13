package main

import "fmt"

func main() {
	fmt.Println(GeneratePassword(8, 1, true, true, false))
	fmt.Println(GeneratePassword(12, 123, true, true, false))
	fmt.Println(GeneratePassword(12, 123, true, true, true))
	fmt.Println(GeneratePassword(8, 1, false, false, false))
}
