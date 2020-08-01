package main

import "strings"

// a e i o u

func reverseVowels(s string) string {
	if len(s) == 0 {
		return ""
	}
	i := 0
	j := len(s) - 1
	tmpByte := []byte(s)
	for i <= j {
		if isVowel(string(s[i])) && isVowel(string(s[j])) {
			tmpByte[i], tmpByte[j] = tmpByte[j], tmpByte[i]
			i++
			j--
		} else if isVowel(string(s[i])) {
			j--
		} else {
			i++
		}
	}
	return s
}

func isVowel(sub string) bool {
	sub = strings.ToLower(sub)
	return sub == "a" || sub == "e" || sub == "i" || sub == "o" || sub == "u"
}
