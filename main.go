package main

import "fmt"

func main() {
	fmt.Println(GeneratePassword(8, 1, true, true, false))
	fmt.Println(GeneratePassword(12, 123, true, true, false))
	fmt.Println(GeneratePassword(12, 123, true, true, true))
	fmt.Println(GeneratePassword(8, 1, false, false, false))
	fmt.Println(GeneratePassword(-3, 42, true, true, false))
	fmt.Println(CheckPassword("abc"))
	fmt.Println(CheckPassword("abcdefgh"))
	fmt.Println(CheckPassword("abcdef1234"))
	fmt.Println(CheckPassword("Abcdef1234"))
	fmt.Println(CheckPassword("Abcdef123!"))
}
