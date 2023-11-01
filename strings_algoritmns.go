package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func reverseString(s string) string {
	// Convert the string to a rune slice
	runes := []rune(s)

	log.Println(runes[1])

	log.Println(string(runes[5]))

	// Reverse the rune slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string
	reversedString := string(runes)

	return reversedString
}

func isAnagram(str1, str2 string) bool {
	// Remove spaces and convert both strings to lowercase for case-insensitive comparison
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Check if the lengths of the strings are different
	if len(str1) != len(str2) {
		return false
	}

	// Convert the strings to slices of characters
	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")

	// Sort both slices
	sort.Strings(s1)
	sort.Strings(s2)

	// Compare the sorted slices
	return strings.Join(s1, "") == strings.Join(s2, "")
}

func occouranceOfString(str string, subStr string) int {

	count := 0

	for i := 0; i < len(str)-2; i++ {

		k := str[i : i+3]
		fmt.Println(k)

		if k == subStr {
			count++
		}
	}

	return count
}

func countVowelsInString(str string) map[string]int {

	m := make(map[string]int)

	vowels := "aeiouAEIOU"

	for _, v := range str {
		if strings.ContainsRune(vowels, v) {
			m[string(v)]++
		}

	}

	return m

}
