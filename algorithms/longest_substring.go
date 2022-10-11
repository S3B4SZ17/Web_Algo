package algorithms

import (
	"strings"
)

/*
3. Longest Substring Without Repeating Characters
Medium

Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func LengthOfLongestSubstring(s string) int {
	longest := 1
	word_split := strings.Split(s, "")
	var next_char int
	var substring string
	sequence := 1

	if len(s) == 0 {
		return 0
	}

	for i := 0; next_char < len(word_split)-1; i++ {
		next_char = i + 1
		substring += word_split[i]
		inSubstring := checkInSubstring(&substring, &word_split[next_char])
		if !inSubstring {
			sequence++
		} else {
			// Reset the values
			substring = getNewSubstring(&substring, &word_split[next_char])
			sequence = len(substring)
			sequence++
		}
		if sequence > longest {
			longest = sequence
		}
	}

	return longest
}

func checkInSubstring(substring *string, next_char *string) bool {
	if strings.Contains(*substring, *next_char) {
		return true
	} else {
		return false
	}
}

func getNewSubstring(substring *string, next_char *string) string {
	var new_substring string
	word_split := strings.Split(*substring, "")
	repeated_index := strings.LastIndex(*substring, *next_char)
	repeated_index++

	for repeated_index < len(word_split) {
		new_substring += word_split[repeated_index]
		repeated_index++
	}

	return new_substring
}
