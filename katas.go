package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GetCount("Success"))
	fmt.Println(Disemvowel("Success"))
	fmt.Println(DuplicateEncode("Success"))
}

// GetCount Return the number (count) of vowels in the given string.
// We will consider a, e, i, o, u as vowels for this Kata (but not y).
// The input string will only consist of lower case letters and/or spaces.
func GetCount(str string) (count int) {
	count = 0
	vowels := "aeiuo"
	for i := 0; i < len(str); i++ {
		if strings.Contains(vowels, string(str[i])) {
			count++
		}
	}
	return count
}

// Disemvowel Trolls are attacking your comment section!
// A common way to deal with this situation is to remove all the vowels from the trolls' comments, neutralizing the threat.
// Your task is to write a function that takes a string and return a new string with all vowels removed.
// For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".
func Disemvowel(comment string) string {
	vowels := []string{"a", "e", "i", "o", "u", "A", "O", "I", "E", "U"}
	for _, i := range vowels {
		comment = strings.Replace(comment, i, "", -1)
	}
	return comment
}

// DuplicateEncode The goal of this exercise is to convert a string to a new string where each character in the new string is
// "(" if that character appears only once in the original string, or ")" if that character appears more than
// once in the original string. Ignore capitalization when determining if a character is a duplicate.
func DuplicateEncode(word string) string {
	var substring string
	bracesString := ""
	for index, i := range word {
		substring = word[:index] + word[index+1:]
		if strings.Contains(strings.ToLower(substring), strings.ToLower(string(i))) {
			bracesString += ")"
		} else {
			bracesString += "("
		}
	}
	return bracesString
}
