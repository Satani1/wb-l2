package main

import (
	"fmt"
	"strings"
)

func prepareWords(words []string) []string {
	result := make([]string, 0, len(words))

	for _, word := range words {
		str := strings.ToLower(word)
		result = append(result, str)
	}

	return result
}

func CheckWords(fString, sString string) bool {
	if len(fString) != len(sString) {
		return false
	}
	exist := make(map[rune]struct{})
	for _, char := range fString {
		exist[char] = struct{}{}
	}
	for _, char := range sString {
		if _, found := exist[char]; !found {
			return false
		}
	}

	return true
}

func SearchAnagram(words []string) map[string][]string {
	result := make(map[string][]string)
	exist := make(map[string]struct{})

	for i := 0; i < len(words); i++ {
		exist[words[i]] = struct{}{}
		for j := i + 1; j < len(words); j++ {
			if _, found := exist[words[j]]; found {
				continue
			}

			if CheckWords(words[i], words[j]) {
				result[words[i]] = append(result[words[i]], words[j])
				exist[words[j]] = struct{}{}
			}
		}
	}
	return result
}

func main() {
	words := []string{"ПЯТКА", "пяТак", "тяпка", "листОк", "Слиток", "свисток", "столИК", "стол"}
	words = prepareWords(words)
	fmt.Printf("Words:\n%v\n", words)
	fmt.Println("Anagrams:")
	fmt.Println(SearchAnagram(words))

}
