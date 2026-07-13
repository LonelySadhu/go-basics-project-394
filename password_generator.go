package main

const lowercase = "abcdefghijklmnopqrstuvwxyz"

func GeneratePassword(length int) string {
	alphabet := lowercase
	len := len(alphabet)
	result := ""
	for i := 0; i < length; i++ {
		index := i % len
		result = result + string(alphabet[index])
	}
	return result
}
