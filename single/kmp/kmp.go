package kmp

import (
	"fmt"
)

type pattern struct {
	str  string
	next []int
}

type Kmp struct {
	Db map[string][]pattern
}

/* You can define a dbKey to compile strings into a specified db */
func (kmp *Kmp) AddStringsToDb(strs []string, dbKey string) {
	if dbKey == "" {
		dbKey = "default"
	}
	if kmp.Db == nil {
		kmp.Db = make(map[string][]pattern)
	}

	for _, str := range strs {
		hasAdded := false
		for _, addedPattern := range kmp.Db[dbKey] {
			if addedPattern.str == str {
				hasAdded = true
				break
			}
		}
		if hasAdded == true {
			continue
		}

		var pat pattern
		pat.str = str
		kmp.Db[dbKey] = append(kmp.Db[dbKey], pat)
	}

	return
}

func (kmp *Kmp) Compile() error {
	return nil
}

func (kmp *Kmp) Show() {
	fmt.Print(kmp)
}

func Speak() {
	fmt.Println("hello")
}
