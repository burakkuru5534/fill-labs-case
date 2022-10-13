package main

import (
	"fmt"
	"github.com/burakkuru5534/src/funcs"
	"testing"
)

func TestSortBunchOfWordsByTheNumberOfCharacterA(t *testing.T) {
	//case 1: sorts a bunch of words by the number of character “a”s within the
	sorted := funcs.SortBunchOfWordsByTheNumberOfCharacterA([]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"})
	expected := []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"}

	if !equal(sorted, expected) {
		t.Errorf("expected %v, got %v", expected, sorted)
	}

}

func TestMostRepeatedDataInArray(t *testing.T) {
	//case3: find the most repeated word in a bunch of words
	mostRepeated := funcs.FindMostRepeatedDataInArray([]string{"apple", "pie", "apple", "red", "red", "red"})
	expected := "red"
	if mostRepeated != expected {
		t.Errorf("expected %v, got %v", expected, mostRepeated)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMostRepeatedForMultipleCases(t *testing.T) {
	var tests = []struct {
		myArr []string
		want  string
	}{
		{[]string{"apple", "pie", "apple", "red", "red", "red"}, "red"},
		{[]string{"madrid", "london", "london", "london", "madrid", "barcelona"}, "london"},
		{[]string{"Ukraine", "Ukraine", "USA", "USA", "UK", "Ukraine"}, "Ukraine"},
		{[]string{"Aws", "Gcp", "Gcp", "Aws", "Aws", "Aws"}, "Aws"},
		{[]string{"Sunny", "Rainy", "Sunny", "Rainy", "raiiny", "Sunny"}, "Sunny"},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%v,%s", tt.myArr, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := funcs.FindMostRepeatedDataInArray(tt.myArr)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestTestSortBunchOfWordsByTheNumberOfCharacterAForMultipleCases(t *testing.T) {
	var tests = []struct {
		myArr []string
		want  []string
	}{
		{[]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"},
			[]string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"}},

		{[]string{"aaaaz", "aaaab", "aaac", "aaab", "ddd", "eeee", "kkkkkk", "zzzzzzzzz", "p", "ca", "l"},
			[]string{"aaaab", "aaaaz", "aaab", "aaac", "ca", "zzzzzzzzz", "kkkkkk", "eeee", "ddd", "l", "p"}},

		{[]string{"aaaaaaaa", "a", "aaaaaaab", "aaaaaadc", "aaaaafde", "aaaacdes", "aaadeswx", "aadfghjk", "afghjklo", "dfghjklm", "cdcdsdws"},
			[]string{"aaaaaaaa", "aaaaaaab", "aaaaaadc", "aaaaafde", "aaaacdes", "aaadeswx", "aadfghjk", "afghjklo", "a", "cdcdsdws", "dfghjklm"}},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%v,%v", tt.myArr, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := funcs.SortBunchOfWordsByTheNumberOfCharacterA(tt.myArr)
			if !(equal(ans, tt.want)) {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
