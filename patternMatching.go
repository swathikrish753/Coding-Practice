package main

import "fmt"

func PatternMatching() {
	str := "FUN_UNCLE"
	pattern := "UNCLE"
	n := len(str)
	m := len(pattern)
	for i := 0; i < n; i++ {
		match := true
		for j := 0; j < m; j++ {
			if i+j >= n || str[i+j] != pattern[j] {
				match = false
				break
			}

		}
		if match {
			fmt.Println("found the pattern")
		}
	}
}
