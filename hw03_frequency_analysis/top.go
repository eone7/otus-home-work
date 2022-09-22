package hw03frequencyanalysis

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var ErrInvalidWordForRegexp = errors.New("panic word for regexp")

func getRegexpWord(w string, re *regexp.Regexp) string {
	if word := re.ReplaceAllString(w, ""); word != "" {
		return strings.ToLower(word)
	}
	return ""
}

func Top10(s string) []string {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(ErrInvalidWordForRegexp)
			return
		}
	}()

	const limitTop = 10
	counterWords := make(map[string]int)
	var uniqWords []string

	re := regexp.MustCompile(`(^'|'$|'[?=,]|[?<=,]')|([,!.]$)|^-$`)

	for _, w := range strings.Fields(s) {
		word := getRegexpWord(w, re)
		if word == "" {
			continue
		}

		if _, ok := counterWords[word]; !ok {
			uniqWords = append(uniqWords, word)
		}
		counterWords[word]++
	}

	sort.Slice(uniqWords, func(i, j int) bool {
		firstWord := uniqWords[i]
		secWord := uniqWords[j]
		if counterWords[firstWord] == counterWords[secWord] {
			return firstWord < secWord
		}
		return counterWords[firstWord] > counterWords[secWord]
	})

	if len(uniqWords) > limitTop {
		return uniqWords[:limitTop]
	}

	return uniqWords
}
