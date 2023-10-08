package main

import "log"

func reverseString(s string) string {
	// Convert the string to a rune slice
	runes := []rune(s)

	log.Println(runes[1])
	log.Println(runes[2])
	log.Println(runes[3])
	log.Println(runes[4])
	log.Println(runes[5])

	log.Println(string(runes[1]))
	log.Println(string(runes[2]))
	log.Println(string(runes[3]))
	log.Println(string(runes[4]))
	log.Println(string(runes[5]))

	// Reverse the rune slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string
	reversedString := string(runes)

	return reversedString
}
