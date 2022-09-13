package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	ErrInvalidString = errors.New("invalid string")
	ErrAtoi          = errors.New("atoi error")
)

var isTwoLetters = func(fc, sc rune) bool {
	return !unicode.IsDigit(fc) && !unicode.IsDigit(sc)
}

var isTwoDigits = func(fc, sc rune) bool {
	return unicode.IsDigit(fc) && unicode.IsDigit(sc)
}

func Unpack(s string) (string, error) {
	/*
	   Функция распаковки строки по след алгоритму:
	   при проходе по строке каждый элемент смотрит на своего соседа,
	   если сосед цифра то данный элемент умножается на цифру
	   остальные элементы остаются прежними
	   недопускаются числа в строке и не допускается начало строки с цифры
	*/

	// первый элемент цифра
	if len(s) > 0 && unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}

	builder := strings.Builder{}
	var w int

	for i := 0; i < len(s)-1; i += w {
		firstRune, width := utf8.DecodeRuneInString(s[i:])
		w = width

		secondRune, _ := utf8.DecodeRuneInString(s[i+w:])

		if isTwoDigits(firstRune, secondRune) {
			return "", ErrInvalidString
		}

		if isTwoLetters(firstRune, secondRune) {
			builder.WriteString(string(firstRune))
			continue
		}

		if unicode.IsDigit(secondRune) {
			number, err := strconv.Atoi(string(secondRune))
			if err != nil {
				return "", ErrAtoi
			}

			repeatChars := strings.Repeat(string(firstRune), number)
			builder.WriteString(repeatChars)
			continue
		}
	}

	// запись последней буквы
	if len(s) >= 1 && !unicode.IsDigit(rune(s[len(s)-1])) {
		builder.WriteString(string(s[len(s)-1]))
	}

	return builder.String(), nil
}
