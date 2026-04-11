package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var trimRe = regexp.MustCompile(`^[^\p{L}-]+|[^\p{L}-]+$`)

func Top10(text string) []string {
	words := strings.Fields(text)
	wordsFreq := make(map[string]int)
	for _, word := range words {
		cleanedWord := cleanWord(word)
		if cleanedWord != "" {
			wordsFreq[cleanedWord]++
		}
	}

	uniqueWords := make([]string, 0, len(wordsFreq))
	for word := range wordsFreq {
		uniqueWords = append(uniqueWords, word)
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		freqI := wordsFreq[uniqueWords[i]]
		freqJ := wordsFreq[uniqueWords[j]]

		if freqI != freqJ {
			return freqI > freqJ
		}
		return uniqueWords[i] < uniqueWords[j]
	})

	result := uniqueWords
	if len(result) > 10 {
		result = result[:10]
	}
	return result
}

func cleanWord(word string) string {
	cleaned := trimRe.ReplaceAllString(word, "")
	if cleaned == "" || cleaned == "-" {
		return ""
	}
	return strings.ToLower(cleaned)
}
