package main

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"
const special = "!@#$%^&*"

func NextRandom(number int) int {
	return (16807 * number) % 2147483647
}

func GeneratePassword(length, seed int, useUppercase, useDigits, useSpecial bool) string {
	if length <= 0 {
		return ""
	}
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

	current := seed
	for i := 0; i < length; i++ {
		current = NextRandom(current)
		index := current % len
		result = result + string(alphabet[index])
	}
	return result
}
