package chapter1

import (
	"sort"
	"unicode"
)

// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
// cannot use additional data structures?
func IsUnique(aString string) bool {
	var stringLookUp map[rune]int = make(map[rune]int, len(aString))
	for _, aRune := range aString {
		if val, ok := stringLookUp[aRune]; ok {
			stringLookUp[aRune] = val + 1
		} else {
			stringLookUp[aRune] = 0
		}
	}

	for _, v := range stringLookUp {
		if v >= 1 {
			return false
		}
	}
	return true
}

// Check Permutation: Given two strings, write a method to decide if one is a permutation of the
// other.
func CheckPermutation(aStringA string, aStringB string) bool {

	//For two strings to be permutations, they have to be equal in length.
	if len(aStringA) != len(aStringB) {
		return false
	}

	var runeArrayA = make([]rune, len(aStringA))
	var runeArrayB = make([]rune, len(aStringB))

	for aIndex, aRune := range aStringA {
		runeArrayA[aIndex] = aRune
	}

	for aIndex, aRune := range aStringB {
		runeArrayB[aIndex] = aRune
	}

	sort.Slice(runeArrayA, func(i, j int) bool {
		return runeArrayA[i] < runeArrayA[j]
	})

	sort.Slice(runeArrayB, func(i, j int) bool {
		return runeArrayB[i] < runeArrayB[j]
	})

	for lIndex := range runeArrayA {
		if runeArrayA[lIndex] != runeArrayB[lIndex] {
			return false
		}
	}
	return true
}

// URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: If implementing in Java, please use a character array so that you can
// perform this operation in place.)
func URLify(aString string, length int) string {
	var stringURL []rune = make([]rune, length)
	var URL string = "%20"
	var URLRuneArray []rune = make([]rune, len(URL))

	for aIndex, aRune := range URL {
		URLRuneArray[aIndex] = aRune
	}

	for lIndex := 0; lIndex < length; lIndex++ {
		var aRune rune = rune(aString[lIndex])
		if unicode.IsSpace(aRune) {
			stringURL = append(stringURL, URLRuneArray...)
		} else {
			stringURL = append(stringURL, aRune)
		}
	}

	return string(stringURL)
}

func contains(s []string, aRune rune) bool {
	for _, v := range s {
		if v == string(aRune) {
			return true
		}
	}

	return false
}

// Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome.
// A palindrome is a word or phrase that is the same forwards and backwards. A permutation
// is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
func IsPalindromePermutation(aString string) bool {
	var stringLookUp map[rune]int = make(map[rune]int, len(aString))
	var lGrammarFilterArray []string = []string{" ", ",", "'", "."}
	for _, aRune := range aString {
		if !contains(lGrammarFilterArray, aRune) {
			var aVal, ok = stringLookUp[aRune]
			if ok {
				stringLookUp[aRune] = aVal + 1
			} else {
				stringLookUp[aRune] = 0
			}
		}
	}

	singleCount := 0
	for _, aVal := range stringLookUp {
		if aVal == 0 {
			singleCount++
		} else if aVal > 1 {
			return false
		}
	}

	return singleCount <= 1
}
