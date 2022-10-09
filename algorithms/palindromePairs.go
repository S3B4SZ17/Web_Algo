package algorithms

import (
	"strings"
)

/*
Given a list of unique words, return all the pairs of the distinct indices (i, j)
in the given list, so that the concatenation of the two words words[i] + words[j] is a palindrome.

Example 1:

Input: words = ["abcd","dcba","lls","s","sssll"]
Output: [[0,1],[1,0],[3,2],[2,4]]
Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]
Example 2:

Input: words = ["bat","tab","cat"]
Output: [[0,1],[1,0]]
Explanation: The palindromes are ["battab","tabbat"]
Example 3:

Input: words = ["a",""]
Output: [[0,1],[1,0]]
*/

func PalindromePairs(words []string) [][]int {
	result := loopWorsArray(&words)
	return result
}

func loopWorsArray(words *[]string) [][]int {
	// Define general variables
	var result [][]int
	i := 0
	next := i + 1

	for i < len(*words)-1 {

		// Controls when to break the loop
		if next == len(*words) {
			i++
			next = i + 1
			if i == len(*words)-1 {
				break
			}
		}

		var combined string = (*words)[i] + (*words)[next]
		arr_result := combineWords(combined, i, next)
		// If the arr_result == {0,0} its an empty array, hence no palindrome
		if arr_result != nil {
			result = append(result, arr_result)
		}

		// We need to swtich the index of the 2 combined words and check if is a palindrome again
		var combined2 string = (*words)[next] + (*words)[i]
		arr_result = combineWords(combined2, next, i)
		if arr_result != nil {
			result = append(result, arr_result)
		}

		next++
	}

	return result
}

func combineWords(combinedWord string, index1 int, index2 int) []int {
	isPalindrome := IsPalindrom(combinedWord)

	if isPalindrome {
		return []int{index1, index2}
	} else {
		return nil
	}
}

func IsPalindrom(phrase string) bool {
	// Edge case where the phrase or word passed is just one letter long
	if len([]rune(phrase)) == 1 {
		return true
	}

	// Preparing the prhase passed. Trimed down all white spaces and make it lowercase
	phrase_trimmed := strings.ReplaceAll(phrase, " ", "")
	phrase_toLittle := strings.ToLower(phrase_trimmed)
	word_split := strings.Split(phrase_toLittle, "")

	next := len(word_split) - 1
	for i := 0; i <= len(word_split)/2; i++ {
		if word_split[i] != word_split[next] {
			return false
		}
		next = next - 1
	}
	return true
}
