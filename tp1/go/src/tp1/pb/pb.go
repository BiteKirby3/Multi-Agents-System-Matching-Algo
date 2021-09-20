package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	anagrams = make(map[string][]string)
	for i := range words {
		_, ok := anagrams[Footprint(words[i])]
		if ok {
			//do something here
			anagrams[Footprint(words[i])] = append(anagrams[Footprint(words[i])], words[i])
		} else {
			newSlice := []string{words[i]}
			anagrams[Footprint(words[i])] = newSlice
		}
	}
	return
}

func DictFromFile(filename string) (dict []string) {
	file, err := os.Open("tp1/go/src/tp1/pb/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dict
}

func main() {
	/*
		dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}
		slice := dict[:]
		fmt.Println(Palindromes(slice))
		fmt.Println(Footprint("AGENT"))
		fmt.Println(Anagrams(slice))
	*/
	dictFr := DictFromFile("dico-scrabble-fr.txt")
	sliceFr := dictFr[:]
	fmt.Println(Palindromes(sliceFr))                  // Le plus long : ESSAYASSE MALAYALAM RESSASSER
	fmt.Println(Anagrams(sliceFr)[Footprint("AGENT")]) // Les anagrammes de agents : [AGENT ETANG GANTE GEANT GENAT]
	ana := Anagrams(sliceFr)
	maxKey := ""
	maxNb := 0
	for k, v := range ana {
		if len(v) > maxNb {
			maxNb = len(v)
			maxKey = k
		}
	}
	fmt.Print(maxKey + " : ")
	fmt.Print(maxNb)
	fmt.Println(ana[maxKey]) // AEINRST (19) : [ARETINS ARISENT ENTRAIS INERTAS INSERAT RATINES RENTAIS RESINAT RETSINA RIANTES SATINER SENTIRA SERIANT SERINAT TANISER TARSIEN TRAINES TRANSIE TSARINE]

	sliceFrPalind := Palindromes(sliceFr)
	fmt.Println(Anagrams(sliceFrPalind)) // Existe-t-il un palindrome qui possède des anagrammes ? Non
}
