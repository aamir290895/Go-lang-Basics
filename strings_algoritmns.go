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

func isAnagram2(s1, s2 string) bool {
	l1 := []rune(s1)
	l2 := []rune(s2)

	if len(l1) != len(l2) {
		return false
	}

	m := make(map[rune]bool)
	for _, v := range s2 {
		if strings.ContainsRune(s1, v) {
			m[v] = true
		} else {
			m[v] = false
		}
	}

	for key, val := range m {
		fmt.Println(key, val)

		if !val {
			return false
		}
	}

	return true
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

func strStr(haystack string, needle string) int {

	//Input: haystack = "sadbutsad", needle = "sad"
	//Output: 0
	// Explanation: "sad" occurs at index 0 and 6.
	//The first occurrence is at index 0, so we return 0.

	s, s1 := len(haystack), len(needle)
	for i := 0; i <= (s - s1); i++ {

		if haystack[i:i+s1] == needle {
			return i
		}

	}
	return -1

}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapping := make(map[byte]byte)
	visited := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if mapped, ok := mapping[s[i]]; ok {
			if mapped != t[i] {
				return false
			}
		} else {
			if visited[t[i]] {
				return false
			}
			mapping[s[i]] = t[i]
			visited[t[i]] = true
		}
	}

	return true
}

func repeatedSubstringPattern(s string) bool {
	// Duplicate the original string to form a longer string
	duplicated := s + s

	// Remove the first and last characters to avoid an exact match
	duplicated = duplicated[1 : len(duplicated)-1]

	fmt.Println(duplicated)
	// Check if the original string is present in the duplicated string
	return strings.Contains(duplicated, s)
}
