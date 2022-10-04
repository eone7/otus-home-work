package hw03frequencyanalysis

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var ErrInvalidWordForRegexp = errors.New("panic word for regexp")

const limitTop = 10

var re = regexp.MustCompile(`(^'|'$|'[?=,]|[?<=,]')|([,!.]$)|^-$`)

func getRegexpWord(w string) string {
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

	counterWords := make(map[string]int)
	var uniqWords []string

	for _, w := range strings.Fields(s) {
		word := getRegexpWord(w)
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
