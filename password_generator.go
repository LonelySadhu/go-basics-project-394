package main

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"
const special = "!@#$%^&*"

func GeneratePassword(length int, useUppercase, useDigits, useSpecial bool) string {
	alphabet := lowercase
	if useUppercase {
		alphabet += uppercase
	}
	if useDigits {
		alphabet += digits
	}
	if useSpecial {
		alphabet += special
	}
	len := len(alphabet)
	result := ""
	for i := 0; i < length; i++ {
		index := i % len
		result = result + string(alphabet[index])
	}
	return result
}
