// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
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
			kk := countLines(f, counts)
			if kk == 1 {
				fmt.Printf("当前文件中有重复行：%s\n", arg)
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) int {
	input := bufio.NewScanner(f)
	mp := map[string]int{}

	ans := 0
	for input.Scan() {
		mp[input.Text()]++
		if mp[input.Text()] >= 2 {
			ans = 1
		}
		counts[input.Text()]++
	}
	return ans
}
