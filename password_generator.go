package main

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"

func GeneratePassword(length int, useUppercase, useDigits bool) string {
	alphabet := lowercase
	if useUppercase {
		alphabet += uppercase
	} else if useDigits {
		alphabet += digits
	}
	len := len(alphabet)
	result := ""
	for i := 0; i < length; i++ {
		index := i % len
		result = result + string(alphabet[index])
	}
	return result
}
