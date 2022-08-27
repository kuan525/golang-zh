package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	// 当长度不够的时候，将其延申到对应长度
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	// i是下标，word是对应下表的整数形式，拆成二进制就是对应的集合
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		// fmt.Println(i, word)
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				// 中间插入一个空格
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				// 这里i是偏移量，当前在第几个位置（64位）中，偏移量是j
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var s IntSet
	s.Add(12)
	s.Add(11)
	s.Add(111)
	fmt.Println(s.String())
}
