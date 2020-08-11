// 统计字符数量
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "I found love, for me. Darling just dive right in, follow my lead."
	var wc wordsCounter
	fmt.Fprint(&wc, s)
	fmt.Printf("The world count is: %d", wc.ToInt())
}

type wordsCounter int

func (w *wordsCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords) // 默认split函数是 bufio.SanLines
	for scanner.Scan() {
		n++
	}

	*w = wordsCounter(n)
	err = nil
	return
}

func (w *wordsCounter) ToInt() int {
	return int(*w)
}
