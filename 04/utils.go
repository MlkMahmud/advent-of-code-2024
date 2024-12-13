package day04

import "regexp"

func CharAt(str string, index int, defaultChar rune) rune {
	if index < 0 || index > len(str)-1 {
		return defaultChar
	}

	return rune(str[index])
}

func IsKeywordMatch(word string, pattern *regexp.Regexp) bool {
	return pattern.MatchString(word)
}

func SafeIndex[T any](arr []T, index int, defaultElement T) T {
	if index >= len(arr) {
		return defaultElement
	}

	return arr[index]
}
