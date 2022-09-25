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

type mapCounter map[string]int

type WordsCounter struct {
	words   []string
	counter mapCounter
}

func NewWordsCounter() *WordsCounter {
	return &WordsCounter{words: nil, counter: make(mapCounter, 0)}
}

func (wc *WordsCounter) AddWord(w string) {
	if _, ok := wc.counter[w]; !ok {
		wc.words = append(wc.words, w)
	}
	wc.counter[w]++
}

func (wc *WordsCounter) GetMostCommonWords() []string {
	sort.Slice(wc.words, func(i, j int) bool {
		firstWord := wc.words[i]
		secWord := wc.words[j]
		if wc.counter[firstWord] == wc.counter[secWord] {
			return firstWord < secWord
		}
		return wc.counter[firstWord] > wc.counter[secWord]
	})

	if len(wc.words) > limitTop {
		return wc.words[:limitTop]
	}

	return wc.words
}

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

	counter := NewWordsCounter()
	re := regexp.MustCompile(`(^'|'$|'[?=,]|[?<=,]')|([,!.]$)|^-$`)

	for _, w := range strings.Fields(s) {
		// getWord
		word := getRegexpWord(w, re)
		if word == "" {
			continue
		}
		counter.AddWord(word)
	}

	return counter.GetMostCommonWords()
}
