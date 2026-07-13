### Hexlet tests and linter status:
[![Actions Status](https://github.com/LonelySadhu/go-basics-project-394/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/LonelySadhu/go-basics-project-394/actions)

## О проекте

Консольная утилита на Go для генерации паролей и оценки их надёжности.

- `GeneratePassword(length, seed int, useUppercase, useDigits, useSpecial bool) string` — генерирует пароль заданной длины на основе псевдослучайного генератора (линейный конгруэнтный метод). Алфавит всегда включает строчные буквы, дополнительно можно подключить заглавные буквы, цифры и спецсимволы (`!@#$%^&*`).
- `CheckPassword(password string) string` — оценивает сложность пароля по 5-балльной шкале (на основе длины и разнообразия использованных категорий символов) и возвращает текстовый вердикт: "Слабый", "Средний", "Надежный" или "Очень надежный".

## Запуск

Требуется установленный Go (версия указана в `go.mod`).

```bash
go run .
```

## Сборка

```bash
go build -o password-generator .
./password-generator
```