package main

import "fmt"

func countSubstrings(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		ans += extendPalindromes(s, i, i)
		ans += extendPalindromes(s, i, i+1)
	}

	return ans
}

func extendPalindromes(s string, l, r int) int {
	count := 0

	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}

	return count
}
