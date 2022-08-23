package main

import (
	"testing"

	chapter1 "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/Chapter1"
)

// Testing the IsUnique Function

func TestCase1_IsUnique(aTest *testing.T) {
	got := chapter1.IsUnique("Hello World")
	if got == true {
		aTest.Errorf("Hello World does not have unique characters.")
	}
}

func TestCase2_IsUnique(aTest *testing.T) {
	got := chapter1.IsUnique("Mouse")
	if got == false {
		aTest.Errorf("Mouse has unique characters.")
	}
}

func TestCase3_IsUnique(aTest *testing.T) {
	got := chapter1.IsUnique("palindrome")
	if got == false {
		aTest.Errorf("Palindrome has unique characters.")
	}
}

func TestCase1_CheckPermutation(aTest *testing.T) {
	var stringA = "Hello"
	var stringB = "Hello"
	got := chapter1.CheckPermutation(stringA, stringB)
	if got == false {
		aTest.Errorf("%s, %s are permutations of each other", stringA, stringB)
	}
}

func TestCase2_CheckPermutation(aTest *testing.T) {
	var stringA = "Different string"
	var stringB = "Completely different string"
	got := chapter1.CheckPermutation(stringA, stringB)
	if got == true {
		aTest.Errorf("%s, %s are not permutations of each other", stringA, stringB)
	}
}

func TestCase3_CheckPermutation(aTest *testing.T) {
	var stringA = "Same Length"
	var stringB = "Serm Length"
	got := chapter1.CheckPermutation(stringA, stringB)
	if got == true {
		aTest.Errorf("%s, %s are not permutations of each other", stringA, stringB)
	}
}

func TestCase4_CheckPermutation(aTest *testing.T) {
	var stringA = "ABC"
	var stringB = "CBA"
	got := chapter1.CheckPermutation(stringA, stringB)
	if got == false {
		aTest.Errorf("%s, %s are permutations of each other", stringA, stringB)
	}
}

// func TestCase5_URLify(aTest *testing.T) {
// 	var lString = "Matty is Cool"
// 	var lLength = 13
// 	var URLString = "Matty%20is%20Cool"
// 	got := chapter1.URLify(lString, lLength)
// 	fmt.Println([]byte{rune(lString...)})
// 	fmt.Println(reflect.ValueOf(URLString).Kind())
// 	if string(got) != string(URLString) {
// 		aTest.Errorf("The URLified expression should be %s. Your expression was %s.", URLString, got)
// 	}
// }

func TestCase6_IsPalindrome(aTest *testing.T) {
	var lString = "ABAB"
	got := chapter1.IsPalindromePermutation(lString)
	if got == false {
		aTest.Errorf("%s has a Palindrome Permutation", lString)
	}
}

func TestCase7_IsPalindrome(aTest *testing.T) {
	var lString = "taco cat"
	got := chapter1.IsPalindromePermutation(lString)
	if got == false {
		aTest.Errorf("%s has a palindrome permutation", lString)
	}
}

func TestCase8_IsPalindrome(aTest *testing.T) {
	var lString = "abcd"
	got := chapter1.IsPalindromePermutation(lString)
	if got == true {
		aTest.Errorf("%s does not have a palindrome permutation", lString)
	}
}

func TestCase9_IsPalindrome(aTest *testing.T) {
	var lString = "hamster"
	got := chapter1.IsPalindromePermutation(lString)
	if got == true {
		aTest.Errorf("%s does not have a palindrome permutation", lString)
	}
}

func TestCase10_IsPalindrome(aTest *testing.T) {
	var lString = "Don't nod."
	got := chapter1.IsPalindromePermutation(lString)
	if got == false {
		aTest.Errorf("%s has a palindrome permutation", lString)
	}
}

func TestCase11_IsPalindrome(aTest *testing.T) {
	var lString = "Step on no pets."
	got := chapter1.IsPalindromePermutation(lString)
	if got == false {
		aTest.Errorf("%s has a palindrome permutation", lString)
	}
}

func TestCase12_OneAway(aTest *testing.T) {
	var lStringA = "bale"
	var lStringB = "pale"
	got := chapter1.OneAway(lStringA, lStringB)
	if got == false {
		aTest.Errorf("%s and %s only have a 1 edit difference.", lStringA, lStringB)
	}
}

func TestCase13_OneAway(aTest *testing.T) {
	var lStringA = "bay"
	var lStringB = "baye"
	got := chapter1.OneAway(lStringA, lStringB)
	if got == false {
		aTest.Errorf("%s and %s only have a 1 edit difference.", lStringA, lStringB)
	}
}

func TestCase14_OneAway(aTest *testing.T) {
	var lStringA = "bay"
	var lStringB = "byy"
	got := chapter1.OneAway(lStringA, lStringB)
	if got == false {
		aTest.Errorf("%s and %s only have a 1 edit difference.", lStringA, lStringB)
	}
}

func TestCase15_OneAway(aTest *testing.T) {
	var lStringA = "pale"
	var lStringB = "bake"
	got := chapter1.OneAway(lStringA, lStringB)
	if got == true {
		aTest.Errorf("%s and %s only have a 2 edit difference.", lStringA, lStringB)
	}
}
