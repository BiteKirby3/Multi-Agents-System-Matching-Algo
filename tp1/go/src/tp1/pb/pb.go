package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(word string) bool {
	chars := []rune(word)
	size := len(chars)
	palindrome := true
	for i := 0; i < size/2; i++ {
		if chars[i] != chars[size-1-i] {
			palindrome = false
			break
		}
	}
	return palindrome

}

func Palindromes(words []string) (l []string) {
	for i := range words {
		if IsPalindrome(words[i]) {
			l = append(l, words[i])
		}
	}
	return
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func Footprint(s string) string {
	//r := StringToRuneSlice(s)
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func Anagrams(words []string) (anagrams map[string][]string) {
	for i := range words {
		//verify if the current word's key already exists
		if val, ok := anagrams[Footprint(words[i])]; ok {
			//do something here
			val = append(val, words[i])
		} else {
			newSlice := []string{words[i]}
			anagrams[Footprint(words[i])] = newSlice
		}
	}
	return
}

func main() {
	dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}
	slice := dict[:]
	fmt.Println(Palindromes(slice))
	fmt.Println(Footprint("AGENT"))
	fmt.Println(Anagrams(slice))

}
