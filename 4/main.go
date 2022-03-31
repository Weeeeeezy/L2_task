package main

import (
	"fmt"
	"sort"
	"strings"
)

func uniqElem(in []string) []string {
	result := make([]string, 0, len(in))
	m := make(map[string]bool)

	for _, v := range in {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}

func anagram(in []string) map[string][]string {
	for i := range in {
		in[i] = strings.ToLower(in[i])
	}
	uniqIn := uniqElem(in)
	m := make(map[string][]string, 0)

	for _, v := range uniqIn {
		word := []rune(v)
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		st := string(word)

		m[st] = append(m[st], v)
	}

	result := make(map[string][]string, 0)

	for _, v := range m {
		if len(v) > 1 {
			result[v[0]] = v
			sort.Strings(v)
		}
	}

	return result

}

func main() {
	input := []string{"арора", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}

	fmt.Println(input)
	fmt.Println(anagram(input))
}
