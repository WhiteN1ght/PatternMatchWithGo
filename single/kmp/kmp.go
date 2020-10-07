package kmp

import (
	"errors"
	"fmt"
	"strings"
)

type Kmp struct {
	pattern     string
	patternLen  int
	next        []int
	nocaseMatch bool
}

func Init(pattern string, nocaseMatch ...bool) (*Kmp, error) {
	var nocase bool
	if pattern == "" {
		return nil, errors.New("[KMP] Can't match null string")
	}
	if nocaseMatch != nil && nocaseMatch[0] == true {
		nocase = true
		pattern = strings.ToLower(pattern)
	}

	patternLen := len(pattern)
	return &Kmp{
		pattern:     pattern,
		patternLen:  patternLen,
		next:        getNextTable(pattern, patternLen),
		nocaseMatch: nocase}, nil
}

func (kmp *Kmp) Match(text string) []int {
	textLen := len(text)
	patternLen := kmp.patternLen
	matchIdxs := make([]int, 0)

	if patternLen > textLen {
		return []int{}
	}

	i := 0
	j := 0
	for i < textLen {
		textChr := text[i]
		if kmp.nocaseMatch && textChr >= 'A' && textChr <= 'Z' {
			textChr += 'a' - 'A'
		}

		if j == patternLen-1 && textChr == kmp.pattern[j] {
			matchIdxs = append(matchIdxs, i-(patternLen-1))
			j = kmp.next[j]
		}
		if textChr == kmp.pattern[j] {
			i++
			j++
		} else {
			j = kmp.next[j]
			if j == -1 {
				i += 1
				j += 1
			}
		}
	}

	return matchIdxs
}

func (kmp *Kmp) ShowNext() {
	fmt.Println(kmp)
}

func getNextTable(pattern string, patternLen int) []int {
	var next = make([]int, len(pattern)+1)

	next[0] = 0
	i := 1
	prev := 0
	for i < patternLen {
		if pattern[i] == pattern[prev] {
			prev += 1
			next[i] = prev
			i += 1
		} else {
			if prev > 0 {
				prev = next[prev-1]
			} else {
				next[i] = prev
				i++
			}
		}
	}

	for i = patternLen; i > 0; i-- {
		next[i] = next[i-1]
	}
	next[0] = -1

	return next[:patternLen]
}
