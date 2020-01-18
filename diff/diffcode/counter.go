
package diffcode

import (
	"unicode"

)

func CharacterCount(str, bincard string) int {
	var count = 0
	var bincardslice []string
	for _, n := range bincard {
		bincardslice = append(bincardslice, string(n))
	}
	for i, r := range str {
		isLetter := unicode.IsLetter(r)
		isNumber := unicode.IsNumber(r)
		if !isLetter && bincardslice[i] != string("1") && !isNumber || string(r) == " " {
			count++
		}
	}
	return count
}

func TagCount(str, bincard string) int {
	var bincardslice []string
	var count = 0
	for _, n := range bincard {
		bincardslice = append(bincardslice, string(n))
	}
	for i, r := range str {
		if string(r) == "<" && bincardslice[i] == string("1") {
			count++
		}
	}
	return count
}

func WordCount(str string) uint {

	var nWords uint

	inWord := false
	for _, r := range str {
		nWords += startOfWord(&inWord, r)
	} 
		
	return nWords
}

func startOfWord(inWord *bool, r rune) uint {
	var isword uint
	if *inWord {
		if unicode.IsLetter(r) {
			isword = 0
		} else {
			*inWord = false
			isword = 0
		}
	} else {
		if unicode.IsLetter(r) {
			isword = 1
			*inWord = true
		} else {
			isword = 0
		}
	}
	return isword
}

