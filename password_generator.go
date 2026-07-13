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
	if len(password) == 0 {
		return ""
	}
	for _, s := range lowercase {
		if strings.ContainsRune(password, s) {
			estimate = 2
			break
		}
	}
	if len(password) <= 3 {
		estimate -= 1
	}

	for _, s := range digits {
		if strings.ContainsRune(password, s) {
			estimate += 1
			break
		}
	}
	for _, s := range uppercase {
		if strings.ContainsRune(password, s) {
			estimate += 1
			break
		}
	}
	for _, s := range special {
		if strings.ContainsRune(password, s) {
			estimate += 1
			break
		}
	}
	answer := fmt.Sprintf("пароль (оценка %d из 5)", estimate)
	switch estimate {
	case 1:
		return "Слабый " + answer
	case 2:
		return "Слабый " + answer
	case 3:
		return "Средний " + answer
	case 4:
		return "Надежный " + answer
	case 5:
		return "Очень надежный " + answer
	default:
		return ""
	}
}
