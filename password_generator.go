package main

import (
	"fmt"
	"strings"
)

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

func CheckPassword(password string) string {
	estimate := 0
	if strings.ContainsAny(password, lowercase) {
		estimate = 2
	}
	if len(password) <= 3 {
		estimate -= 1
	}
	if strings.ContainsAny(password, digits) {
		estimate++
	}
	if strings.ContainsAny(password, uppercase) {
		estimate++
	}
	if strings.ContainsAny(password, special) {
		estimate++
	}
	if estimate < 0 {
		estimate = 0
	}

	answer := fmt.Sprintf("пароль (оценка %d из 5)", estimate)
	switch {
	case estimate <= 1:
		return "Слабый " + answer
	case estimate == 2:
		return "Слабый " + answer
	case estimate == 3:
		return "Средний " + answer
	case estimate == 4:
		return "Надёжный " + answer
	default:
		return "Очень надёжный " + answer
	}
}
