package passphrase

import (
	"io/ioutil"
	"sort"
	"strings"
)

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

func readPassphrases(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var contents = strings.Trim(string(data), "\n")
	return strings.Split(contents, "\n")
}

func getValidPassphraseCount(passphrases []string, isInvalid func(string) bool) (numValid int) {
	for _, passphrase := range passphrases {
		if !isInvalid(passphrase) {
			numValid += 1
		}
	}
	return
}

func Part1(args []string) interface{} {
	var passphrases = readPassphrases(args[0])
	return getValidPassphraseCount(passphrases, containsDuplicateWords)
}

func Part2(args []string) interface{} {
	var passphrases = readPassphrases(args[0])
	return getValidPassphraseCount(passphrases, containsAnagrams)
}
