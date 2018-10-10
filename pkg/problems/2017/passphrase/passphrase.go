package passphrase

import (
	"sort"
	"strings"
)

// Runes type
type Runes []rune

func (r Runes) Len() int           { return len(r) }
func (r Runes) Swap(i, j int)      { r[j], r[i] = r[i], r[j] }
func (r Runes) Less(i, j int) bool { return r[i] < r[j] }

func sortedString(s string) string {
	var runes Runes
	for _, r := range s {
		runes = append(runes, r)
	}
	sort.Sort(runes)
	return string(runes)
}

func containsAnagrams(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	seen := make(map[string]bool)
	for _, word := range words {
		word = sortedString(word)
		if seen[word] == true {
			return true
		}
		seen[word] = true
	}
	return false
}

func containsDuplicateWords(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	seen := make(map[string]bool)
	for _, word := range words {
		if seen[word] == true {
			return true
		}
		seen[word] = true
	}
	return false
}

func readPassphrases(input string) []string {
	return strings.Split(input, "\n")
}

func getValidPassphraseCount(passphrases []string, isInvalid func(string) bool) (numValid int) {
	for _, passphrase := range passphrases {
		if !isInvalid(passphrase) {
			numValid++
		}
	}
	return
}

// Part1 function
func Part1(input string) interface{} {
	var passphrases = readPassphrases(input)
	return getValidPassphraseCount(passphrases, containsDuplicateWords)
}

// Part2 function
func Part2(input string) interface{} {
	var passphrases = readPassphrases(input)
	return getValidPassphraseCount(passphrases, containsAnagrams)
}
