package anagram

import (
	"fmt"
	"slices"
	"strings"
)

func SetsOfAnagrams(in []string) *map[string][]string {
	out := make(map[string][]string)

	tempHashes := makeMapOfSortedHashesUsingRunes(in)
	// tempHashes := makeMapOfSortedHashesUsingStrings(in)

	for _, anagrams := range tempHashes {
		if len(anagrams) > 1 {
			firstWord := anagrams[0]
			out[firstWord] = append([]string{}, anagrams...)
		}
	}
	return &out
}

func makeMapOfSortedHashesUsingRunes(in []string) map[string][]string {
	out := make(map[string][]string)
	for i := range in {
		str := strings.ToLower(in[i])
		runes := []rune(str)
		slices.Sort(runes)
		hash := string(runes)
		out[hash] = append(out[hash], str)
	}
	return out
}

func makeMapOfSortedHashesUsingStrings(in []string) map[string][]string {
	out := make(map[string][]string)
	for i := range in {
		str := strings.ToLower(in[i])
		letters := strings.Split(str, "")
		slices.Sort(letters)
		hash := strings.Join(letters, "")
		out[hash] = append(out[hash], str)
	}
	return out
}

func Run() {
	in := []string{"Пятак", "пЯтка", "тяпкА", "Листок", "слИток", "стОлик", "одувАн"}

	fmt.Println(SetsOfAnagrams(in))
}
