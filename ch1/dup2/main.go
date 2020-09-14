// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type DupLine struct {
	files []string
	n     int // count of dup lines
}

func main() {
	counts := make(map[string]DupLine)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, dupLine := range counts {
		if dupLine.n > 1 {
			fmt.Printf("%v\t%d\t%s\n", dupLine.files, dupLine.n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]DupLine) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		dupLine := counts[input.Text()]
		dupLine.files = append(dupLine.files, f.Name())
		dupLine.n++
		counts[input.Text()] = dupLine
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
