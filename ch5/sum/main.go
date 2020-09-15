// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"sort"
	"strings"
)

//!+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func Join(sep string, args ...string) string {
	return strings.Join(args, sep)
}

//!-

func max(vals ...int) int {
	length := len(vals)
	if length == 0 {
		return 0
	}
	sort.Ints(vals)
	return vals[length-1]
}

func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	//!-slice

	//!+max
	fmt.Println(max())
	fmt.Println(max(7, 2, 3, 4))
}
