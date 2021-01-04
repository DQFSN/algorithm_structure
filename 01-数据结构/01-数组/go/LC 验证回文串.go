package main

import (
	"strings"
)

func isPalindrome(s string) bool {
    s = strings.ToUpper(s)
    chs := []rune{}
    for _, v := range s {
        if v >= 'A' && v <= 'Z' || v >= '0' && v <= '9' {
            chs = append(chs, v)
        }
    }
    l, r := 0, len(chs)-1

    for l < r {
        if chs[l] != chs[r] {
            return false
        }
        l++
        r--
    }

    return true
}