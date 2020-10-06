package kmp

import "testing"

func TestKmp(t *testing.T) {
	var test Kmp
	strs := []string{"http", "com"}
	test.AddStringsToDb(strs, "url")
	test.Show()
}
